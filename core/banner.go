package core

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const (
	VERSION = "3.8.2"
)

func putAsciiArt(s string) {
	for _, c := range s {
		d := string(c)
		switch string(c) {
		case "#":
			color.Set(color.BgRed)
			d = " "
		case "@":
			color.Set(color.BgBlack)
			d = " "
		case ":":
			color.Set(color.BgGreen)
			d = " "
		case "$":
			color.Set(color.BgYellow)
			d = " "
		case "/":
			color.Set(color.BgBlue)
			d = " "
		case " ":
			color.Unset()
		}
		fmt.Print(d)
	}
	color.Unset()
}

func printLogo(s string) {
	for _, c := range s {
		d := string(c)
		switch string(c) {
		case "_":
			color.Set(color.FgWhite)
		case "\n":
			color.Unset()
		default:
			color.Set(color.FgHiBlack)
		}
		fmt.Print(d)
	}
	color.Unset()
}

func printUpdateName() {
	nameClr := color.New(color.FgHiCyan)
	txt := nameClr.Sprintf("            -- Evil Twist Edition --")
	fmt.Fprintf(color.Output, "%s", txt)
}

func printOneliner1() {
	handleClr := color.New(color.FgHiBlue)
	versionClr := color.New(color.FgGreen)
	textClr := color.New(color.FgHiBlack)
	spc := strings.Repeat(" ", 10-len(VERSION))
	txt := textClr.Sprintf("   Modified by EvilTwist @EVILTWIST ") + handleClr.Sprintf("https://t.me/eviltwist telegram") + spc + textClr.Sprintf("version ") + versionClr.Sprintf("%s", VERSION)
	fmt.Fprintf(color.Output, "%s", txt)
}

func printOneliner2() {
	textClr := color.New(color.FgHiBlack)
	red := color.New(color.FgRed)
	white := color.New(color.FgWhite)
	txt := textClr.Sprintf("           no ") + red.Sprintf("nginx") + white.Sprintf(" - ") + textClr.Sprintf("pure ") + red.Sprintf("evil")
	fmt.Fprintf(color.Output, "%s", txt)
}

