package stringutils

import (
	"fmt"
	"github.com/fatih/color"
)

// Functions that are created via the library.
var Rbg, Ybg, Gbg, Rfg, Yfg, Gfg func(...interface{}) string
var Wfg, Blubg, Cyanbg func(...interface{}) string

func init() {
	// Create SprintXxx functions to mix strings
	// with other non-colorized strings:

	Gbg = color.New(color.BgHiGreen).SprintFunc()
	Ybg = color.New(color.BgHiYellow).SprintFunc()
	Rbg = color.New(color.BgHiRed).SprintFunc()
	Blubg = color.New(color.BgBlue).SprintFunc()
	Cyanbg = color.New(color.BgCyan).SprintFunc()

	Gfg = color.New(color.FgHiGreen).Add(color.Bold).SprintFunc()
	Yfg = color.New(color.FgHiYellow).Add(color.Bold).SprintFunc()
	Rfg = color.New(color.FgHiRed).Add(color.Bold).SprintFunc()
	Wfg = color.New(color.FgHiWhite).Add(color.Bold).SprintFunc()
	// Rfg = color.New(color.FgRed).Add(color.Bold).SprintFunc()
}

// ColorDemo can be run from anywhere anytime to see samples.
func ColorDemo() {
	fmt.Printf("stringutils.ColorDemo: \n      %s         %s         %s \n"+
		" %s \n      %s         %s         %s \n %s \n",
		Gbg(" Gbg:Okay "), Ybg(" Ybg:Warn "), Rbg(" Rbg:Err! "),
		"Gbg(\" Gbg:Okay \"), Ybg(\" Ybg:Warn \"), Rbg(\" Rbg:Err! \")",
		Gfg(" Gfg:Okay "), Yfg(" Yfg:Warn "), Rfg(" Rfg:Err! "),
		"Gfg(\" Gfg:Okay \"), Yfg(\" Yfg:Warn \"), Rfg(\" Rfg:Err! \")")
}
