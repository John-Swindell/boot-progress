# 12 — Testing: tables, subtests, benchmarks, fuzzing

> **Goal:** become fluent with Go's built-in testing toolkit. Idiomatic Go code
> ships with idiomatic tests; both look identical across companies because
> there's only one testing library.

## What you'll learn

- Naming + layout: `_test.go` suffix, same package or `_test` package
- Table-driven tests + `t.Run` subtests
- `t.Helper()`, `t.Fatalf` vs `t.Errorf`
- `*testing.B` benchmarks (and what they actually measure)
- `*testing.F` fuzzing — Go's killer feature for catching weird inputs
- Useful flags: `-run`, `-race`, `-bench`, `-count`, `-cover`

## Coming from Python

| pytest | Go testing | Note |
|---|---|---|
| `pytest` | `go test` | Built in. No fixtures library. |
| `@pytest.mark.parametrize` | Table-driven test (slice of structs + range loop) | Pure Go, no metaclass tricks. |
| `pytest.fail` | `t.Fatalf` (stops test) or `t.Errorf` (continues) | Prefer `Errorf` so you see all failures. |
| Fixtures | Plain helper funcs + `t.Helper()` | If a helper calls `t.Fatal`, mark `t.Helper()` so the line points at the caller. |
| `pytest-benchmark` | `go test -bench=.` | Built in. Output is reps/sec + ns/op. |
| `hypothesis` | `*testing.F` fuzzing | Built in since Go 1.18. Run with `go test -fuzz=Fuzz`. |

## Table-driven test idiom

```go
func TestReverse(t *testing.T) {
    cases := []struct {
        name, in, want string
    }{
        {"empty", "", ""},
        {"single", "a", "a"},
        {"ascii", "hello", "olleh"},
        {"emoji", "go👍", "👍og"},
    }
    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            if got := Reverse(tc.in); got != tc.want {
                t.Errorf("Reverse(%q) = %q, want %q", tc.in, got, tc.want)
            }
        })
    }
}
```

`t.Run` gives you per-case isolation, parallel execution (`t.Parallel()` inside),
and friendly names you can target with `-run TestReverse/emoji`.

## Benchmarks

```go
func BenchmarkReverse(b *testing.B) {
    s := "the quick brown fox"
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = Reverse(s)
    }
}
```

Run with `go test -bench=. -benchmem ./12-testing`. Use `-count=10` for stability.

## Fuzzing — find inputs you didn't think of

```go
func FuzzReverseRoundtrip(f *testing.F) {
    f.Add("hello")
    f.Add("👍")
    f.Fuzz(func(t *testing.T, s string) {
        if Reverse(Reverse(s)) != s {
            t.Errorf("roundtrip failed for %q", s)
        }
    })
}
```

`go test -fuzz=Fuzz` runs forever (or `-fuzztime=30s`), generating mutations
based on your seed corpus. When it finds a failure, it adds the input to
`testdata/fuzz/...` so it becomes a regression test.

## Your turn

```sh
go test ./12-testing
go test -bench=. -benchmem ./12-testing
go test -fuzz=FuzzReverseRoundtrip -fuzztime=5s ./12-testing
```
