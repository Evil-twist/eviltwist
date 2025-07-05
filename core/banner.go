package core

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const (
	VERSION = "3.8.0"
)

func putAsciiArt(s string) {
	for _, c := range s {
		d := string(c)
		switch string(c) {
		case " ":
			color.Set(color.BgBlack)
			d = " "
		case "@":
			color.Set(color.BgRed)
			d = " "
		case "#":
			color.Set(color.BgRed)
			d = " "
		case "W":
			color.Set(color.BgWhite)
			d = " "
		case "_":
			color.Unset()
			d = " "
		case "║"
            color.set(color.BgWhite)
            d= " "
        case "═"
            color.set(color.BgWhite)		
		case "\n":
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
	nameClr := color.New(color.FgHiWhite)
	txt := nameClr.Sprintf("               - --  EvilTwist Edition  -- -")
	fmt.Fprintf(color.Output, "%s", txt)
}

func printOneliner1() {
	handleClr := color.New(color.FgHiBlue)
	versionClr := color.New(color.FgGreen)
	textClr := color.New(color.FgHiBlack)
	spc := strings.Repeat(" ", 10-len(VERSION))
	txt := textClr.Sprintf("      modified by Jesali Blurry (") + handleClr.Sprintf("@Eviltwist") + textClr.Sprintf(")") + spc + textClr.Sprintf("version ") + versionClr.Sprintf("%s", VERSION)
	fmt.Fprintf(color.Output, "%s", txt)
}

func printOneliner2() {
	textClr := color.New(color.FgHiBlack)
	red := color.New(color.FgRed)
	white := color.New(color.FgWhite)
	txt := textClr.Sprintf("                   NO ") + red.Sprintf("PEACE") + white.Sprintf(" - ") + textClr.Sprintf("DO ") + red.Sprintf("EVIL")
	fmt.Fprintf(color.Output, "%s", txt)
}

func Banner() {
	fmt.Println()
	putAsciiArt("	                                           "#"                                                 ")	
	fmt.Println()
	putAsciiArt("                                            "#####"                                               ")
	fmt.Println()
	putAsciiArt("                                       "##############"                                           ")	
	fmt.Println()
	putAsciiArt("                                     "###################"                                        ")
	fmt.Println()
	putAsciiArt("                                   "#######################"                                      ")	
	fmt.Println()
	putAsciiArt("                                  "##########################"                                    ")	
	fmt.Println()
	putAsciiArt("                                 "#############################"                                  ")	
	fmt.Println()
	putAsciiArt("                                "###############################"                                 ")	
	fmt.Println()
	putAsciiArt("                               "###########"         "###########"                                ")	
	fmt.Println()
	putAsciiArt("                              "#########"               "#########"                               ")	
	fmt.Println()
	putAsciiArt("                             "########"                    "########"                             ")	
	fmt.Println()
	putAsciiArt("                            "########"    ##         ##     "########"                            ")
	printLogo(`                   @@@@@@@╗@@╗   @@╗@@╗@@╗  @@@@@@@@╗@@╗    @@╗@@╗@@@@@@@╗@@@@@@@@╗                 `)
	fmt.Println()
	putAsciiArt("                           "########"    ####       ####     "$######"                            ")
	printLogo(`                   @@╔════╝@@║   @@║@@║@@║  ╚══@@╔══╝@@║    @@║@@║@@╔════╝╚══@@╔══╝                 `)	
	fmt.Println()
	putAsciiArt("                           "#######"      ##         ##       "#######"                           ")
        printLogo(`                   @@@@@╗  @@║   @@║@@║@@║     @@║   @@║ @╗ @@║@@║@@@@@@@╗   @@║                    `)	
	fmt.Println()
	putAsciiArt("                           "#######"                           "######"                           ")
	printLogo(`                   @@╔══╝  ╚@@╗ @@╔╝@@║@@║     @@║   @@║@@@╗@@║@@║╚════@@║   @@║                    `)
	fmt.Println()
	putAsciiArt("                            "######"                            "#####"                           ")
	printLogo(`                   @@@@@@@╗ ╚@@@@╔╝ @@║@@@@@@@╗@@║   ╚@@@╔@@@╔╝@@║@@@@@@@║   @@║                    `)
	fmt.Println()
	putAsciiArt("                            "######"        ww        ww         "######"                          ")
	printLogo(`                   ╚══════╝  ╚═══╝  ╚═╝╚══════╝╚═╝    ╚══╝╚══╝ ╚═╝╚══════╝   ╚═╝                    `)
	fmt.Println()
	putAsciiArt("                            "######"          wwwwwwwwww         "#####"                          ")	
	fmt.Println()
	putAsciiArt("                             "######"           wwwww           "#####"                           ")	
	fmt.Println()
	putAsciiArt("                              "#####"                           "####"                            ")	
	fmt.Println()
	putAsciiArt("                            "########"                         "####"                             ")	
	fmt.Println()
	putAsciiArt("                          "############"                      "#######"                           ")	
	fmt.Println()
	putAsciiArt("                       "#################"                  "############"                        ")	
	fmt.Println()
	putAsciiArt("                      "######################"@EVILTWIST"##################"                      ")	
	fmt.Println()
	putAsciiArt("                     "#######################################################"                    ")	
	fmt.Println()
	putAsciiArt("                    "#########################################################"                   ")	
	//printOneliner2()
	//fmt.Println()
	putAsciiArt("                   "############################################################"                 ")
	printOneliner1()
	fmt.Println()
	putAsciiArt("                  "##############################################################"                ")
	fmt.Println()
}
