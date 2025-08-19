package core

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const (
	VERSION = "4.1.0"
)

func putAsciiArt(s string) {
	for _, c := range s {
		d := string(c)
		switch {
		case d == " ":
			color.Set(color.FgBlack) // Default to black for spaces
			d = " "
		case strings.Contains("☠", d): // Border elements in yellow
			color.Set(color.FgMagenta)
		case d == ":" || d == ":": // ":" or ":" in red
			color.Set(color.FgRed)
		case d == "!" || d == "!": // "!" or "!" in blue
			color.Set(color.FgRed)
		case d == "@": // Specific character "@" in red
			color.Set(color.FgRed)	
		case d == "\n": // Reset color on newline
			color.Unset()
		case strings.Contains("EXODUS", d): // PRO VERSION text in yellow
			color.Set(color.FgMagenta)
        case strings.Contains("v4.1.0", d): // VERSION number in green
			color.Set(color.FgGreen)
		case strings.Contains("Buy on https://www.clirpa.com", d):
			color.Set(color.FgGreen)
		default:
			color.Set(color.FgHiBlack) // Default all other characters to black
		}
		fmt.Print(d)
	}
	color.Unset()
}


func printLogo(s string) {
	for _, c := range s {
		d := string(c)
		switch d {
		case "_":
			color.Set(color.FgYellow) // Change color to yellow
		case "\n":
			color.Unset()
		default:
			color.Set(color.FgHiBlack)
		}
		fmt.Print(d)
	}
	color.Unset()
}


func Banner() {
	fmt.Println()

	// Display the new ASCII art banner with yellow border and red content
	putAsciiArt(`
☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠
☠     @@@@@@@@  @@@  @@@   @@@@@@   @@@@@@@   @@@  @@@   @@@@@@         ☠
☠     @@@@@@@@  @@@  @@@  @@@@@@@@  @@@@@@@@  @@@  @@@  @@@@@@@         ☠
☠     @@!       @@!  !@@  @@!  @@@  @@!  @@@  @@!  @@@  !@@             ☠
☠     !@!       !@!  @!!  !@!  @!@  !@!  @!@  !@!  @!@  !@!             ☠
☠     @!!!:!     !@@!@!   @!@  !@!  @!@  !@!  @!@  !@!  !!@@!!          ☠
☠     !!!!!:      @!!!    !@!  !!!  !@!  !!!  !@!  !!!   !!@!!!         ☠
☠     !!:        !: :!!   !!:  !!!  !!:  !!!  !!:  !!!       !:!        ☠
☠     :!:       :!:  !:!  :!:  !:!  :!:  !:!  :!:  !:!      !:!         ☠
☠     !:: ::::  !::  :::  ::::: ::  !:::: ::  :::::!::  ::::!::         ☠
☠     : :: ::   !:   ::    : :  :   :: :  :    : :  :   :: : :          ☠
☠     Buy on https://www.clirpa.com                   EXODUS v4.1.0     ☠
☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠
`)

}
