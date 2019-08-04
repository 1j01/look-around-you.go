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
	// typeOutLines(codeLines, 0.05*time.Second, 0.5*time.Second) doesn't work because Duration is int based so it would lose precision
	typeOutLines(
		codeLines,
		time.Duration(0.05*float64(time.Second)),
		time.Duration(0.5*float64(time.Second)),
	)
	for {
		fmt.Printf(message)
	}
}
