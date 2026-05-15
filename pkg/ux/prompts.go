package ux

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
)

type OnboardProfile struct {
	Name      string
	AccessKey string
	SecretKey string
	Region    string
	Output    string
}

// ConfirmOnboard asks the user whether to onboard a new AWS profile.
func ConfirmOnboard() (bool, error) {
	var selected string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("No AWS profiles were found. Would you like to onboard a new profile?").
				Options(
					huh.NewOption("Yes", "yes"),
					huh.NewOption("No", "no"),
				).
				Value(&selected),
		),
	)

	err := form.Run()
	if err != nil {
		return false, err
	}

	return selected == "yes", nil
}

// PromptOnboardProfile collects profile information for onboarding.
func PromptOnboardProfile() (OnboardProfile, error) {
	var profile OnboardProfile

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("AWS Profile Name").
				Description("Enter a local profile name to store AWS credentials.").
				Placeholder("default").
				Validate(func(str string) error {
					if strings.TrimSpace(str) == "" {
						return fmt.Errorf("profile name is required")
					}
					return nil
				}).
				Value(&profile.Name),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("AWS Access Key ID").
				Description("Enter your AWS access key ID.").
				Validate(func(str string) error {
					if strings.TrimSpace(str) == "" {
						return fmt.Errorf("access key ID is required")
					}
					return nil
				}).
				Value(&profile.AccessKey),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("AWS Secret Access Key").
				Description("Enter your AWS secret access key.").
				Validate(func(str string) error {
					if strings.TrimSpace(str) == "" {
						return fmt.Errorf("secret access key is required")
					}
					return nil
				}).
				Value(&profile.SecretKey),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Default AWS Region").
				Description("Enter the default AWS region for this profile.").
				Placeholder("us-east-1").
				Validate(func(str string) error {
					if strings.TrimSpace(str) == "" {
						return fmt.Errorf("region is required")
					}
					return nil
				}).
				Value(&profile.Region),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Default AWS Output Format").
				Description("Enter the default output format for this profile.").
				Placeholder("json").
				Value(&profile.Output),
		),
	)

	err := form.Run()
	if err != nil {
		return OnboardProfile{}, err
	}

	profile.Name = strings.TrimSpace(profile.Name)
	profile.AccessKey = strings.TrimSpace(profile.AccessKey)
	profile.SecretKey = strings.TrimSpace(profile.SecretKey)
	profile.Region = strings.TrimSpace(profile.Region)
	profile.Output = strings.TrimSpace(profile.Output)

	return profile, nil
}

// SelectProfile presents a list of profiles and returns the selected one.
func SelectProfile(profiles []string) (string, error) {
	if len(profiles) == 0 {
		return "", fmt.Errorf("no AWS profiles found")
	}

	var selected string
	var options []huh.Option[string]
	for _, p := range profiles {
		options = append(options, huh.NewOption(p, p))
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select an AWS Profile").
				Options(options...).
				Value(&selected),
		),
	)

	err := form.Run()
	if err != nil {
		return "", err
	}

	return selected, nil
}

// PromptMfaCode asks the user for the 6-digit MFA code.
func PromptMfaCode() (string, error) {
	var code string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("MFA Code").
				Description("Enter the 6-digit code from your authenticator app").
				Placeholder("123456").
				Validate(func(str string) error {
					if len(str) != 6 {
						return fmt.Errorf("code must be 6 digits")
					}
					return nil
				}).
				Value(&code),
		),
	)

	err := form.Run()
	if err != nil {
		return "", err
	}

	return code, nil
}
