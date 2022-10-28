## integrationcli integrations executions suspensions list

List all suspensions of an integration

### Synopsis

List all suspensions of an integration

```
integrationcli integrations executions suspensions list [flags]
```

### Options

```
      --filter string      Filter results
  -h, --help               help for list
  -n, --name string        Integration flow name
      --orderBy string     The results would be returned in order
      --pageSize int       The maximum number of versions to return (default -1)
      --pageToken string   A page token, received from a previous call
```

### Options inherited from parent commands

```
  -a, --account string       Path Service Account private key in JSON
      --apigee-integration   Use Apigee Integration; default is false (Application Integration)
      --disable-check        Disable check for newer versions
      --no-output            Disable printing API responses from the control plane
  -p, --proj string          Integration GCP Project name
  -r, --reg string           Integration region name
  -t, --token string         Google OAuth Token
```

### SEE ALSO

* [integrationcli integrations executions suspensions](integrationcli_integrations_executions_suspensions.md)	 - Manage suspensions of an integrations flow version

###### Auto generated by spf13/cobra on 28-Oct-2022