package main

import (
	"bytes"
	"fmt"
)

type Sleeper interface {
	Sleep()
}

type spy struct {
	calls int
}

func (s *spy) Sleep() {
	fmt.Println("Slept once")
	s.calls++
}

func main() {
	// Countdown()
}

func Countdown(out *bytes.Buffer, sleeper spy) {
	sleeper.Sleep()
	sleeper.Sleep()
	fmt.Fprint(out, "3")
}
