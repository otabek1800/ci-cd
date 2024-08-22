package main

import (
	"testing"
)
func TestSeyHello(t *testing.T){
	masg := seyHello("go")
	want := "Hello go"

	if masg != want {
		t.Errorf("got %q, want %q", masg, want)
	}

}

func TestSeyGoodbye(t *testing.T) {
	msg := seyGoodbye("go")
	want := "Goodbye go"

	if msg != want {
		t.Errorf("got %q, want %q", msg, want)
	}
}

func TestSeyThankYou(t *testing.T) {
	msg := seyThankYou("go")
	want := "Thank you go"

	if msg != want {
		t.Errorf("got %q, want %q", msg, want)
	}
}

func TestSeyWelcome(t *testing.T) {
	msg := seyWelcome("go")
	want := "Welcome go"

	if msg != want {
		t.Errorf("got %q, want %q", msg, want)
	}
}

func TestSeyGoodMorning(t *testing.T) {
	msg := seyGoodMorning("go")
	want := "Good morning go"

	if msg != want {
		t.Errorf("got %q, want %q", msg, want)
	}
}

func TestSeyGoodNight(t *testing.T) {
	msg := seyGoodNight("go")
	want := "Good night go"

	if msg != want {
		t.Errorf("got %q, want %q", msg, want)
	}
}

func TestSeyGoodLuck(t *testing.T) {
	msg := seyGoodLuck("go")
	want := "Good luck go"

	if msg != want {
		t.Errorf("got %q, want %q", msg, want)
	}
}

func TestSeyCongratulations(t *testing.T) {
	msg := seyCongratulations("go")
	want := "Congratulations go"

	if msg != want {
		t.Errorf("got %q, want %q", msg, want)
	}
}

func TestSeySorry(t *testing.T) {
	msg := seySorry("go")
	want := "Sorry go"

	if msg != want {
		t.Errorf("got %q, want %q", msg, want)
	}
}

func TestSeySeeYou(t *testing.T) {
	msg := seySeeYou("go")
	want := "See you go"

	if msg != want {
		t.Errorf("got %q, want %q", msg, want)
	}
}