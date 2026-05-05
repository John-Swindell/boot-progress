# 04 — Structs & Methods

> **Goal:** model data with structs, attach behavior with methods, and *correctly
> choose value vs pointer receivers* — the #1 thing newcomers get wrong.

## What you'll learn

- Defining structs and constructing them
- Methods: value receivers vs pointer receivers
- When pointer receivers are required (mutation, large structs, interface satisfaction)
- Struct embedding: composition over inheritance
- Field tags (preview — used heavily in chapter 09 for JSON)

## Coming from Python

| Python | Go | Note |
|---|---|---|
| `class Counter: ...` | `type Counter struct { ... }` | Data is in struct fields. |
| `def inc(self):` | `func (c *Counter) Inc() { ... }` | The receiver replaces `self` and goes *before* the method name. |
| `__init__` | Just construct: `&Counter{Limit: 100}` | No `__init__`; sometimes a `NewCounter(...)` factory function. |
| `class Logging(Counter):` | `type Logging struct { *Counter; ... }` | **Embedding**, not inheritance. |
| `super().inc()` | `c.Counter.Inc()` (the embedded field is named after the type) | No `super`. |

## Value vs pointer receivers — the rules

```go
func (c Counter) Value() int  { return c.n }   // value receiver — copy
func (c *Counter) Inc()       { c.n++ }        // pointer receiver — mutates
```

**Use a pointer receiver when:**
1. You need to **mutate** the receiver.
2. The struct is **large** (avoid copy cost).
3. The type contains a `sync.Mutex` or similar — copying breaks them.
4. You need the type to satisfy an interface only via pointer methods.

**Use a value receiver when:**
- The type is small and immutable-ish (e.g. `time.Time`, a `Point`).

**Be consistent within a type.** Don't mix value and pointer receivers on the same
type — it confuses readers and the rules around method sets get subtle.

## Embedding (composition)

```go
type Counter struct{ n int }
func (c *Counter) Inc()        { c.n++ }
func (c *Counter) Value() int  { return c.n }

type AuditCounter struct {
    *Counter             // embedded: AuditCounter "has-a" Counter
    Audit []time.Time
}

func (a *AuditCounter) Inc() {
    a.Audit = append(a.Audit, time.Now())
    a.Counter.Inc()      // explicitly call the embedded method
}
```

`AuditCounter.Value()` is automatically promoted — you call `a.Value()` and it
forwards to `a.Counter.Value()`. This is *not* inheritance — `AuditCounter` is
not a subtype of `Counter`. It's just delegation with sugar.

## Field tags (teaser)

```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age,omitempty"`
}
```

Tags are string metadata read at runtime via reflection — `encoding/json` uses
them. We'll lean on this in chapter 09.

## Your turn

```sh
go test ./04-structs-methods
```
