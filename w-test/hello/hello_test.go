package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("kamil")
	want := "hello kamil"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
