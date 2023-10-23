package netln

import (
	"sourcecode.social/reiver/go-utf8"

	"errors"
	"fmt"
	"io"
)

// readRune deals with reading a single rune.
//
// It also makes sure that any error condition is represented as a Go error,
// rather than a utf8.RuneError.
//
// It also wraps any errors, and provides a more appropriate error message.
func readRune(reader io.Reader) (r rune, size int, err error) {

	r, size, err = utf8.ReadRune(reader)
	if nil != err && !errors.Is(err, io.EOF) {
		err = fmt.Errorf("problem reading UTF-8 character: %w", err)
/////////////// RETURN
		return
	}
	if utf8.RuneError == r {
		err = fmt.Errorf("problem reading UTF-8 character: %w", errRuneError)
/////////////// RETURN
		return
	}

/////// RETURN
	return
}
