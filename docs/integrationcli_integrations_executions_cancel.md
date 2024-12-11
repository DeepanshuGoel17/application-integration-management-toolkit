## integrationcli integrations executions cancel

Cancels an execution of an integration

### Synopsis

Cancels an execution of an integration

```
integrationcli integrations executions cancel [flags]
```

### Options

```
  -c, --cancel-reason string   Cancel Reason
  -e, --execution-id string    Execution ID
  -h, --help                   help for cancel
  -n, --name string            Integration flow name
```

### Options inherited from parent commands

```
  -a, --account string   Path Service Account private key in JSON
      --api api          Sets the control plane API. Must be one of prod, staging or autopush; default is prod
      --default-token    Use Google default application credentials access token
      --disable-check    Disable check for newer versions
      --metadata-token   Metadata OAuth2 access token
      --no-output        Disable printing all statements to stdout
      --print-output     Control printing of info log statements (default true)
  -p, --proj string      Integration GCP Project name
  -r, --reg string       Integration region name
  -t, --token string     Google OAuth Token
      --verbose          Enable verbose output from integrationcli
```

### SEE ALSO

* [integrationcli integrations executions](integrationcli_integrations_executions.md)	 - Manage integration executions for an integration

###### Auto generated by spf13/cobra on 10-Dec-2024