## snd spark

Manage Spark jobs

### Synopsis

Manage Spark jobs

### Examples

```
$ snd spark request-status --id 54284cb9-8e58-4d92-93cb-6543
$ snd spark runtime-info --id 54284cb9-8e58-4d92-93cb-6543
$ snd spark logs --id 54284cb9-8e58-4d92-93cb-6543
$ snd spark submit --job-name configuration-name --overrides ./overrides.json
$ snd spark configuration --name configuration-name 

```

### Options

```
  -a, --auth-provider string        Specify the OAuth provider name (default "azuread")
      --custom-service-url string   Specify the service url (default "https://beast-v3.%s.sneaksanddata.com")
  -e, --env string                  Target environment (default "test")
  -h, --help                        help for spark
  -i, --id string                   Specify the  Job ID
```

### SEE ALSO

* [snd](snd.md)	 - SnD CLI
* [snd spark configuration](snd_spark_configuration.md)	 - Get a deployed SparkJob configuration.

The name of the SparkJob should be provided as an argument.

* [snd spark encrypt](snd_spark_encrypt.md)	 - Encrypt a value from a file or stdin using encryption key from a corresponding Spark Runtime
* [snd spark logs](snd_spark_logs.md)	 - Get logs from a Spark Job
* [snd spark request-status](snd_spark_request-status.md)	 - Get the status of a Spark Job
* [snd spark runtime-info](snd_spark_runtime-info.md)	 - Get the runtime info of a Spark Job
* [snd spark submit](snd_spark_submit.md)	 - Runs the provided Beast V3 job with optional overrides

The overrides should be provided as a JSON file with the structure below.

If the 'clientTag' is not provided, a random tag will be generated.

If 'extraArguments', 'projectInputs', 'projectOutputs', or 'expectedParallelism' are not provided, the job will run with the default arguments.

<pre><code>
{
 "clientTag": "<string> (optional) - A tag for the client making the submission",
 "extraArguments": "<object> (optional) - Any additional arguments for the job",
 "projectInputs": [{
	"alias": "<string> (optional) - An alias for the input",
	"dataPath": "<string> (required) - The path to the input data",
	"dataFormat": "<string> (required) - The format of the input data"
	}
		// More input objects can be added here
	],
 "projectOutputs": [{
	"alias": "<string> (optional) - An alias for the output",
	"dataPath": "<string> (required) - The path where the output data should be stored",
	"dataFormat": "<string> (required) - The format of the output data"
	}
		// More output objects can be added here
	],
 "expectedParallelism": "<integer> (optional) - The expected level of parallelism for the job"
}
</code></pre>


###### Auto generated by spf13/cobra on 29-Apr-2024