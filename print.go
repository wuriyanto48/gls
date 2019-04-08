package gls

import (
	"fmt"
)

// PrintRedColor will print text to red text
func PrintRedColor(s string) {
	fmt.Printf("\033[31m%s\033[0m%s", s, "\n")
}

// PrintGreenColor will print text to green text
func PrintGreenColor(s string) {
	fmt.Printf("\033[32m%s\033[0m%s", s, "\n")
}

// PrintYellowColor will print text to yellow text
func PrintYellowColor(s string) {
	fmt.Printf("\033[33m%s\033[0m%s", s, "\n")
}

// PrintCyanColor will print text to cyan text
func PrintCyanColor(s string) {
	fmt.Printf("\033[36m%s\033[0m%s", s, "\n")
}
