# PingOne Terraform Provider

Terraform provider to support PingOne IAC configuration.

## Requirements

- Terraform 1.0+
- Go 1.18

# Adding new stuff

If you need to create a new resource create a folder under `internal/` following the folder structure below.

```
terraform-provider-pingone
│ README.md
│ main.go
└─internal
│ └─things
│   │ expander.go
│   │ flattener.go [REQUIRED]
│   │ sources.go
│   │ resources.go
```

After defining these resources you must add new block(s) to [provider.go](internal/provider/provider.go).

```go
import {
  ...

  "github.com/ronniehicks/terraform-provider-pingone/internal/thing"
}

func Provider() *schema.Provider {
  Schema: ...,
  DataSourcesMap: map[string]*schema.Resource{
    "new_things": thing.DataSource,
    ...
  },
  ResourcesMap: map[string]*schema.Resource{
    "new_thing": thing.Resource,
  },
}
```

- `expander.go` - contains methods to "expand" from TF resources to provider specific object structures (PingOne API)
  > To be used with `resources.go`
- `flattener.go` - contains methods to "flatten" from provider specific object structures (PingOne API) to TF state definitions
  > To be used with `resources.go` and `sources.go`
- `sources.go` (optional) - entrypoint for TF defined "data sources" and will always contain references to the flattener
  > Note: This maps to the `data` keyword in a \*.tf file
- `resources.go` (optional) - entrypoint for TF defined "resources" and will always contain references to the expander and flattener
  > Note: This maps to the `resource` keyword in a \*.tf file

# Build and Test

To debug run the "Debug Terraform Provider" configuration provided. This will create an output like below. Copy the bottom line of that output and paste it before your `terraform` command.

```shell
Detaching and terminating target process
dlv dap (59345) exited with code: 0
Starting: /go/bin/dlv dap --check-go-version=false --listen=127.0.0.1:39257 --log-dest=3 from /workspaces/terraform-provider-pingone
DAP server listening at: 127.0.0.1:39257
Type 'dlv help' for list of commands.
{"@level":"debug","@message":"plugin address","@timestamp":"2022-05-12T21:56:30.343059Z","address":"/tmp/plugin810917772","network":"unix"}
Provider started. To attach Terraform CLI, set the TF_REATTACH_PROVIDERS environment variable with the following:

	TF_REATTACH_PROVIDERS='{"registry.terraform.io/ronniehicks/pingone":{"Protocol":"grpc","ProtocolVersion":5,"Pid":60925,"Test":true,"Addr":{"Network":"unix","String":"/tmp/plugin810917772"}}}'

```

## Example

1. F5/Debug
2. Copy `TF_REATTACH_PROVIDERS` statement
3. Run this command in a terminal window from the `examples/` folder

- `TF_REATTACH_PROVIDERS='{"registry.terraform.io/ronniehicks/pingone":{"Protocol":"grpc","ProtocolVersion":5,"Pid":60925,"Test":true,"Addr":{"Network":"unix","String":"/tmp/plugin810917772"}}}' terraform apply`
