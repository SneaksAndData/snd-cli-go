# Command-line interface for Sneaks & Data

This repository contains a command-line interface for internal and external services used in Sneaks & Data.
 
## Installation
Retrieve installation script from GitHub
```bash
curl -L -o install.sh https://github.com/SneaksAndData/snd-cli-go/releases/latest/download/install.sh
```

#### Grant execute permission
```bash
chmod +x install.sh
```

#### Run the installation script
##### Linux
```bash
./install.sh
```
##### macOS

```bash
sh ./install.sh
```


## Requirements (deprecated)

To be able to retrieve the installation script and the cli binary you need to have Azure CLI installed. Instructions on
how to install it can be found here https://learn.microsoft.com/en-us/cli/azure/install-azure-cli.

## Installation (deprecated)
#### Login into Azure
```bash
az login
```
#### Retrieve the install.sh script from the blob storage
```bash
az storage blob download --blob-url https://esddatalakeproduction.blob.core.windows.net/dist/snd-cli-go/install.sh --auth-mode login --file "install.sh"
```
#### Grant execute permission
```bash
chmod +x install.sh
```

#### Run the installation script
##### Linux
```bash
./install.sh
```
##### macOS

```bash
sh ./install.sh
```


## Uninstall

TBD

## Usage

See commands documentation [here](./docs/snd.md).
