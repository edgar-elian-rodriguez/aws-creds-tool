package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/edgar-elian-rodriguez/aws-creds-tool/pkg/awsclient"
	"github.com/edgar-elian-rodriguez/aws-creds-tool/pkg/ux"
	"github.com/spf13/cobra"
)

var (
	sourceProfile string
	targetProfile string
	duration      int32
)

var mfaCmd = &cobra.Command{
	Use:   "mfa",
	Short: "Generate temporary MFA credentials",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		// 1. Discover or select source profile
		if sourceProfile == "" {
			profiles, err := awsclient.ListProfiles()
			if err != nil {
				log.Fatalf("Failed to list profiles: %v", err)
			}

			if len(profiles) == 0 {
				onboard, err := ux.ConfirmOnboard()
				if err != nil {
					log.Fatalf("Onboarding prompt failed: %v", err)
				}

				if onboard {
					sourceProfile, err = runOnboard()
					if err != nil {
						log.Fatalf("Onboarding failed: %v", err)
					}
				} else {
					log.Fatal("No AWS profiles found. Aborting. Run 'aws-creds-tool onboard' to create one.")
				}
			} else {
				sourceProfile, err = ux.SelectProfile(profiles)
				if err != nil {
					log.Fatalf("Profile selection failed: %v", err)
				}
			}
		}

		fmt.Printf("Using profile: %s\n", sourceProfile)

		// 2. Discover MFA ARN
		fmt.Println("Discovering MFA device...")
		mfaArn, err := awsclient.GetMfaSerialNumber(ctx, sourceProfile)
		if err != nil {
			log.Fatalf("Failed to discover MFA ARN: %v", err)
		}
		fmt.Printf("MFA Device: %s\n", mfaArn)

		// 3. Prompt for MFA code
		mfaCode, err := ux.PromptMfaCode()
		if err != nil {
			log.Fatalf("MFA code input failed: %v", err)
		}

		// 4. Get Session Token
		fmt.Println("Requesting temporary credentials...")
		sessionToken, err := awsclient.GetSessionToken(ctx, sourceProfile, mfaArn, mfaCode, duration)
		if err != nil {
			log.Fatalf("Failed to get session token: %v", err)
		}

		// 5. Save credentials
		if targetProfile == "" {
			targetProfile = fmt.Sprintf("%s-mfa", sourceProfile)
		}

		err = awsclient.UpdateCredentials(
			targetProfile,
			*sessionToken.Credentials.AccessKeyId,
			*sessionToken.Credentials.SecretAccessKey,
			*sessionToken.Credentials.SessionToken,
		)
		if err != nil {
			log.Fatalf("Failed to save credentials: %v", err)
		}

		fmt.Printf("\n✓ MFA credentials configured successfully!\n")
		fmt.Printf("Profile Name: %s\n", targetProfile)
		fmt.Printf("Expiration: %v\n", sessionToken.Credentials.Expiration)
	},
}

func init() {
	rootCmd.AddCommand(mfaCmd)

	mfaCmd.Flags().StringVarP(&sourceProfile, "profile", "p", "", "Source AWS profile")
	mfaCmd.Flags().StringVarP(&targetProfile, "target", "t", "", "Target AWS profile for MFA credentials")
	mfaCmd.Flags().Int32VarP(&duration, "duration", "d", 43200, "Session duration in seconds")
}
