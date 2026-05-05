# 09 — Standard Library Tour

> **Goal:** get hands-on with the packages you'll use literally every day:
> `fmt`, `strings`, `strconv`, `time`, `encoding/json`.

## What you'll learn

- `fmt.Sprintf`, `fmt.Errorf` (and the `%w` verb)
- `strings.Fields`, `Split`, `TrimSpace`, `HasPrefix`, `Builder`
- `strconv.Itoa`, `Atoi`, `ParseInt`, `FormatFloat`
- `time.Time`, `time.Duration`, `time.Parse` (the cursed reference layout)
- `encoding/json` Marshal/Unmarshal with struct tags

## Coming from Python

| Python | Go | Note |
|---|---|---|
| `f"hello {name}"` | `fmt.Sprintf("hello %s", name)` | `%s`, `%d`, `%v` (any), `%q` (quoted), `%T` (type). |
| `str(42)` | `strconv.Itoa(42)` | Or `fmt.Sprint(42)` (slower, more general). |
| `int("42")` | `strconv.Atoi("42")` | Returns `(int, error)`. |
| `s.split()` | `strings.Fields(s)` (whitespace) or `strings.Split(s, ",")` | |
| `"".join(parts)` | `strings.Join(parts, "")` | |
| `"".join(...)` for hot loops | `strings.Builder` | Avoids quadratic copy. |
| `json.dumps(d)` | `json.Marshal(v)` | Returns `([]byte, error)`. |
| `json.loads(s)` | `json.Unmarshal(b, &v)` | **Pass a pointer** to fill. |
| `dict` round-trip | `map[string]any` | Ergonomic for unknown shapes. |

## Time formatting — the cursed layout

Go's `time.Format` and `time.Parse` use a **reference time** instead of
`%Y-%m-%d`-style codes:

```
Mon Jan 2 15:04:05 MST 2006
   = 1   2  3  4  5     6
```

`01/02 03:04:05PM '06 -0700` is the mnemonic. Common layouts are pre-defined:

```go
time.RFC3339         // 2006-01-02T15:04:05Z07:00
time.RFC3339Nano
time.DateTime        // 2006-01-02 15:04:05
```

Always parse and emit RFC3339 for logs/configs unless you have a reason not to.

## JSON struct tags

```go
type User struct {
    Name    string `json:"name"`
    Email   string `json:"email,omitempty"`   // omit when zero value
    Created time.Time `json:"created_at"`
    Secret  string `json:"-"`                 // never marshal this field
}
```

Default field names are the Go name (CapitalCase). Tags let you match the
JSON convention (snake_case usually). `omitempty` skips zero values. `"-"`
fully excludes.

## Your turn

```sh
go test ./09-stdlib-tour
```
