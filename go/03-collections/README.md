# 03 — Slices, Arrays, Maps

> **Goal:** master slices (Go's most-used and most-bug-causing type) and maps,
> and understand the aliasing trap that snags every Pythonist.

## What you'll learn

- Arrays (fixed length, value semantics) vs slices (the thing you actually use)
- The slice header: pointer, len, cap — and why it matters
- The slice aliasing trap (the bug that *will* find you)
- Maps: declaration, the comma-ok idiom, randomized iteration order
- `append`, `copy`, `make`, the difference between `nil` and empty
- `range` semantics

## Coming from Python

| Python | Go | Note |
|---|---|---|
| `xs = []` | `var xs []int` (nil) or `xs := []int{}` (empty) | Both have `len == 0`. `nil` slice is fine to `append` to. |
| `xs.append(x)` | `xs = append(xs, x)` | `append` returns a new header — must reassign. |
| `xs[1:4]` | `xs[1:4]` | Looks the same. Behaves *very* differently — see Aliasing below. |
| `xs.copy()` | `dst := make([]int, len(xs)); copy(dst, xs)` | No method — explicit allocate + copy. |
| `d = {}` | `d := map[string]int{}` or `make(map[string]int)` | Always make/init before use. Reading a `nil` map is OK; writing panics. |
| `d.get(k, 0)` | `v := d[k]` (zero value if missing) | Maps return the zero value, never raise. |
| `if k in d:` | `v, ok := d[k]; if ok { ... }` | "comma-ok" idiom — also works for type assertions and channel recv. |
| `for k, v in d.items():` | `for k, v := range m { }` | **Order is randomized.** Sort keys yourself if you need determinism. |

## The Aliasing Trap

A slice is a *view* into a backing array. Taking a sub-slice does not copy:

```go
a := []int{1, 2, 3, 4, 5}
b := a[1:4]          // b == [2 3 4]
b[0] = 99            // mutates a too: a == [1 99 3 4 5]
```

When in doubt, `copy()` to detach. Especially when returning sub-slices from
internal data — callers can mutate your invariants.

`append` is even sneakier: if `cap(xs) > len(xs)`, append mutates the backing
array in place. If not, it allocates a new one. Two callers holding overlapping
slice headers can see different things after an `append`. **Rule:** if you
share slices, document who owns them.

## `make` and zero values

```go
var s []int           // nil, len 0, cap 0 — append-safe
s := []int{}          // empty (not nil), len 0, cap 0
s := make([]int, 5)   // len 5, all zeros, cap 5
s := make([]int, 0, 100) // len 0, cap 100 — preallocated
```

Preallocate (`make([]T, 0, n)`) when you know the size — avoids `O(log n)`
reallocs as you `append`.

## Your turn

```sh
go test ./03-collections
```
