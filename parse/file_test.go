package parse

import (
	"testing"
)

func TestFile(t *testing.T) {
	filename := "../testdata/helloworld/helloworld.go"
	_, err := File(filename)
	if err != nil {
		t.Error(err)
	}
}
