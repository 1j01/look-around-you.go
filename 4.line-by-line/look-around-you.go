package main

import "fmt"

func main() {
	message := "LOOK AROUND YOU! "
	codeLines := [...]string{
		"10 PRINT \"" + message + "\"",
		"20 GOTO 10",
		"RUN",
	}
	for _, codeLine := range codeLines {
		fmt.Printf(codeLine)
	}
	for {
		fmt.Printf(message)
	}
}
