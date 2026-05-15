param(
    [string]$Version = "latest",
    [string]$InstallDir = "$env:LOCALAPPDATA\Programs\aws-creds-tool"
)

$Repo = "edgar-elian-rodriguez/aws-creds-tool"
if ($Version -eq "latest") {
    $DownloadBase = "https://github.com/$Repo/releases/latest/download"
} else {
    $DownloadBase = "https://github.com/$Repo/releases/download/$Version"
}

$arch = if ($env:PROCESSOR_ARCHITECTURE -eq "AMD64") { "x86_64" } elseif ($env:PROCESSOR_ARCHITECTURE -eq "ARM64") { "arm64" } else { "x86_64" }
$asset = "aws-creds-tool_Windows_$arch.zip"
$assetUrl = "$DownloadBase/$asset"
$checksumUrl = "$DownloadBase/checksums.txt"

$tmp = New-Item -ItemType Directory -Path ([System.IO.Path]::GetTempPath()) -Name (New-Guid) -Force
try {
    Write-Host "Downloading $assetUrl"
    Invoke-WebRequest -Uri $assetUrl -OutFile "$tmp\$asset" -UseBasicParsing

    Write-Host "Downloading checksums"
    Invoke-WebRequest -Uri $checksumUrl -OutFile "$tmp\checksums.txt" -UseBasicParsing

    $expected = Select-String -Path "$tmp\checksums.txt" -Pattern " $asset$" | ForEach-Object { $_.Line.Split(' ')[0] }
    if (-not $expected) {
        throw "Unable to find checksum for $asset"
    }

    $actual = Get-FileHash -Path "$tmp\$asset" -Algorithm SHA256 | Select-Object -ExpandProperty Hash
    if ($expected -ne $actual) {
        throw "Checksum mismatch"
    }

    Write-Host "Extracting $asset"
    Expand-Archive -Path "$tmp\$asset" -DestinationPath "$tmp\extract" -Force

    $binary = Get-ChildItem -Path "$tmp\extract" -Filter aws-creds-tool -Recurse | Select-Object -First 1
    if (-not $binary) {
        throw "Binary not found in archive"
    }

    if (-not (Test-Path -Path $InstallDir)) {
        New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
    }

    Copy-Item -Path $binary.FullName -Destination "$InstallDir\aws-creds-tool.exe" -Force
    Write-Host "Installed aws-creds-tool to $InstallDir\aws-creds-tool.exe"
    Write-Host "If $InstallDir is not on your PATH, add it and reopen your shell."
} finally {
    Remove-Item -Recurse -Force -Path $tmp
}
