package main

import (
	"fmt"
	"time"
)

// Uppercase functions, types, and constants are public (exported)
// lowercase is private

func typeOutLine(line string, charDelay time.Duration) {
	for _, c := range line {
		fmt.Printf(string(c))
		time.Sleep(charDelay)
	}
}

func typeOutLines(lines []string, charDelay time.Duration, lineDelay time.Duration) {
	for _, line := range lines {
		typeOutLine(line, charDelay)
		time.Sleep(lineDelay)
		fmt.Printf("\n")
	}
}

func main() {
	message := "LOOK AROUND YOU! "
	codeLines := []string{
		"10 PRINT \"" + message + "\"",
		"20 GOTO 10",
		"RUN",
	}
	// typeOutLines(codeLines, 0.05*time.Second, 0.5*time.Second)
	// "[go] constant 0.05 truncated to integer"
	// it doesn't work because Duration is int based so it would lose precision
	// (an insignificant amount concidering it's stored as nanoseconds)
	// we can convince it we don't care by admitting we're converting between types
	typeOutLines(
		codeLines,
		time.Duration(0.05*float64(time.Second)),
		time.Duration(0.5*float64(time.Second)),
	)
	for {
		fmt.Printf(message)
	}
}
