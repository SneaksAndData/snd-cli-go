## snd login

Get internal authorization token

### Synopsis

Retrieve the internal authorization token generated by Boxer by providing an authentication provider

```
snd login [flags]
```

### Examples

```
$ snd login -a azuread -e test
$ snd login --auth-provider azuread --env production

```

### Options

```
  -a, --auth-provider k8s   Specify the OAuth provider name 
                            For in-cluster Kubernetes auth specify name of your kubernetes cluster context prefixed with k8s
                            for example `k8s-esd-airflow-dev-0` (default "azuread")
  -e, --env string          Target environment (default "test")
  -h, --help                help for login
```

### SEE ALSO

* [snd](snd.md)	 - SnD CLI

###### Auto generated by spf13/cobra on 29-Apr-2024