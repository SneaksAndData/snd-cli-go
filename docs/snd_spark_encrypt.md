## snd spark encrypt

Encrypt a value from a file or stdin using encryption key from a corresponding Spark Runtime

```
snd spark encrypt [flags]
```

### Options

```
  -h, --help                 help for encrypt
  -s, --secret-path string   Optional Vault secret path to Spark Runtime encryption key
  -v, --value string         Value to encrypt
```

### Options inherited from parent commands

```
  -a, --auth-provider string        Specify the OAuth provider name (default "azuread")
      --custom-service-url string   Specify the service url (default "https://beast-v3.%s.sneaksanddata.com")
  -e, --env string                  Target environment (default "test")
  -i, --id string                   Specify the  Job ID
```

### SEE ALSO

* [snd spark](snd_spark.md)	 - Manage Spark jobs

###### Auto generated by spf13/cobra on 29-Apr-2024