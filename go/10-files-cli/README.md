# 10 — Files & CLI

> **Goal:** read files efficiently, write them safely, parse command-line flags
> like a grown-up. Build small Unix-shaped tools.

## What you'll learn

- `os.Open` / `os.Create` / `os.WriteFile` / `os.ReadFile`
- `bufio.Scanner` for line-by-line reads
- `io.Reader` / `io.Writer` — the universal interfaces
- The `flag` package: `flag.String`, `flag.Int`, `flag.Bool`, `flag.Parse`
- Atomic writes via `os.Rename` (the only sane way to update a file)

## Coming from Python

| Python | Go | Note |
|---|---|---|
| `open(path)` | `os.Open(path)` returns `(*os.File, error)` | Both implement `io.Reader`. |
| `with open(...) as f:` | `f, err := os.Open(...); defer f.Close()` | |
| `for line in f:` | `s := bufio.NewScanner(f); for s.Scan() { line := s.Text() }` | Check `s.Err()` after the loop. |
| `f.read()` (whole file) | `os.ReadFile(path)` returns `([]byte, error)` | |
| `argparse` | `flag` (stdlib) or `cobra` (3rd-party for big CLIs) | `flag` is enough for ~80% of tools. |
| `os.replace(tmp, path)` | `os.Rename(tmp, path)` | Atomic on the same filesystem. |

## bufio.Scanner — the standard way to read lines

```go
f, err := os.Open(path)
if err != nil { return err }
defer f.Close()

s := bufio.NewScanner(f)
for s.Scan() {
    line := s.Text()  // current line, no trailing \n
    // ...
}
if err := s.Err(); err != nil {   // ALWAYS check after the loop
    return err
}
```

Default max line is 64KiB. For larger lines use `s.Buffer(make([]byte, 0, 1024*1024), 1024*1024)`.

## Atomic writes — the rename trick

`os.WriteFile` is not atomic — a crash mid-write leaves a partial file. The
classic fix:

```go
tmp := path + ".tmp"
if err := os.WriteFile(tmp, data, 0o644); err != nil {
    return err
}
if err := os.Rename(tmp, path); err != nil {
    os.Remove(tmp)
    return err
}
```

`Rename` on the same filesystem is atomic — the file at `path` either points
at the old data or the new data, never half-written. Every config-file writer
in production code does this.

## flag package basics

```go
var (
    addr    = flag.String("addr", ":8080", "listen address")
    workers = flag.Int("workers", 4, "concurrent workers")
    debug   = flag.Bool("debug", false, "enable debug logging")
)

func main() {
    flag.Parse()
    log.Printf("addr=%s workers=%d debug=%v", *addr, *workers, *debug)
    // positional args: flag.Args()
}
```

For exercises we'll bypass globals and use a `*flag.FlagSet` so we can test it.

## Your turn

```sh
go test ./10-files-cli
```
