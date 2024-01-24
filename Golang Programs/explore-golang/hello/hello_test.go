package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, World Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("checking if hello on empty string", func(t *testing.T){
		got := Hello("")
		want := "Hello, World "

		if got  != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
