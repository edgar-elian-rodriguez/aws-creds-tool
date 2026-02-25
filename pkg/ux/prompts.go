package ux

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

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
