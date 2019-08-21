package stringutils

import (
	"fmt"

	"github.com/fatih/color"
)

var R, Y, G func(...interface{}) string

func init() {
	// Create SprintXxx functions to mix strings
	// with other non-colorized strings:
	G = color.New(color.BgHiGreen).SprintFunc()
	Y = color.New(color.BgHiYellow).SprintFunc()
	R = color.New(color.FgWhite).Add(color.BgRed).Add(color.Bold).SprintFunc()
	fmt.Printf("mu.colors.init: %s %s %s \n",
		G("Okay"), Y("Warn"), R("ERR!"))
}
