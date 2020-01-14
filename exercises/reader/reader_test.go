package reader

import (
	"testing"
)

func TestThatMyReaderEmitsInfiniteA(t *testing.T) {
	myReader := MyReader{}
	text := []byte{1, 2, 65}

	i, e := myReader.Read(text)

	if e != nil {
		t.Error("Something went wrong in the implementation, error must be nil!")
	}

	if i > 1 {
		t.Error("It must return only one character")
	}
}
