# Day 1 — Types, Variables, Zero Values

**Date:** 2026-04-05

---

## What We Built

`day1/main.go` — A `Document` struct modelling a RAG pipeline chunk, used to explore Go's type system and variable declaration patterns.

---

## Concepts Covered

### Zero Values
Every Go variable is initialized to a safe default — no null, no uninitialized state.

| Type | Zero Value |
|---|---|
| `int` | `0` |
| `float64` | `0.0` |
| `string` | `""` |
| `bool` | `false` |
| struct | each field at its zero value |

| | Undefined field | Declared but uninitialized field |
|---|---|---|
| **Python** | Runtime `AttributeError` | Runtime `AttributeError` |
| **Java** | Compile error | Compiles, but field is `null` → NPE at runtime |
| **Go** | Compile error | Compiles, field is zero value → always safe |

```python
# Python case 1 — undefined field, blows up at runtime
class Document:
    pass

doc = Document()
print(doc.content)  # AttributeError: 'Document' object has no attribute 'content'

# Python case 2 — declared but uninitialized (set to None), blows up at runtime
class Document:
    def __init__(self):
        self.content = None  # declared, but None

doc = Document()
print(doc.content.upper())  # AttributeError: 'NoneType' object has no attribute 'upper'
```

```java
// Java — undefined field caught at compile time
// but declared field defaults to null → NPE at runtime
class Document {
    String content;  // declared but not initialized
}
Document doc = new Document();
System.out.println(doc.content.length()); // NullPointerException at runtime
```

```go
// Go — undefined field caught at compile time
// declared field always has a zero value → no crash
var doc Document
fmt.Println(doc.Content == "") // true — safe, no panic
```

**Go's guarantee:** The real win over Java isn't catching undefined fields (both do that) — it's that declared fields are never `null`. Zero values make the uninitialized state safe by design.

### Variable Declaration
- `var doc Document` — explicit declaration, gets zero value
- `title := "..."` — short declaration with type inference, only valid inside functions
- `var` works at package level; `:=` does not

### Struct Literals
- Set only the fields you care about
- Unset fields automatically get their zero values
- `%+v` format verb prints field names alongside values

### Constants
- Declared with `const`, evaluated at compile time
- No runtime cost

### Visibility (Access Control)
- **Uppercase** first letter = exported (public)
- **Lowercase** first letter = package-private
- No `public`, `private`, `protected` keywords — capitalization is the entire model

### Multiple Return Values
- Functions can return `(int, string)` natively
- In Java you'd need a wrapper object; in Python you'd use a tuple
- This is how Go handles errors — preview for Day 2

---

## Errors Hit

**Error:** `package command-line-arguments is not a main package`  
**Cause:** File had `package day1` instead of `package main`  
**Fix:** Every executable Go program must declare `package main`. Library code uses a descriptive package name.

---

## Key Insight

Zero values aren't just "happens to be zero" — they're a **language guarantee**. You can always safely compare, use, or pass a zero-value struct. In a production LLM gateway (Month 3), this means a missing request field gives you a predictable zero value to check — not a nil panic.

---

## What's Next — Day 2
Error handling: errors are return values, not exceptions. No stack unwinding, no invisible control flow. The compiler forces you to acknowledge them.

Build: `loadDocument(path string) (Document, error)`
