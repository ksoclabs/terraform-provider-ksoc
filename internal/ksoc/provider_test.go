package ksoc

import (
	"testing"
)

func TestProvider(t *testing.T) {
	if err := New("test")().InternalValidate(); err != nil {
		t.Fatal(err)
	}
}
