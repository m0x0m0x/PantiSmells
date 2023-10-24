
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
	fmt.Println("This is the help message.")
	// Add more information about how to use your application
}