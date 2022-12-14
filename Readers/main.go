package main

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	if len(b) == 0 {
		err := fmt.Errorf("math: square root of negative number %s", b)
		return 0, err
	}

	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}

	return len(b), nil
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}
