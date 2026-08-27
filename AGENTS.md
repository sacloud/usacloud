# Usacloud Agent Instructions

## Project overview

`usacloud` is the official CLI for Sakura Cloud. `main.go` delegates to `pkg.Run()`, which registers Cobra commands and executes the root command.

- `pkg/commands/root` defines persistent configuration flags and the root Cobra command.
- `pkg/resources.go` registers top-level resources.
- IaaS resources are collected in `pkg/commands/iaas/resources.go`.
- They are attached beneath `usacloud iaas` and exposed as hidden root-level commands for backward compatibility.
- Each resource package under `pkg/commands/<platform>/<resource>` defines a `core.Resource` and registers `core.Command` values from `init()`. Parameter structs combine shared `cflag` types with resource-specific fields.
- `pkg/core` drives the common command lifecycle.
- It loads configuration and profiles, initializes clients and output, and validates parameters.
- It also expands selectors across zones, confirms destructive operations, invokes commands, and prints results.
- Commands without a custom `core.Command.Func` use generated adapters in `pkg/services/<platform>/*_services_gen.go` to call the matching service library. The generated `zz_*_gen.go` files in command packages create Cobra flags from parameter struct tags.
- `tools/gen-commands` reads registered resources and commands to generate those adapters and flag files. Do not manually edit files ending in `_gen.go` or marked `Code generated`; change the resource/command definition or generator template, then regenerate.

## Commands

Use Go `1.25.8` (from `go.mod`).

| Purpose | Command |
| --- | --- |
| Install development tools | `make tools` |
| Build the CLI | `make build` |
| Regenerate command/service adapters, licenses, and formatting | `make gen` |
| Run all unit tests | `make test` |
| Run one package's tests | `go test ./pkg/config -run '^TestFillDefaults_ZonesEmpty$' -v` |
| Run one test case in a command package | `go test ./pkg/commands/iaas/disk -run '^TestCreate_ConvertToServiceRequest$' -v` |
| Run end-to-end tests | `make e2e-test` |
| Run one end-to-end test | `make e2e-test TESTARGS='-run ^TestE2E_minimum$'` |
| Go lint and formatting | `make lint-go` |
| Text lint | `make lint-text` |
| GitHub Actions workflow lint | `make lint-action` |

`make test` runs `go test ./...` with the race detector. End-to-end tests have the `e2e` build tag, install the local CLI first, and can take up to 240 minutes.

## Command implementation conventions

- Define CLI flags and service request mapping declaratively on parameter structs. Embedded shared parameters normally use both `cli:",squash"` and `mapconv:",squash"`; fields that must not reach service requests use `mapconv:"-"`.
- Use `validate` tags for standard validation. Add named value mappings in `pkg/vdef/definitions.go` when a flag needs `options=...`, filtering, or key/value conversion. Put validation that depends on several fields in the command's custom validator.
- Set `SelectorTypeRequireSingle` or `SelectorTypeRequireMulti` for commands that accept resource ID, partial name, or tag arguments.
- The core handles lookup and per-zone execution. Account for `--zone=all` and use `ctx.WithResource` rather than duplicating selector logic.
- Keep a command's `Category` and `Order` consistent with the shared categories. `Resource.AddCommand` fails for an unknown category and sorts help output by these values.
- Add new IaaS resources to `pkg/commands/iaas/resources.go`; add non-IaaS top-level resources to `pkg/resources.go`. Set `ServiceType` to the corresponding service-library type so generation can create adapters.
- Prefer the generated default service call where the service method matches the command. Provide `ServiceFuncAltName` for naming mismatches; use a custom `Func` only for behavior the service adapter cannot express.
- Commands that change resources should embed `cflag.ConfirmParameter`; noninteractive callers then require `--assumeyes`/`-y`. List commands conventionally set `NoProgress: true` and define explicit `ColumnDefs`.

## Repository conventions

- Preserve the Apache license header and the repository copyright year pattern in new Go files. `make gen` runs `set-license`, `gofmt`, and `gosimports`.
- `golangci-lint` deliberately excludes generated files and runs with `--fix`; review the resulting formatting changes before committing.
- Unit tests use the standard `testing` package with `testify/require`. End-to-end tests are isolated under `e2e/` behind the `e2e` build tag.
