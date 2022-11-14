# go-netln

**go-netln** provides tools for parsing "net lines", for the Go programming language —
i.e., lines that end with a "\r\n".

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-netln

[![GoDoc](https://godoc.org/github.com/reiver/go-netln?status.svg)](https://godoc.org/github.com/reiver/go-netln)

## Examples

```go
import "github.com/reiver/go-netln"

// ...

var storage strings.Builder
var writer io.Writer = &storage

netln.CopyLine(writer, reader)

line := storage.String()
```
