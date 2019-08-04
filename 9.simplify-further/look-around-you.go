package main

import (
	"fmt"
	"time"
)

const {
	charDelay = time.Duration(0.1*float64(time.Second))
	beforeNewLineDelay = time.Duration(0.5 * float64(time.Second))
	betweenLinesDelay = time.Duration(0.8 * float64(time.Second))
	afterRUNDelay = time.Duration(0.1 * float64(time.Second))
}

func main() {
	message := "LOOK AROUND YOU! "
	codeLines := []string{
		"10 PRINT \"" + message + "\"",
		"20 GOTO 10",
		"RUN",
	}

	for i, line := range codeLines {
		for _, c := range line {
			fmt.Printf(string(c))
			time.Sleep(charDelay)
		}

		time.Sleep(beforeNewLineDelay)
		fmt.Printf("\n")
		if i < len(codeLines)-1 {
			// in between lines (human pause)
			time.Sleep(betweenLinesDelay)
		} else {
			// after "RUN"
			time.Sleep()
		}
	}

	for {
		fmt.Printf(message)
	}
}
