# 00 — Setup & Toolchain Tour

> **Goal:** make sure your toolchain works, get a feel for the five commands you'll
> run a thousand times, and understand the shape of a Go module.

## What you'll learn

- The five commands: `go run`, `go build`, `go test`, `go fmt`, `go vet`
- What `go.mod` is and why every project has one
- How packages and import paths line up with directories

## Coming from Python

| Python | Go | Why it matters |
|---|---|---|
| `python script.py` | `go run ./pkg` | `go run` compiles + runs in one shot. No `__main__`. |
| `pip install foo` → `requirements.txt` | `go get foo` → `go.mod` + `go.sum` | `go.mod` lists deps; `go.sum` pins their checksums. |
| Virtualenv | Module per repo | The repo *is* the virtualenv. The `go.mod` at the root scopes it. |
| `pytest` | `go test` (built in) | Testing is part of the toolchain, not a third-party library. |
| `black` / `ruff format` | `gofmt` (built in) | Formatting is **not negotiable** in Go. There is one style. Run it. |
| `mypy` | `go vet` (built in) + the compiler | Compiler is strict; `go vet` catches additional bugs (printf mismatches, etc.). |

## Concept walkthrough

A Go module is a directory tree rooted at a `go.mod` file. The `module` line sets
the import path. Subdirectories are packages, named by their `package` clause:

```go
// 00-setup/exercises.go
package setup

func Greeting() string {
    return "go lab is alive"
}
```

Anywhere in the module, `import "example.com/golab/00-setup"` would pull this in
(though we won't import it — each chapter is its own world).

The toolchain commands you need today:

```sh
go version          # what's installed
go mod tidy         # add missing / drop unused deps in go.mod
go build ./...      # compile every package; ./... means "this dir + recursively"
go test ./...       # run every test
go fmt ./...        # reformat in place
go vet ./...        # static checks
go run ./00-setup   # compile-and-run a main package (this chapter has no main)
```

## Your turn

Open [`exercises.go`](./exercises.go). Implement `Greeting()` so the test passes.

```sh
go test ./00-setup
```

Green? Run the scoreboard:

```sh
./scripts/progress.sh
```
