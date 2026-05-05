# 02 — Functions & Errors

> **Goal:** Internalize Go's error-as-value philosophy. You will write
> `if err != nil { return err }` more times than you've eaten lunch. Make peace.

## What you'll learn

- Multiple return values, named returns, blank identifier
- `defer` (and how it stacks LIFO)
- The `error` interface and the `if err != nil` idiom
- Wrapping errors with `%w` and `fmt.Errorf`
- Sentinel errors, custom error types, `errors.Is` / `errors.As`
- `panic` and `recover` (and when not to use them — almost always)

## Coming from Python

| Python | Go | Note |
|---|---|---|
| `raise ValueError("bad")` | `return fmt.Errorf("bad")` | Errors are *returned*, not thrown. |
| `try: ... except X as e:` | `if errors.Is(err, X)` / `errors.As(err, &target)` | No try/except. Inspect the value. |
| Tuples | Multi-return: `func f() (int, error)` | Native syntax, no need to pack. |
| `with open(...) as f:` | `f, err := os.Open(...); defer f.Close()` | `defer` schedules cleanup at function exit. |
| `raise ... from e` | `fmt.Errorf("...: %w", e)` | `%w` preserves the wrapped chain for `errors.Is/As`. |
| Exceptions for control flow | Don't | `panic` is for *unrecoverable* program bugs. Not flow control. |

## The error interface

```go
type error interface {
    Error() string
}
```

That's it. Anything with an `Error() string` method is an error. So you can
build domain-specific errors as structs:

```go
type ValidationError struct {
    Field string
}

func (e *ValidationError) Error() string {
    return "invalid: " + e.Field
}
```

And the caller can recover the structured info:

```go
var verr *ValidationError
if errors.As(err, &verr) {
    log.Printf("bad field: %s", verr.Field)
}
```

## Wrapping

When you catch an error and re-return, **wrap** it with `%w` so the caller can
still `errors.Is` it:

```go
if err := openConfig(); err != nil {
    return fmt.Errorf("startup: %w", err)
}
```

`%w` preserves the chain. `%v` and `%s` flatten to a string and lose it.

## defer

```go
func write(path string, data []byte) error {
    f, err := os.Create(path)
    if err != nil {
        return err
    }
    defer f.Close()    // runs when write() returns, even on error
    _, err = f.Write(data)
    return err
}
```

Multiple defers run in **reverse order**. Use them for `Close()`, `Unlock()`,
`cancel()`, recovery handlers.

## panic / recover

`panic` unwinds the stack. `recover` (only meaningful inside a `defer`) stops
the unwind. Use them only at HTTP/RPC handler boundaries to convert "the
program is on fire" into a 500 instead of a crash. **Never** as an exception
substitute.

## Your turn

Open [`exercises.go`](./exercises.go).

```sh
go test ./02-functions-errors
```
