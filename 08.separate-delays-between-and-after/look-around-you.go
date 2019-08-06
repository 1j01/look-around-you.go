package main

import (
	"fmt"
	"time"
)

func typeOutLine(line string, charDelay time.Duration) {
	for _, c := range line {
		fmt.Printf(string(c))
		time.Sleep(charDelay)
	}
}

// Sometimes you want to remove a function,
// if you realize you're only using it once
// and(/or) especially if it's more tied to the specific use case (or use cases) than you thought
//
// (I'd made typeOutLines into a function out of some vague notion of potential reusability,
// but actually I want a lot of behavior specific to what I'm doing (the one use case for it).
// Specifically, I want to delay in between the lines of code being typed out more than after the last line,
// because the last line makes it run, whereas the between ones are meant to be human pauses.
// I still want a bit of a delay after "RUN" so you can see it before the screen is flooded,
// so that's another delay, that I would probably do outside of the function typeOutLines,
// after making it only delay in between items - and so if I'm calling time.Sleep() outside as well,
// that's basically just spreading the logic around unnecessarily)
//
// In short: pulling stuff out into a function is very useful, but it's a two way street.
// You should question your abstractions, because there's a mental cost, versus something direct/unabstract.
//
// Related reading: Semantic Compression https://caseymuratori.com/blog_0015
// (a bit pretentious, but some useful tips)

func main() {
	message := "LOOK AROUND YOU! "
	codeLines := []string{
		"10 PRINT \"" + message + "\"",
		"20 GOTO 10",
		"RUN",
	}

	for i, line := range codeLines {
		typeOutLine(line, time.Duration(0.1*float64(time.Second)))
		time.Sleep(time.Duration(0.5 * float64(time.Second)))
		fmt.Printf("\n")
		if i < len(codeLines)-1 {
			// in between lines (human pause)
			time.Sleep(time.Duration(0.8 * float64(time.Second)))
		} else {
			// after "RUN"
			time.Sleep(time.Duration(0.1 * float64(time.Second)))
		}
	}

	for {
		fmt.Printf(message)
	}
}
