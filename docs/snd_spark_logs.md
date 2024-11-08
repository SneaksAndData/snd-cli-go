## snd spark logs

Get logs from a Spark Job

```
snd spark logs [flags]
```

### Examples

```
snd spark logs --id 14abbec-e517-4135-bf01-fc041a4e
```

### Options

```
  -h, --help        help for logs
  -t, --trim-logs   Trims log to anything after STDOUT
```

### Options inherited from parent commands

```
  -a, --auth-provider string        Specify the OAuth provider name (default "azuread")
      --custom-auth-url string      Specify the auth service uri
      --custom-service-url string   Specify the service url (default "https://beast.%s.sneaksanddata.com")
  -e, --env string                  Target environment (default "awsd")
      --gen-docs                    Generate Markdown documentation for all commands
  -i, --id string                   Specify the  Job ID
```

### SEE ALSO

* [snd spark](snd_spark.md)	 - Manage Spark jobs

###### Auto generated by spf13/cobra on 4-Nov-2024
