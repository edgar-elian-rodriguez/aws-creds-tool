package cmd

import (
	"fmt"
	"log"

	"github.com/Elian-Rodriguez-Dev/aws-creds-tool/pkg/awsclient"
	"github.com/spf13/cobra"
)

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Administrate local AWS profiles",
}

var listProfilesCmd = &cobra.Command{
	Use:   "list",
	Short: "List all discovered profiles",
	Run: func(cmd *cobra.Command, args []string) {
		profiles, err := awsclient.ListProfiles()
		if err != nil {
			log.Fatalf("Failed to list profiles: %v", err)
		}

		fmt.Println("Discovered AWS Profiles:")
		for _, p := range profiles {
			fmt.Printf(" - %s\n", p)
		}
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
	profileCmd.AddCommand(listProfilesCmd)
}
