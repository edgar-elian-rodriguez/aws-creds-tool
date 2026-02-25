package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aws-creds-tool",
	Short: "A tool for administrating AWS credentials and MFA tokens",
	Long: `AWS Credential Administrator is a multi-platform Go CLI for 
managing your local AWS profiles and generating temporary MFA session tokens.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Root flags can be defined here
}
