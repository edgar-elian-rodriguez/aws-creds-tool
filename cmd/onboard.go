package cmd

import (
	"fmt"
	"log"

	"github.com/edgar-elian-rodriguez/aws-creds-tool/pkg/awsclient"
	"github.com/edgar-elian-rodriguez/aws-creds-tool/pkg/ux"
	"github.com/spf13/cobra"
)

var onboardCmd = &cobra.Command{
	Use:   "onboard",
	Short: "Create a new AWS profile on this machine",
	Run: func(cmd *cobra.Command, args []string) {
		profileName, err := runOnboard()
		if err != nil {
			log.Fatalf("Onboarding failed: %v", err)
		}

		fmt.Printf("\n✓ AWS profile '%s' created successfully.\n", profileName)
		fmt.Println("You can now run `aws-creds-tool mfa` to generate temporary credentials.")
	},
}

func init() {
	rootCmd.AddCommand(onboardCmd)
}

func runOnboard() (string, error) {
	profile, err := ux.PromptOnboardProfile()
	if err != nil {
		return "", err
	}

	err = awsclient.CreateProfile(profile.Name, profile.AccessKey, profile.SecretKey, profile.Region, profile.Output)
	if err != nil {
		return "", err
	}

	return profile.Name, nil
}
