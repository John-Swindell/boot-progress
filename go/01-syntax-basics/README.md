# 01 — Syntax Basics

> **Goal:** get fluent with Go's variable declaration, types, control flow, and
> the fact that `for` is the only loop keyword.

## What you'll learn

- `var` vs `:=` (and when each is required)
- The numeric tower (`int`, `int64`, `float64`, `byte`, `rune`) and zero values
- `if` / `else` with the optional **init statement**
- `switch` (no `break` needed, multi-value cases)
- The four shapes of `for`

## Coming from Python

| Python | Go | Note |
|---|---|---|
| `x = 5` | `x := 5` (inside funcs) or `var x = 5` | `:=` is **declare + assign**; `=` is just assign. |
| Dynamic types | Static types — every var has one at compile time | `var x int` defaults to `0` (the zero value). |
| `True` / `False` | `true` / `false` | Lowercase. |
| `if 0:` is falsy | `if 0:` doesn't compile | You compare explicitly: `if x != 0`. |
| `while cond:` | `for cond { }` | No `while` — `for` does it. |
| `for x in range(10):` | `for i := 0; i < 10; i++ { }` | C-style three-clause for, or `for i := range slice`. |
| `if x in {1, 2, 3}:` | `switch x { case 1, 2, 3: }` | `switch` cases can list multiple values. |
| `elif` | `else if` | Two words. |

### The optional `if` init statement

This idiom shows up everywhere:

```go
if err := doThing(); err != nil {
    return err
}
// err is out of scope here
```

Declare a variable inside the `if`, scope it to the `if`/`else` block, gone after.
You'll use this constantly with errors.

### The four `for` shapes

```go
for i := 0; i < 10; i++ { }       // C-style
for x < 100 { }                    // while
for { break }                      // infinite
for i, v := range slice { }        // range
```

## Concept walkthrough

```go
package syntax

func Abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}
```

No parentheses around the condition. Braces are mandatory even for one-liners.

```go
func Grade(score int) string {
    switch {
    case score >= 90:
        return "A"
    case score >= 80:
        return "B"
    default:
        return "F"
    }
}
```

`switch` with no expression == chained `if/else if`. No `break` — cases don't
fall through (use `fallthrough` if you really want C semantics, you almost never
do).

## Your turn

Open [`exercises.go`](./exercises.go). Implement each function to spec.

```sh
go test ./01-syntax-basics
```
