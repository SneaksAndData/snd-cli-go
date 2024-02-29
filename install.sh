#!/bin/bash
# Installation script for SnD CLI
echo "Determining target OS and architecture..."
ARCH=$(uname -m)
if [[ "$OSTYPE" =~ ^darwin  &&  "$ARCH" =~ arm64 ]]; then
    bin_name=snd-darwin-arm64
elif [[ "$OSTYPE" =~ ^linux ]]; then
  if [[ "$ARCH" =~ arm64 ]]; then
    bin_name=snd-linux-arm64
  elif [[ "$ARCH" =~ amd64 ]]; then
    bin_name=snd-linux-amd64
elif [[ "$OSTYPE" =~ ^linux &&  "$ARCH" =~ amd64 ]]; then
   bin_name=snd-windows-amd64
  fi
else
    echo "Error: Unsupported OS type or architecture: $OSTYPE $ARCH "
    exit 1
fi
echo "Target OS: $OSTYPE"
echo "Target ARCH: $ARCH"

BUNDLE_URL="https://esddatalakeproduction.blob.core.windows.net/dist/snd-cli-go/$bin_name"

# Define base path for the application
base_path="$HOME/.local/snd-cli"
mkdir -p "$base_path"

# Check if az cli is installed
echo "Check if Azure CLI is installed..."
if ! command -v az &> /dev/null
then
    echo "az cli is not installed. Please install it from https://docs.microsoft.com/en-us/cli/azure/install-azure-cli"
    exit
fi

# Login into azure
echo "Please log in to Azure..."
az login

echo "Downloading the binary from $BUNDLE_URL"
# Get file
az storage blob download --blob-url $BUNDLE_URL --auth-mode login --file "$base_path/$bin_name"

chmod +x "$base_path/$bin_name"

if [ -e "$HOME/.local/bin/snd" ]; then
   echo "Removing symlink..."
   rm $HOME/.local/bin/snd
fi

# Create a symbolic link to the application
echo "Creating the symlink..."
ln -s "$base_path/$bin_name" "$HOME/.local/bin/snd"

echo "Please restart your terminal for the changes to take effect."
echo "After restarting, you can try running 'snd --help'. :)"