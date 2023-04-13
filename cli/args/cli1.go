// Example to demonstrate reading input from command-line arguments
// provided by [os.Args].
//
// Author: Everton Cavalcante (everton.cavalcante@ufrn.br)
// Date: April 3, 2023

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		args := os.Args[1:]
		fmt.Println("Arguments provided:", args)
	} else {
		fmt.Println("No arguments have been provided")
	}
}
