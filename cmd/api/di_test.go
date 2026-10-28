package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T){
	buffer := bytes.Buffer{}
	Greet(&buffer, "Sami")

	got := buffer.String()
	want := "Hello, Sami"

	assertStringEquals(t, got, want)
}