func Banner() {
	fmt.Println()
	putAsciiArt("                                                                                                  ")
	fmt.Println()
	putAsciiArt("                                            $#####$                                               ")
	fmt.Println()
	putAsciiArt("                                         /##########/                                             ")
	fmt.Println()
	putAsciiArt("                                       $###############$                                          ")
	fmt.Println()
	putAsciiArt("                                     $###################$                                        ")
	fmt.Println()
	putAsciiArt("                                    /######################/                                      ")
	fmt.Println()
	putAsciiArt("                                   /#########################/                                    ")
	fmt.Println()
	putAsciiArt("                                  /###########################/                                   ")
	fmt.Println()
	putAsciiArt("                                 /#########/$:::::::$/#########/                                  ")
	fmt.Println()
	putAsciiArt("                                /#######/:::::::::::::::/#######$                                 ")
	fmt.Println()
	putAsciiArt("                               $######/:::::::::::::::::::/######$                                ")
	fmt.Println()
	putAsciiArt("                              /######/::::::::::::::::::::::######/                               ")
	fmt.Println()
	putAsciiArt("                              $######$$$$$$$$/$::::::$/$$$$$$$####$                               ")
	fmt.Println()
	putAsciiArt("                              /#####/$$$####/::::::::/####$$$/####/                               ")
	fmt.Println()
	putAsciiArt("                              /#####:::::/#: :::::::: :#/::::/####/                               ")
	fmt.Println()
	putAsciiArt("                              /####/::::::::::::::::::::::::://###/                               ")
	fmt.Println()
	putAsciiArt("                              /####/:::::::::::::::::::::::::////#$                               ")
	fmt.Println()
	putAsciiArt("                              $###//$:::::::::::::::::::::::$/####                                ")
	fmt.Println()
	putAsciiArt("                               $####/$::::::::::::::::::::::$/###$                                ")
	fmt.Println()
	putAsciiArt("                                /###//$:::::::::::::::::::$///##/                                 ")
	fmt.Println()
	putAsciiArt("                              /#######//::::::::::::::::://///#/                                  ")
	fmt.Println()
	putAsciiArt("                            /##########///::::::::::::://////####$                                ")
	fmt.Println()
	putAsciiArt("                         /###############/////::::::///    //########/                            ")
	fmt.Println()
	putAsciiArt("                        /#################  $/######/$    /############/                          ")
	fmt.Println()
	putAsciiArt("                       /##################  #########$   /##############/                         ")
	fmt.Println()
	putAsciiArt("                      /###################  #########   /################                         ")
	fmt.Println()
	putAsciiArt("                      ####################  ########   /#################/                        ")
	fmt.Println()
	putAsciiArt("                     /#####################  ########$ /##################                        ")
	fmt.Println()
	putAsciiArt("                     /######################/########$$  /################/                       ")
	fmt.Println()
	putAsciiArt("                     ######################$$##########    /###############                       ")
	fmt.Println()
	putAsciiArt("                    /####################################### /#############/                      ")
	fmt.Println()
	putAsciiArt("                    /########################################//#############                      ")
	fmt.Println()
	putAsciiArt("                    ########################################################/                     ")
	fmt.Println()
	putAsciiArt("                    ######//////////////////////////////////////////////////#                     ")
	fmt.Println()
	putAsciiArt("                    #####/                                                   /                    ")
	fmt.Println()
	putAsciiArt("                    /#####/                      .;'                         /#                    ")
	fmt.Println()
	putAsciiArt("                    #####$/               .'`  `-'  ``.                     /$#                    ")
	fmt.Println()
	putAsciiArt("                    #####/               ; EVILTTWIST ;                     /#                     ")
	fmt.Println()
	putAsciiArt("                    /#####/                `.__.-.__;                      /$                     ")
	fmt.Println()
	putAsciiArt("                     /####/                                              /$/                     ")
	fmt.Println()
	putAsciiArt("                       /###/                                                                      ")
	fmt.Println()
	putAsciiArt("                          /$/                                                                     ")
	fmt.Println()
	printLogo(`  ███████╗██╗   ██╗██╗██╗  ████████╗██╗    ██╗██╗███████╗████████╗         ██╗██╗███╗   ██╗██╗  ██╗`)
    fmt.Println()██╔════╝██║   ██║██║██║  ╚══██╔══╝██║    ██║██║██╔════╝╚══██╔══╝         ██║██║████╗  ██║╚██╗██╔╝`)
    printLogo(`  █████╗  ██║   ██║██║██║     ██║   ██║ █╗ ██║██║███████╗   ██║            ██║██║██╔██╗ ██║ ╚███╔╝ `)
    fmt.Println()██╔══╝  ╚██╗ ██╔╝██║██║     ██║   ██║███╗██║██║╚════██║   ██║       ██   ██║██║██║╚██╗██║ ██╔██╗ `)
    printLogo(`  ███████╗ ╚████╔╝ ██║███████╗██║   ╚███╔███╔╝██║███████║   ██║       ╚█████╔╝██║██║ ╚████║██╔╝ ██╗`)
    fmt.Println()╚══════╝  ╚═══╝  ╚═╝╚══════╝╚═╝    ╚══╝╚══╝ ╚═╝╚══════╝   ╚═╝        ╚════╝ ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝`)
    printLogo(`  Evilginx to telegram bot modified by Blurry(@Eviltwist)`)
    fmt.Println()
    printUpdateName()                                                                                                
    fmt.Println()
	printOneliner1()
	fmt.Println()
	printOneliner2()
	fmt.Println()
	fmt.Println()
}

func main() {
	Banner()
}
