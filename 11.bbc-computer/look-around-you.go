package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func clearScreen() {
	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "darwin": // macOS
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		// panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

const (
	initialDelay          = time.Duration(0.8 * float64(time.Second))
	charDelay             = time.Duration(0.1 * float64(time.Second))
	beforeNewLineDelay    = time.Duration(0.5 * float64(time.Second))
	betweenLinesDelay     = time.Duration(0.8 * float64(time.Second))
	beforeScreenSpamDelay = time.Duration(0.1 * float64(time.Second))
)

func main() {
	systemText := `
BBC Computer

Acorn DFS

BASIC

>`
	spamMessage := "LOOK AROUND YOU "
	codeLines := []string{
		"10 PRINT \"" + spamMessage + "\"",
		"20 GOTO 10",
		"RUN",
	}

	clearScreen()
	fmt.Printf(systemText)

	time.Sleep(initialDelay)

	for i, line := range codeLines {
		for _, c := range line {
			fmt.Printf(string(c))
			time.Sleep(charDelay)
		}

		time.Sleep(beforeNewLineDelay)

		if i < len(codeLines)-1 {
			fmt.Printf("\n>")
			time.Sleep(betweenLinesDelay)
		}
	}
	fmt.Printf("\n")
	time.Sleep(beforeScreenSpamDelay)
	for {
		fmt.Printf(spamMessage)
	}
}
