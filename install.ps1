# AWS Credential Tool Installer for Windows
# Run with: iwr https://raw.githubusercontent.com/.../install.ps1 | iex

$ErrorActionPreference = "Stop"

$Owner = "Elian-Rodriguez-Dev"
$Repo = "aws-creds-tool"
$BinaryName = "aws-creds-tool"

# Detect Architecture
$Arch = if ($Is64BitProcess) { "x86_64" } else { "i386" }
if ($env:PROCESSOR_ARCHITECTURE -eq "ARM64") { $Arch = "arm64" }

Write-Host "Detected Windows ($Arch)..." -ForegroundColor Cyan

# Fetch latest release version
$LatestRelease = Invoke-RestMethod -Uri "https://api.github.com/repos/$Owner/$Repo/releases/latest"
$Version = $LatestRelease.tag_name

Write-Host "Installing $Repo $Version..." -ForegroundColor Cyan

# Download URL
# Pattern: aws-creds-tool_Windows_x86_64.zip
$Filename = "${BinaryName}_Windows_${Arch}.zip"
$Asset = $LatestRelease.assets | Where-Object { $_.name -eq $Filename }

if (-not $Asset) {
    Write-Error "Could not find asset $Filename for version $Version"
}

$Url = $Asset.browser_download_url

# Temporary location
$TempDir = Join-Path $env:TEMP "aws-creds-tool-install"
if (Test-Path $TempDir) { Remove-Item $TempDir -Recurse -Force }
New-Item -ItemType Directory -Path $TempDir | Out-Null

$ZipPath = Join-Path $TempDir $Filename
Write-Host "Downloading $Url..."
Invoke-WebRequest -Uri $Url -OutFile $ZipPath

# Extract
Write-Host "Extracting..."
Expand-Archive -Path $ZipPath -DestinationPath $TempDir -Force

# Define Install Location
$InstallDir = Join-Path $env:USERPROFILE ".aws-creds-tool"
if (-not (Test-Path $InstallDir)) { New-Item -ItemType Directory -Path $InstallDir | Out-Null }

$ExePath = Join-Path $InstallDir "${BinaryName}.exe"
Move-Item -Path (Join-Path $TempDir "${BinaryName}.exe") -Destination $ExePath -Force

# Add to PATH if not already there
$UserPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($UserPath -notlike "*$InstallDir*") {
    Write-Host "Adding $InstallDir to User PATH..." -ForegroundColor Yellow
    [Environment]::SetEnvironmentVariable("Path", "$UserPath;$InstallDir", "User")
    $env:Path += ";$InstallDir"
}

# Cleanup
Remove-Item $TempDir -Recurse -Force

Write-Host "`n✓ $BinaryName installed successfully!" -ForegroundColor Green
Write-Host "Location: $ExePath"
Write-Host "Restart your terminal and run '$BinaryName --help' to get started!" -ForegroundColor Green
