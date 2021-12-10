package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {
	_, err := rot13.r.Read(b)

	if err != nil {
		return 0, err
	}

	for index, value := range b {

		if value >= 65 && value <= 90 {
			// For Capital Letters
			temp := (((value - 65) + 13) % 26) + 65
			b[index] = temp
		} else if value >= 97 && value <= 122 {
			// For Non-Capital Letters
			temp := (((value - 97) + 13) % 26) + 97
			b[index] = temp
		}
		// Do nothing for non alphabets
	}
	return len(b), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
