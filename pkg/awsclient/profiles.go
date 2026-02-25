package awsclient

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

// Profile represents an AWS profile found in the config or credentials file.
type Profile struct {
	Name string
}

// GetProfileFiles returns the paths to the AWS credentials and config files.
func GetProfileFiles() (string, string) {
	home, _ := os.UserHomeDir()
	credentialsPath := filepath.Join(home, ".aws", "credentials")
	configPath := filepath.Join(home, ".aws", "config")
	return credentialsPath, configPath
}

// ListProfiles reads the AWS credentials and config files and returns a list of unique profile names.
func ListProfiles() ([]string, error) {
	credentialsPath, configPath := GetProfileFiles()
	profiles := make(map[string]bool)

	// Helper to load and extract profiles
	loadProfiles := func(path string) {
		cfg, err := ini.Load(path)
		if err != nil {
			return
		}
		for _, section := range cfg.Sections() {
			name := section.Name()
			if name == "DEFAULT" {
				continue
			}
			// In config file, profiles are often named [profile name]
			if filepath.Base(path) == "config" {
				if len(name) > 8 && name[:8] == "profile " {
					name = name[8:]
				}
			}
			profiles[name] = true
		}
	}

	loadProfiles(credentialsPath)
	loadProfiles(configPath)

	var profileList []string
	for p := range profiles {
		profileList = append(profileList, p)
	}
	return profileList, nil
}

// UpdateCredentials updates the ~/.aws/credentials file with new temporary credentials.
func UpdateCredentials(profileName string, accessKey, secretKey, sessionToken string) error {
	credentialsPath, _ := GetProfileFiles()
	
	cfg, err := ini.Load(credentialsPath)
	if err != nil {
		if os.IsNotExist(err) {
			cfg = ini.Empty()
		} else {
			return fmt.Errorf("failed to load credentials file: %w", err)
		}
	}

	section := cfg.Section(profileName)
	section.Key("aws_access_key_id").SetValue(accessKey)
	section.Key("aws_secret_access_key").SetValue(secretKey)
	section.Key("aws_session_token").SetValue(sessionToken)

	err = cfg.SaveTo(credentialsPath)
	if err != nil {
		return fmt.Errorf("failed to save credentials file: %w", err)
	}

	return nil
}
