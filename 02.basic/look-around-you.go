package main

import "fmt"

func main() {
	var message string = "LOOK AROUND YOU! ";
	fmt.Printf("10 PRINT \"" + message + "\"\n20 GOTO 10\nRUN\n")
	for {
		fmt.Printf(message)
	}
}
