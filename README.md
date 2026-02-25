# AWS Credential Administrator (`aws-creds-tool`)

A multi-platform Go CLI for managing your local AWS profiles and generating temporary MFA session tokens.

## Why this tool?

When Multi-Factor Authentication (MFA) is enabled on an AWS account, using the AWS CLI requires obtaining temporary credentials via the `sts:GetSessionToken` API. This process is often manual and repetitive (logging into the console to get the MFA ARN, running CLI commands, and copying/pasting tokens into the `credentials` file).

`aws-creds-tool` was created to automate this workflow. It handles the MFA prompt, interacts with AWS STS to get temporary credentials, and automatically updates your local AWS credentials file, making it seamless to switch between profiles and keep your environment ready for use.

## Built With

- **[Go](https://go.dev/)**: Chosen for its ability to produce high-performance, cross-platform static binaries without external runtime requirements.
- **[Cobra](https://github.com/spf13/cobra)**: The industry standard for building powerful and structured CLI applications in Go.
- **[Huh?](https://github.com/charmbracelet/huh)**: A simple, powerful library for building terminal-based forms and interactive TUIs, giving the tool its user-friendly interface.
- **[GoReleaser](https://goreleaser.com/)**: Used to automate the build and release process for multiple platforms (Windows, macOS, Linux).
- **[GitHub Actions](https://github.com/features/actions)**: Provides the automation pipeline to build and publish releases whenever a new tag is pushed.

## Features
- Support for Windows, macOS, and Linux.
- Interactive TUI for selecting profiles and inputting MFA codes.
- Automatic updating of AWS credentials files.

## Installation

### From GitHub Releases
1. Go to the [Releases](https://github.com/Elian-Rodriguez-Dev/aws-creds-tool/releases) page.
2. Download the archive for your operating system and architecture.
3. Extract the binary and move it to a folder in your `PATH`.
   - **Linux/macOS:** `/usr/local/bin` is a common choice.
   - **Windows:** You can add a custom folder to your System Environment Variables.

### Using `go install`
If you have Go installed, you can install it directly:
```bash
go install github.com/Elian-Rodriguez-Dev/aws-creds-tool@latest
```

## Usage
Once installed, you can call the tool from your command shell:
```bash
aws-creds-tool --help
```

## Contributing
Please feel free to open issues or submit pull requests.

## License
[MIT](LICENSE) (or your preferred license)
