# 13 — Modules, Builds, & Tooling

> **Goal:** speak fluent `go.mod`, stamp version metadata into your binaries
> with `-ldflags`, and learn the small set of community tools every Go shop
> uses.

## What you'll learn

- `go.mod`, `go.sum`, the module graph
- Common workflow: `go mod tidy`, `go mod download`, `go mod vendor`
- The `replace` directive (for local development)
- `-ldflags` for compile-time variable injection
- `golangci-lint` — the standard meta-linter
- Cross-compiling: `GOOS=linux GOARCH=amd64 go build`

## go.mod basics

```
module example.com/golab

go 1.22

require (
    k8s.io/client-go v0.29.0
    gopkg.in/yaml.v3 v3.0.1
)
```

- `module`: import path that everything under this dir is rooted at.
- `go`: minimum toolchain version (also affects language features available).
- `require`: direct + indirect dependencies, version-pinned.

`go.sum` records cryptographic checksums for every module version — committed,
verified on `go mod download`.

## Common commands

```sh
go mod init example.com/foo       # new project
go get example.com/bar@v1.2.3      # add/upgrade a dep
go mod tidy                         # add missing, drop unused — run after editing imports
go mod download                     # populate the local module cache
go mod vendor                       # copy deps into ./vendor (rarely needed; air-gapped builds)
go list -m all                      # show full dep graph
```

## The `replace` directive — your friend during local dev

When you're hacking on `example.com/foo` and you need it to use your local
copy of `example.com/bar`:

```
// go.mod
replace example.com/bar => ../bar
```

Go now resolves `bar` from the local path. **Strip before you commit** unless
the replace is intentional (some monorepos keep them).

## Stamping versions with -ldflags

```go
package version

var (
    Version = "dev"
    Commit  = "none"
    Date    = "unknown"
)
```

```sh
go build -ldflags "-X 'example.com/golab/13-modules-tooling.Version=v1.2.3' \
                   -X 'example.com/golab/13-modules-tooling.Commit=$(git rev-parse --short HEAD)' \
                   -X 'example.com/golab/13-modules-tooling.Date=$(date -u +%Y-%m-%dT%H:%M:%SZ)'" \
   ./...
```

`-X importpath.var=value` overwrites a string variable at link time. Use it for
version, commit, build date, environment.

## Cross compile

```sh
GOOS=linux   GOARCH=amd64  go build -o myapp-linux-amd64 ./cmd/myapp
GOOS=darwin  GOARCH=arm64  go build -o myapp-darwin-arm64 ./cmd/myapp
GOOS=windows GOARCH=amd64  go build -o myapp.exe ./cmd/myapp
```

No toolchain swap needed — Go ships every cross-target out of the box.

## golangci-lint

Install once: `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`.
Run: `golangci-lint run ./...`. It bundles `gofmt`, `govet`, `staticcheck`,
`errcheck`, and ~20 others. Most teams configure it via `.golangci.yml`. Catches
bugs `go vet` doesn't (especially around error handling and unused returns).

## Your turn

This chapter is more reading than coding — but there's one tiny exercise so
the scoreboard ticks green. See [`exercises.go`](./exercises.go).

```sh
go test ./13-modules-tooling
```

Bonus: run `go build -ldflags "-X example.com/golab/13-modules-tooling.Version=v0.1.0" ./13-modules-tooling/cmd/buildinfo` and inspect.
