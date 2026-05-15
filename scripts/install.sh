#!/usr/bin/env bash
set -euo pipefail

REPO="edgar-elian-rodriguez/aws-creds-tool"
TAG="${1:-latest}"
INSTALL_DIR="${2:-/usr/local/bin}"
TMPDIR="$(mktemp -d)"

cleanup() {
  rm -rf "$TMPDIR"
}
trap cleanup EXIT

if [ "$(uname -s)" = "Linux" ]; then
  OS="Linux"
elif [ "$(uname -s)" = "Darwin" ]; then
  OS="Darwin"
else
  echo "Unsupported OS: $(uname -s)"
  exit 1
fi

case "$(uname -m)" in
  x86_64|amd64)
    ARCH="x86_64"
    ;;
  arm64|aarch64)
    ARCH="arm64"
    ;;
  i386|i686)
    ARCH="i386"
    ;;
  *)
    echo "Unsupported architecture: $(uname -m)"
    exit 1
    ;;
esac

ASSET="aws-creds-tool_${OS}_${ARCH}.tar.gz"
RELEASE_BASE="https://github.com/$REPO/releases"
if [ "$TAG" = "latest" ]; then
  DOWNLOAD_BASE="$RELEASE_BASE/latest/download"
else
  DOWNLOAD_BASE="$RELEASE_BASE/download/$TAG"
fi

ASSET_URL="$DOWNLOAD_BASE/$ASSET"
CHECKSUM_URL="$DOWNLOAD_BASE/checksums.txt"

cd "$TMPDIR"

echo "Downloading $ASSET_URL"
curl -fsSL -o "$ASSET" "$ASSET_URL"

echo "Downloading checksums"
curl -fsSL -o checksums.txt "$CHECKSUM_URL"

if command -v sha256sum >/dev/null 2>&1; then
  SHASUM="sha256sum"
elif command -v shasum >/dev/null 2>&1; then
  SHASUM="shasum -a 256"
else
  echo "Missing sha256 checksum utility"
  exit 1
fi

EXPECTED_CHECKSUM=$(grep " $ASSET$" checksums.txt | awk '{print $1}')
if [ -z "$EXPECTED_CHECKSUM" ]; then
  echo "Unable to find checksum for $ASSET"
  exit 1
fi

ACTUAL_CHECKSUM=$(eval "$SHASUM $ASSET" | awk '{print $1}')
if [ "$EXPECTED_CHECKSUM" != "$ACTUAL_CHECKSUM" ]; then
  echo "Checksum mismatch"
  exit 1
fi

echo "Extracting $ASSET"
mkdir -p extracted
tar -xzf "$ASSET" -C extracted

BINARY_PATH="$(find extracted -type f -name aws-creds-tool | head -n 1)"
if [ -z "$BINARY_PATH" ]; then
  echo "Binary not found in archive"
  exit 1
fi

if [ ! -w "$INSTALL_DIR" ]; then
  SUDO=sudo
else
  SUDO=""
fi

$SUDO mkdir -p "$INSTALL_DIR"
$SUDO install -m 0755 "$BINARY_PATH" "$INSTALL_DIR/aws-creds-tool"

echo "Installed aws-creds-tool to $INSTALL_DIR/aws-creds-tool"
echo "If $INSTALL_DIR is not on your PATH, add it and reopen your shell."
