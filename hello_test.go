package hello

import "testing"

func TestHello(t *testing.T) {
	msg := Hello()

	if msg != "Hello World" {
		t.Errorf(`Expected "Hello World", got "%s"`, msg)
	}
}
