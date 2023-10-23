# go-netln

**go-netln** provides tools for parsing "net lines", for the Go programming language —
i.e., lines that end with a "\r\n".

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/sourcecode.social/reiver/go-netln

[![GoDoc](https://godoc.org/sourcecode.social/reiver/go-netln?status.svg)](https://godoc.org/sourcecode.social/reiver/go-netln)

## Examples

```go
import "sourcecode.social/reiver/go-netln"

// ...

var storage strings.Builder
var writer io.Writer = &storage

netln.CopyLine(writer, reader)

line := storage.String()
```
