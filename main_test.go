package main

import (
	"testing"
)
func TestSeyHello(t *testing.T){
	masg := seyHello("go")
	want := "Hello "

	if masg != want {
		t.Errorf("got %q, want %q", masg, want)
	}

}