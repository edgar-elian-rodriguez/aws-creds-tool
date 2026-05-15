# AWS Credential Administrator (`aws-creds-tool`)

[![CI](https://github.com/edgar-elian-rodriguez/aws-creds-tool/actions/workflows/ci.yml/badge.svg)](https://github.com/edgar-elian-rodriguez/aws-creds-tool/actions/workflows/ci.yml)

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

### Using `go install` (Recommended)

Requires [Go](https://go.dev/dl/) 1.25 or later:

```bash
go install github.com/edgar-elian-rodriguez/aws-creds-tool@latest
```

The binary is placed in `$GOPATH/bin` (usually `~/go/bin`). Make sure that directory is on your `PATH`.

### From GitHub Releases

Pre-built binaries for Windows, macOS, and Linux are available on the [Releases](https://github.com/edgar-elian-rodriguez/aws-creds-tool/releases) page.

1. Download the archive for your OS and architecture.
2. Extract the binary and place it in a directory on your `PATH`.
   - **Linux/macOS:** `/usr/local/bin` is a common choice.
  - **Windows:** run `install.ps1` from the extracted archive to copy `aws-creds-tool.exe` into `%LOCALAPPDATA%\Programs\aws-creds-tool`, or place the binary manually in a directory on your `PATH`.

#### Linux packages
Each release also publishes Debian and Red Hat package artifacts.

- Debian:
  ```bash
  sudo dpkg -i aws-creds-tool_<version>_amd64.deb
  ```
- Red Hat:
  ```bash
  sudo rpm -i aws-creds-tool-<version>.rpm
  ```

#### Install scripts
A quick way to install on Linux or macOS is the included shell installer:

```bash
curl -fsSL https://raw.githubusercontent.com/edgar-elian-rodriguez/aws-creds-tool/main/scripts/install.sh | bash
```

For Windows, use the PowerShell installer script:

```powershell
Invoke-WebRequest -Uri "https://raw.githubusercontent.com/edgar-elian-rodriguez/aws-creds-tool/main/scripts/install.ps1" -OutFile "$env:TEMP\install.ps1"
pwsh $env:TEMP\install.ps1
```

#### Homebrew
A Homebrew tap exists at `edgar-elian-rodriguez/homebrew-tap`. Once the formula is published, install with:

```bash
brew install edgar-elian-rodriguez/tap/aws-creds-tool
```

#### Scoop
A Scoop bucket exists at `edgar-elian-rodriguez/scoop-bucket`. Once the manifest is published, install with:

```powershell
scoop bucket add aws-creds-tool https://github.com/edgar-elian-rodriguez/scoop-bucket
scoop install aws-creds-tool
```

## Usage
Once installed, you can call the tool from your command shell:
```bash
aws-creds-tool --help
```

### Onboarding a new AWS profile
If you do not have any AWS profiles configured locally, create one with:
```bash
aws-creds-tool onboard
```
The command prompts for a profile name, AWS access key ID, secret access key, default region, and output format, then writes the profile to both `~/.aws/credentials` and `~/.aws/config`.

If `aws-creds-tool mfa` runs and no AWS profiles are found, it will ask whether you want to onboard a new profile and will launch the onboard flow automatically if you choose Yes.

## Contributing
Please feel free to open issues or submit pull requests.

## License
[MIT](LICENSE) (or your preferred license)
