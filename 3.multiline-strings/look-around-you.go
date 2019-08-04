package main

import "fmt"

func main() {
	var message string = "LOOK AROUND YOU! ";
	fmt.Printf(
`10 PRINT "` + message + `"
20 GOTO 10
RUN
`)
	for {
		fmt.Printf(message)
	}
}
