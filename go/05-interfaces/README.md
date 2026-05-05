# 05 — Interfaces

> **Goal:** Use Go's most powerful (and simplest) abstraction. Interfaces in Go
> are *implicit* — types satisfy them by having the right method set, not by
> declaring intent.

## What you'll learn

- Defining and implementing interfaces (no `implements` keyword)
- The empty interface `any` (alias for `interface{}`) and when to use it
- Type assertions: `v, ok := x.(T)`
- Type switches: `switch v := x.(type)`
- Standard interfaces you'll lean on every day: `error`, `fmt.Stringer`, `io.Reader`, `io.Writer`

## Coming from Python

| Python | Go | Note |
|---|---|---|
| Duck typing at runtime | Static implicit interfaces | Compiler verifies the method set at compile time. |
| `class X(Protocol)` | `type X interface { ... }` | `Protocol` and Go interfaces are conceptually nearly identical. |
| `__str__` | `String() string` (satisfies `fmt.Stringer`) | `fmt` checks for it automatically. |
| `isinstance(x, T)` | `_, ok := x.(T)` (assertion) or type switch | Go has no `isinstance`, but assertions are the equivalent. |
| `raise NotImplementedError` | Don't — interfaces are method sets, you simply don't have a method to call | If a type doesn't implement an interface, it doesn't satisfy it. End of story. |

## The Big Idea

```go
type Notifier interface {
    Notify(msg string) error
}
```

Any type that has a `Notify(msg string) error` method **is** a `Notifier`. No
declaration, no inheritance, no registration. A logging notifier, an email
notifier, a Slack notifier, a no-op notifier — they all satisfy this interface
just by having the method.

This lets you write code that takes `Notifier` and works with any
implementation, including ones written *after* your code that you've never
seen. That's the whole point.

## Keep interfaces small

> "The bigger the interface, the weaker the abstraction." — Rob Pike

`io.Reader` is one method. `io.Writer` is one method. `error` is one method.
You compose larger interfaces from small ones (`io.ReadWriter` is just
`Reader` + `Writer`). Define interfaces **at the consumer side**, not on the
implementer's side.

## Type switches

```go
switch v := x.(type) {
case int:
    return "int"
case string:
    return "string of len " + strconv.Itoa(len(v))
case []byte:
    return "bytes"
default:
    return "unknown"
}
```

Each `case` binds `v` to the asserted type — you can use it directly without
re-asserting.

## Your turn

```sh
go test ./05-interfaces
```
