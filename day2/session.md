# Day 2 — Error Handling

**Date:** 2026-04-08

---

## What We Built

`day2/main.go` — `loadDocument(path string) (Document, error)` that reads a file, wraps OS errors with context, and returns a zero-value Document on failure.

---

## Concepts Covered

### Errors as Return Values

| | Java | Python | Go |
|---|---|---|---|
| Mechanism | `try/catch/throw` | `try/except/raise` | `(result, error)` return value |
| Ignoring errors | Possible | Possible | Compiler forces acknowledgment |
| Control flow | Invisible — exception unwinds the stack | Invisible — exception unwinds the stack | Explicit — just an `if` after the call |

**Why Go made this choice:** Errors in distributed systems need to be handled at the call site, not caught somewhere upstream. Invisible control flow (exceptions) makes it hard to reason about what happens when things fail. Go makes failure as visible as success.

### The Core Pattern

```go
func loadDocument(path string) (Document, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return Document{}, fmt.Errorf("loadDocument: %w", err)
    }
    return Document{Content: string(data)}, nil
}

doc, err := loadDocument("file.txt")
if err != nil {
    // handle
}
```

### Error Wrapping — `%w` vs `%v`

- `fmt.Errorf("loadDocument: %w", err)` — wraps the original error; callers can inspect the cause with `errors.Is()` / `errors.As()`
- `fmt.Errorf("loadDocument: %v", err)` — formats as string only; original error is lost
- Convention: each layer prefixes its name → `loadDocument: open file.txt: no such file or directory`

### `nil` for Errors

`error` is an interface in Go. `nil` is its zero value — meaning "no error." On the happy path, return `nil` as the error.

### `:=` vs `=`

- `:=` — declares **and** assigns; requires at least one new variable on the left
- `=` — assigns only; variable must already be declared

```go
doc, err := loadDocument("a.txt")  // first call — both new, use :=
doc, err = loadDocument("b.txt")   // second call — both exist, use =
```

### Blank Identifier `_`

`_` explicitly discards a return value. Using it on an error is valid Go but a code smell — any reviewer will flag it.

```go
doc, _ = loadDocument("missing.txt")
fmt.Println(doc.Content) // prints "" — zero value, no crash
```

**Why this is safe (but wrong):** Zero values + explicit errors work together. Even when you ignore an error, you get a safe zero-value struct back — not a nil panic. But in production, always handle the error.

---

## Key Insight

Go's error handling isn't just a style choice — it's an architectural one. When you see `if err != nil` after every call, the control flow is completely transparent. No hidden exception paths, no wondering "what could throw here." That explicitness is what makes Go services easy to operate in production.

---

## What's Next — Day 3

Structs, methods, interfaces. No inheritance. No `implements` keyword. Structural interfaces — if your struct has the right methods, it satisfies the interface automatically.

Build: `Retriever` interface → implement with `VectorRetriever` and `KeywordRetriever`.
