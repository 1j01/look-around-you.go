package main

import (
	"fmt"
	"time"
)

func main() {
	message := "LOOK AROUND YOU! "
	codeLines := [...]string{
		"10 PRINT \"" + message + "\"",
		"20 GOTO 10",
		"RUN",
	}
	for _, codeLine := range codeLines {
		fmt.Printf(codeLine + "\n")
		time.Sleep(1 * time.Second)
	}
	for {
		fmt.Printf(message)
	}
}
