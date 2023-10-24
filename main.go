
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	helpFlag := flag.Bool("help", false, "Display help")
	flag.Parse()

	if *helpFlag {
		displayHelp()
		os.Exit(0)
	}

	fmt.Println("SmellPussy")
	SmellPussy()
}

func displayHelp() {
	var help = ` 
What is big breasted ass fetish 
smell he ass and pyusy now 
	`
	fmt.Println(help)
	// Add more information about how to use your application
}