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

func typeOutLines(lines []string, charDelay time.Duration, beforeLineDelay time.Duration, afterLineDelay time.Duration) {
	for _, line := range lines {
		typeOutLine(line, charDelay)
		time.Sleep(beforeLineDelay)
		fmt.Printf("\n")
		time.Sleep(afterLineDelay)
	}
}

func main() {
	message := "LOOK AROUND YOU! "
	codeLines := []string{
		"10 PRINT \"" + message + "\"",
		"20 GOTO 10",
		"RUN",
	}
	typeOutLines(
		codeLines,
		time.Duration(0.1*float64(time.Second)),
		time.Duration(0.5*float64(time.Second)),
		time.Duration(0.2*float64(time.Second)),
	)
	for {
		fmt.Printf(message)
	}
}
