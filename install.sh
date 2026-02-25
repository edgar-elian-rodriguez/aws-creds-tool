#!/bin/bash

# AWS Credential Tool Installer
# Works on macOS and Linux

set -e

OWNER="Elian-Rodriguez-Dev"
REPO="aws-creds-tool"
BINARY_NAME="aws-creds-tool"

# Detect OS
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
case "$OS" in
    darwin*) OS="Darwin" ;;
    linux*) OS="Linux" ;;
    *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

# Detect Architecture
ARCH=$(uname -m)
case "$ARCH" in
    x86_64) ARCH="x86_64" ;;
    arm64|aarch64) ARCH="arm64" ;;
    *) echo "Unsupported Architecture: $ARCH"; exit 1 ;;
esac

echo "Detected $OS ($ARCH)..."

# Fetch latest release version
VERSION=$(curl -s "https://api.github.com/repos/$OWNER/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$VERSION" ]; then
    echo "Error: Could not determine latest version."
    exit 1
fi

echo "Installing $REPO $VERSION..."

# Download URL
# Pattern: aws-creds-tool_Darwin_arm64.tar.gz or aws-creds-tool_Linux_x86_64.tar.gz
FILENAME="${BINARY_NAME}_${OS}_${ARCH}.tar.gz"
URL="https://github.com/$OWNER/$REPO/releases/download/$VERSION/$FILENAME"

# Download and Extract
TEMP_DIR=$(mktemp -d)
curl -L "$URL" -o "$TEMP_DIR/$FILENAME"
tar -xzf "$TEMP_DIR/$FILENAME" -C "$TEMP_DIR"

# Move to bin folder
INSTALL_PATH="/usr/local/bin/$BINARY_NAME"
echo "Moving binary to $INSTALL_PATH..."

if [ -w "/usr/local/bin" ]; then
    mv "$TEMP_DIR/$BINARY_NAME" "$INSTALL_PATH"
else
    echo "Permission denied. Requesting sudo to install to /usr/local/bin..."
    sudo mv "$TEMP_DIR/$BINARY_NAME" "$INSTALL_PATH"
fi

chmod +x "$INSTALL_PATH"

# Cleanup
rm -rf "$TEMP_DIR"

echo ""
echo "✓ $BINARY_NAME installed successfully to $INSTALL_PATH"
echo "Run '$BINARY_NAME --help' to get started!"
