package di

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
)

func NotTestableGreet(name string) {
	fmt.Printf("Hello, %s", name)
}

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Kazim")

	got := buffer.String()
	want := "Hello, Kazim"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func Greet(writer io.Writer, name string) {
	//NotTestableGreet(s)
	fmt.Fprintf(writer, "Hello, %s", name)
}

func mainStdout() {
	Greet(os.Stdout, "Kamuran")
}

func GreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func mainHttp() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(GreeterHandler)))
}
