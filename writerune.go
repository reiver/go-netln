package netln

import (
	"sourcecode.social/reiver/go-utf8"

	"fmt"
	"io"
)

// writeRune deals with writing a single rune.
//
// It also makes sure that any error condition is represented as a Go error.
//
// It also wraps any errors, and provides a more appropriate error message.
func writeRune(writer io.Writer, r rune, expectedWritten int) error {

	n, err := utf8.WriteRune(writer, r)
	if nil != err {
		return fmt.Errorf("problem writing UTF-8 character %U: %w", r, err)
	}
	if expectedWritten != n {
		return fmt.Errorf("problem writing UTF-8 character %U — expected to write %d bytes, but actually wrote %b bytes", r, expectedWritten, n)
	}

	return nil
}
