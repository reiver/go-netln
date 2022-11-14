package netln

import (
	"errors"
	"io"
)

// CopyLine reads in a line from ‘reader’ and writes that line to ‘writer’ (without the line terminator).
//
// For CopyLine, a line is terminated by the "\r\n" characters.
//
// So, for example, if from ‘reader’ CopyLine read in:
//
//	"Hello world!\r\n"
//
// Then what CopyLine would write to ‘writer’ is:
//
//	"Hello world!"
//
// (Notice that the "\r\n" at the end is missing.)
//
// Note that ‘n64’ represents how many bytes were written, not read.
func CopyLine(writer io.Writer, reader io.Reader) (n64 int64, err error) {

	for {
		var eof bool

		var r0 rune
		var size0 int
		{
			r0, size0, err = readRune(reader)
			eof = errors.Is(err, io.EOF)
			if nil != err && !eof {
/////////////////////////////// RETURN
				return
			}
		}

		if size0 <= 0 {
	/////////////// BREAK
			break
		}

		if '\r' != r0 || eof {

			err = writeRune(writer, r0, size0)
			if nil != err {
/////////////////////////////// RETURN
				return
			}
			n64 += int64(size0)

			if eof {
	/////////////////////// BREAK
				break
			} else {
	/////////////////////// CONTINUE
				continue
			}
		}

		var r1 rune
		var size1 int
		{
			r1, size1, err = readRune(reader)
			eof = errors.Is(err, io.EOF)
			if nil != err && !eof {
/////////////////////////////// RETURN
				return
			}
		}

		if size1 <= 0 {
			err = writeRune(writer, r0, size0)
			if nil != err {
/////////////////////////////// RETURN
				return
			}
			n64 += int64(size0)
	/////////////// BREAK
			break
		}

		if '\n' != r1 {
			{
				err = writeRune(writer, r0, size0)
				if nil != err {
/////////////////////////////////////// RETURN
					return
				}
				n64 += int64(size0)
			}

			{
				err = writeRune(writer, r1, size1)
				if nil != err {
/////////////////////////////////////// RETURN
					return
				}
				n64 += int64(size1)
			}

			if eof {
/////////////////////////////// BREAK
				break
			} else {
/////////////////////////////// CONTINUE
				continue
			}
		}

	/////// BREAK
		break
	}

	return n64, nil
}
