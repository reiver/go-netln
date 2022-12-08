package netln_test

import (
	"github.com/reiver/go-netln"

	"io"
	"strings"

	"testing"
)

func TestCopyLine(t *testing.T) {

	tests := []struct{
		Src string
		Expected string
	}{
		{
			Src:      "apple banana cherry",
			Expected: "apple banana cherry",
		},
		{
			Src:      "apple banana cherry"+"\r",
			Expected: "apple banana cherry"+"\r",
		},
		{
			Src:      "apple banana cherry"+"\n",
			Expected: "apple banana cherry"+"\n",
		},
		{
			Src:      "apple banana cherry"+"\r\n",
			Expected: "apple banana cherry",
		},
		{
			Src:      "apple banana cherry"+"\r"+"hello world",
			Expected: "apple banana cherry"+"\r"+"hello world",
		},
		{
			Src:      "apple banana cherry"+"\n"+"hello world",
			Expected: "apple banana cherry"+"\n"+"hello world",
		},
		{
			Src:      "apple banana cherry"+"\r\n"+"hello world",
			Expected: "apple banana cherry",
		},
		{
			Src:      "apple banana cherry"+"\r"+"hello world"+"\r",
			Expected: "apple banana cherry"+"\r"+"hello world"+"\r",
		},
		{
			Src:      "apple banana cherry"+"\r"+"hello world"+"\n",
			Expected: "apple banana cherry"+"\r"+"hello world"+"\n",
		},
		{
			Src:      "apple banana cherry"+"\r"+"hello world"+"\r\n",
			Expected: "apple banana cherry"+"\r"+"hello world",
		},
		{
			Src:      "apple banana cherry"+"\n"+"hello world"+"\r",
			Expected: "apple banana cherry"+"\n"+"hello world"+"\r",
		},
		{
			Src:      "apple banana cherry"+"\n"+"hello world"+"\n",
			Expected: "apple banana cherry"+"\n"+"hello world"+"\n",
		},
		{
			Src:      "apple banana cherry"+"\n"+"hello world"+"\r\n",
			Expected: "apple banana cherry"+"\n"+"hello world",
		},
		{
			Src:      "apple banana cherry"+"\r\n"+"hello world"+"\r\n",
			Expected: "apple banana cherry",
		},



		{
			Src:      "r۵≡🙂",
			Expected: "r۵≡🙂",
		},
		{
			Src:      "r۵≡🙂"+"\r",
			Expected: "r۵≡🙂"+"\r",
		},
		{
			Src:      "r۵≡🙂"+"\n",
			Expected: "r۵≡🙂"+"\n",
		},
		{
			Src:      "r۵≡🙂"+"\r\n",
			Expected: "r۵≡🙂",
		},
		{
			Src:      "r۵≡🙂"+"\r"+"once twice thrice fource",
			Expected: "r۵≡🙂"+"\r"+"once twice thrice fource",
		},
		{
			Src:      "r۵≡🙂"+"\r"+"once twice thrice fource"+"\r",
			Expected: "r۵≡🙂"+"\r"+"once twice thrice fource"+"\r",
		},
		{
			Src:      "r۵≡🙂"+"\r"+"once twice thrice fource"+"\n",
			Expected: "r۵≡🙂"+"\r"+"once twice thrice fource"+"\n",
		},
		{
			Src:      "r۵≡🙂"+"\r"+"once twice thrice fource"+"\r\n",
			Expected: "r۵≡🙂"+"\r"+"once twice thrice fource",
		},
		{
			Src:      "r۵≡🙂"+"\n"+"once twice thrice fource",
			Expected: "r۵≡🙂"+"\n"+"once twice thrice fource",
		},
		{
			Src:      "r۵≡🙂"+"\n"+"once twice thrice fource"+"\r",
			Expected: "r۵≡🙂"+"\n"+"once twice thrice fource"+"\r",
		},
		{
			Src:      "r۵≡🙂"+"\n"+"once twice thrice fource"+"\n",
			Expected: "r۵≡🙂"+"\n"+"once twice thrice fource"+"\n",
		},
		{
			Src:      "r۵≡🙂"+"\n"+"once twice thrice fource"+"\r\n",
			Expected: "r۵≡🙂"+"\n"+"once twice thrice fource",
		},
		{
			Src:      "r۵≡🙂"+"\r\n"+"once twice thrice fource",
			Expected: "r۵≡🙂",
		},
		{
			Src:      "r۵≡🙂"+"\r\n"+"once twice thrice fource"+"\r",
			Expected: "r۵≡🙂",
		},
		{
			Src:      "r۵≡🙂"+"\r\n"+"once twice thrice fource"+"\n",
			Expected: "r۵≡🙂",
		},
		{
			Src:      "r۵≡🙂"+"\r\n"+"once twice thrice fource"+"\r\n",
			Expected: "r۵≡🙂",
		},



		{
			Src:      "once"+"\r\n"+"twice"+"\r\n"+"thrice"+"\r\n"+"fource"+"\r\n",
			Expected: "once",
		},
	}

	for testNumber, test := range tests {

		var actualStorage strings.Builder

		var reader io.Reader = strings.NewReader(test.Src)

		actualN, err := netln.CopyLine(&actualStorage, reader)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %q", err, err)
			t.Logf("SRC: %q", test.Src)
			t.Logf("ACTUAL-N: %d", actualN)
			continue
		}

		{
			var expected int64 = int64(len(test.Expected))
			var actual   int64 = actualN

			if expected != actual {
				t.Errorf("For test #%d, the actual number of bytes written is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d bytes", expected)
				t.Logf("ACTUAL:   %d bytes", actual)
				t.Logf("SRC:          %q", test.Src)
				t.Logf("EXPECTED-DST: %q", test.Expected)
				t.Logf("ACTUAL-DST:   %q", actualStorage.String())
				continue
			}
		}

		{
			var expected string = test.Expected
			var actual   string = actualStorage.String()

			if expected != actual {
				t.Errorf("For test #%d, the actual value of what was written is not what was expected", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("SRC:      %q", test.Src)
				continue
			}
		}
	}
}
