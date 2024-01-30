package main

import "testing"

func TestHello(t *testing.T){
	got:= Hello("Chris")
	want:= "hello , Chris"

	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
}