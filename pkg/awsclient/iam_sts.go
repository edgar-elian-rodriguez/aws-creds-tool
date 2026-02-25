package awsclient

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// GetMfaSerialNumber discovers the first MFA device ARN for the current user.
func GetMfaSerialNumber(ctx context.Context, profile string) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile(profile))
	if err != nil {
		return "", fmt.Errorf("unable to load SDK config: %v", err)
	}

	iamClient := iam.NewFromConfig(cfg)
	
	// List MFA devices for the current user (calling with nil user name lists for calling user)
	output, err := iamClient.ListMFADevices(ctx, &iam.ListMFADevicesInput{})
	if err != nil {
		return "", fmt.Errorf("failed to list MFA devices: %v", err)
	}

	if len(output.MFADevices) == 0 {
		return "", fmt.Errorf("no MFA devices found for this profile")
	}

	return *output.MFADevices[0].SerialNumber, nil
}

// GetSessionToken calls STS to get temporary credentials using MFA.
func GetSessionToken(ctx context.Context, profile, serialNumber, tokenCode string, duration int32) (*sts.GetSessionTokenOutput, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile(profile))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config: %v", err)
	}

	stsClient := sts.NewFromConfig(cfg)
	
	input := &sts.GetSessionTokenInput{
		SerialNumber:    aws.String(serialNumber),
		TokenCode:       aws.String(tokenCode),
		DurationSeconds: aws.Int32(duration),
	}

	output, err := stsClient.GetSessionToken(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to get session token: %v", err)
	}

	return output, nil
}
