package customvalidator

import (
	"testing"
)

func TestFormatError1(t *testing.T) {
	res := FormatError(nil)
	if res != nil {
		t.Fatal("Must be nil")
	}
}
