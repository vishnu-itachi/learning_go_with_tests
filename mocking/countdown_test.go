package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	sleeper := &spy{}
	Countdown(buffer, *sleeper)

	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
