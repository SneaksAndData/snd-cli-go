# Command-line interface for Sneaks & Data

This repository contains a command-line interface for internal and external services used in Sneaks & Data.

## Requirements

To be able to retrieve the installation script and the cli binary you need to have Azure CLI installed. Instructions on
how to install it can be found here https://learn.microsoft.com/en-us/cli/azure/install-azure-cli.

## Installation

```bash
# Login into Azure
az login
# Retrieve the install.sh script fro the blob storage
az storage blob download --blob-url https://esddatalakeproduction.blob.core.windows.net/dist/snd-cli-go/install.sh --auth-mode login --file "install.sh"
# Run the installation script
# Add execute permission
chmod +x install.sh
```

### Linux

```bash
./install.sh
```

### macOS

```bash
sh ./install.sh
```

### Windows
```bash
bash ./install.sh
```


## Uninstall

TBD

## Usage

CLI supports the following command groups: `login`, `claim`, `spark`, `algorithm`. Each command group is described in
respective
section below.

```bash
Usage:
  snd [command]

Auth Commands
  login       Get internal authorization token

Claim Commands
  claim       Manage claims

ML Algorithm Commands
  algorithm   Manage ML/AI algorithm runs and related resources

Spark Commands
  spark       Manage Spark jobs and related resources

Additional Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -h, --help   help for snd
```

### Login

Login in to SnD CLI. This command will retrieve your internal authorization token using the specified identity provider,
defaults to AzureAD. For now, we only support AzureAD as provider.

```bash
Usage:
  snd login [flags]

Flags:
  -a, --auth-provider string   Specify the OAuth provider name (default "azuread")
  -e, --env string             Target environment (default "test")
  -h, --help                   help for login

```

#### Examples

```bash
# Authenticate with AzureAD on test environment
$ snd login
$ snd login -a azuread -e test 
# Authenticate with AzureAD on production environment
$ snd login -e production
$ snd login -a azuread -e production
```

### Claim

Manage Boxer claims and users. This feature is available only for platform administrators.

```bash
Usage:
  snd claim [command]

Available Commands:
  add         Add a new claim to an existing user
  get         Retrieves claims assigned to an existing user
  remove      Removes a claim from an existing user
  user        Manage (add/remove) a user

Flags:
  -a, --auth-provider string     Specify the OAuth provider name (default "azuread")
      --claims-provider string   Specify the claim provider
  -e, --env string               Target environment (default "test")
  -h, --help                     help for claim
  -u, --user string              Specify the user ID
```

#### Get claims

```bash
$ snd claim get --claims-provider azuread -u test@ecco.com
$ snd claim get --claims-provider azuread --user test@ecco.com
```

#### Add claims

```bash
$ snd claim add --claims-provider azuread -u test@ecco.com -c "test1.test.sneaksanddata.com/.*:.*"
$ snd claim add --claims-provider azuread --user test@ecco.com --claims "test1.test.sneaksanddata.com/.*:.*"
```

#### Remove claims

```bash
$ snd claim remove --claims-provider azuread -u test@ecco.com -c "test1.test.sneaksanddata.com/.*:.*"
$ snd claim remove --claims-provider azuread --user test@ecco.com --claims "test1.test.sneaksanddata.com/.*:.*"
```

#### Manage users

##### Add user

```bash
$ snd claim user add --claims-provider azuread -u test@ecco.com 
$ snd claim user add --claims-provider azuread --user test@ecco.com 
```

##### Remove user

```bash
$ snd claim user remove --claims-provider azuread -u test@ecco.com 
$ snd claim user remove --claims-provider azuread --user test@ecco.com 
```

### Algorithm

Run and retrieve information about ML algorithm jobs.

```bash
Usage:
  snd algorithm [command]

Available Commands:
  get         Get the result for a ML Algorithm run
  run         Run a ML Algorithm

Flags:
      --algorithm string       Specify the algorithm name
  -a, --auth-provider string   Specify the OAuth provider name (default "azuread")
  -e, --env string             Target environment (default "test")
  -h, --help                   help for algorithm
```

#### Run algorithm

```bash
$ snd algorithm run --algorithm store-auto-replenishment-crystal-orchestrator -p ./crystal-payload.json
$ snd algorithm run --algorithm store-auto-replenishment-crystal-orchestrator --payload ./crystal-payload.json
```

#### Get algorithm job result

```bash
$ snd algorithm get -i fa1d02af-c294-4bf6-989f-1234 --algorithm store-auto-replenishment-crystal-orchestrator
$ snd algorithm get --id fa1d02af-c294-4bf6-989f-1234 --algorithm store-auto-replenishment-crystal-orchestrator
```

### Spark

Retrieve information related to spark jobs by using job ID.

```bash
Usage:
  snd spark [command]

Available Commands:
  configuration  Get a deployed SparkJob configuration
  encrypt        Encrypt a value from a file or stdin using encryption key from a corresponding Spark Runtime
  logs           Get logs from a Spark Job
  request-status Get the status of a Spark Job
  runtime-info   Get the runtime info of a Spark Job
  submit         Runs the provided Beast V3 job with optional overrides

Flags:
  -a, --auth-provider string   Specify the OAuth provider name (default "azuread")
  -e, --env string             Target environment (default "test")
  -h, --help                   help for spark
  -i, --id string              Specify the  Job ID

```

#### Request status

```bash
$ snd spark request-status -i 54284cb9-8e58-4d92-93cb-6543
```

#### Runtime info

```bash
$ snd spark runtime-info -i 54284cb9-8e58-4d92-93cb-6543
```

#### Request logs

```bash
$ snd spark logs -i 54284cb9-8e58-4d92-93cb-6543
```

#### Submit

```bash
$ snd spark submit -n configuration-name -o ./overrides.json
$ snd spark submit --job-name configuration-name --overrides ./overrides.json
```

overrides.json structure

```json
{
  "extra_arguments": {},
  "project_outputs": []
}
```

#### Get configuration

```bash
$ snd spark configuration -n configuration-name 
$ snd spark configuration --name configuration-name 
```

#### Encrypt

```bash
NOT YET SUPPORTED
```