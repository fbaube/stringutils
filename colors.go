package stringutils

import (
	"fmt"

	"github.com/fatih/color"
)

var Rbg, Ybg, Gbg, Rfg, Yfg, Gfg func(...interface{}) string

func init() {
	// Create SprintXxx functions to mix strings
	// with other non-colorized strings:
	Gbg = color.New(color.BgHiGreen).SprintFunc()
	Ybg = color.New(color.BgHiYellow).SprintFunc()
	Rbg = color.New(color.BgHiRed).SprintFunc()
	Gfg = color.New(color.FgHiGreen).Add(color.Bold).SprintFunc()
	Yfg = color.New(color.FgHiYellow).Add(color.Bold).SprintFunc()
	Rfg = color.New(color.FgHiRed).Add(color.Bold).SprintFunc()

	Rfg = color.New(color.FgRed).Add(color.Bold).SprintFunc()
	fmt.Printf("su.colors.init: \n (bg) %s %s %s \n (fg) %s %s %s \n",
		Gbg(" Okay "), Ybg(" Warn "), Rbg(" Err! "),
		Gfg(" Okay "), Yfg(" Warn "), Rfg(" Err! "))
}
