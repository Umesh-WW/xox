package main

import "testing"

func TestMain(t *testing.T) {
    if got, want := "Hellos, Go!", "Hello, Go!"; got != want {
        t.Errorf("got %v, want %v", got, want)
    }
}
