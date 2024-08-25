package mocking

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func Countdown(buffer io.Writer) {
	fmt.Fprint(buffer, "3")
}

func main() {
	Countdown(os.Stdout)
}
