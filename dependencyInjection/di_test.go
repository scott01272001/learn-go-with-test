package dependencyinjection

import (
	"bytes"
	"os"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	t.Run("use stdout", func(t *testing.T) {
		Greet(os.Stdout, "Elodie")
	})
}
