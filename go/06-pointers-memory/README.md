# 06 — Pointers & Memory

> **Goal:** know exactly when to take a pointer, when to dereference, and what
> "escape to the heap" means in practice.

## What you'll learn

- `&` (address-of) and `*` (dereference)
- Pointer vs value semantics in function calls
- `nil` pointers and how to safely check
- Escape analysis intuition — why you don't need to think about heap vs stack 99% of the time
- A taste of `unsafe.Sizeof` to understand layout (read-only)

## Coming from Python

Python has *no* explicit pointers. Every object is implicitly a reference, and
mutation happens via method calls. Go is more like C: you control whether you
pass a value (copy) or a pointer (alias).

| Python | Go | Note |
|---|---|---|
| All objects are references | Values copy by default; you take a pointer with `&` | `func(s string)` copies the string header (3 words). `func(s *string)` passes a pointer. |
| `x is None` | `x == nil` | Only meaningful for pointers, slices, maps, channels, funcs, interfaces. |
| `id(x)` | `&x` (the address) | Useful for understanding aliasing. |
| Mutating arg via methods | Pass a pointer or pass a slice/map | Slices and maps are already reference-y under the hood (they point to backing storage). |

## When to use a pointer

- **Mutation:** `func (c *Counter) Inc()` must take a pointer or it mutates a copy.
- **Large struct:** copying a 200-byte struct on every call is wasteful.
- **Optional value:** `*int` distinguishes "0" from "absent" (`nil`).
- **Identity matters:** if two callers must share *the same* object.

## When NOT to use a pointer

- For small read-only values (`int`, `time.Time`, small structs). Copies are cheap and avoid aliasing bugs.
- Just to "feel like Python." Don't.

## Escape analysis (intuition only)

The Go compiler decides whether each allocation lives on the **stack** (cheap,
auto-freed when function returns) or the **heap** (GC'd). It does this with
*escape analysis*: if a pointer to the value can outlive the function, it
escapes to the heap.

```go
func newPoint() *Point {
    p := Point{1, 2}   // would be stack-only...
    return &p          // ...but we return its address, so it escapes to heap
}
```

You usually don't need to think about this. But: returning pointers, passing
to interfaces, capturing in closures, sending on channels — these often force
heap allocation. If you ever need to look, run:

```sh
go build -gcflags='-m' ./05-interfaces 2>&1 | grep escapes
```

## Pointer pitfall — capturing the loop variable (pre-Go 1.22)

In Go 1.21 and earlier, this bug bit everyone:

```go
for _, x := range xs {
    go func() { fmt.Println(x) }()   // BUG: all goroutines see the same x
}
```

In Go 1.22+ (this lab's toolchain) the loop var is per-iteration, so this is
fixed. But you'll see the workaround in older code:

```go
for _, x := range xs {
    x := x  // shadow into a fresh per-iter variable
    go func() { fmt.Println(x) }()
}
```

## Your turn

```sh
go test ./06-pointers-memory
```
