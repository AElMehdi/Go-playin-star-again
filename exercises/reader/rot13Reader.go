package reader

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	in io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	char, err := r.in.Read(b)

	for index, char := range b {
		switch {
		case char >= 'A' && char <= 'M' || char >= 'a' && char <= 'm':
			b[index] += 13
		case char >= 'N' && char <= 'Z' || char >= 'n' && char <= 'z':
			b[index] -= 13
		}
	}
	return char, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
