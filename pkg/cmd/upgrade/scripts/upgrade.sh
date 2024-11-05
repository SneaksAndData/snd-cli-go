#!/bin/bash
# Installation script for SnD CLI
set -e
REPO_OWNER="SneaksAndData"
REPO_NAME="snd-cli-go"
LATEST_RELEASE_TAG=$(curl --silent "https://api.github.com/repos/$REPO_OWNER/$REPO_NAME/releases/latest" | jq -r .tag_name)

VERSION=$(echo $LATEST_RELEASE_TAG | sed 's/^v//')
BASE_BIN_NAME="snd-cli-go_$VERSION"

echo "Determining target OS and architecture..."
ARCH=$(uname -m)
if [[ "$OSTYPE" =~ ^darwin ]]; then
     if [[ "$ARCH" =~ arm64 ]]; then
        asset_name="${BASE_BIN_NAME}_darwin_arm64.tar.gz"
      elif [[ "$ARCH" =~ (amd64|x86_64|x64) ]]; then
        asset_name="${BASE_BIN_NAME}_darwin_amd64.tar.gz"
      fi
elif [[ "$OSTYPE" =~ ^linux ]]; then
  if [[ "$ARCH" =~ arm64 ]]; then
    asset_name="${BASE_BIN_NAME}_linux_arm64.tar.gz"
  elif [[ "$ARCH" =~ (amd64|x86_64|x64) ]]; then
    asset_name="${BASE_BIN_NAME}_linux_amd64.tar.gz"
  fi
else
    echo "Error: Unsupported OS type or architecture: $OSTYPE $ARCH "
    exit 1
fi
echo "Target OS: $OSTYPE"
echo "Target ARCH: $ARCH"

BUNDLE_URL="https://github.com/$REPO_OWNER/$REPO_NAME/releases/download/$LATEST_RELEASE_TAG/$asset_name"

# Define base path for the application
base_path="$HOME/.local/snd-cli"
mkdir -p "$base_path"

echo "Downloading the binary from $BUNDLE_URL"
# Get file
curl -L -o "$base_path/$asset_name" "$BUNDLE_URL"

echo "Extracting the binary..."
tar -xzf "$base_path/$asset_name" -C "$base_path"

chmod +x "$base_path/snd-cli-go"

if [ -e "$HOME/.local/bin/snd" ]; then
   echo "Removing symlink..."
   rm $HOME/.local/bin/snd
fi

# Create a symbolic link to the application
echo "Creating the symlink..."
ln -s "$base_path/snd-cli-go" "$HOME/.local/bin/snd"

echo "Please restart your terminal for the changes to take effect."
echo "After restarting, you can try running 'snd --help'. :)"
