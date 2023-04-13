// Example to demonstrate reading user input from the console
// with the [bufio.Reader] object from the [bufio] package.
// Method ReadString is used to read input as a string until
// the new line character is found.
//
// Author: Everton Cavalcante (everton.cavalcante@ufrn.br)
// Date: April 3, 2023

package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Print("Please enter some input: ")	
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("The input was:", input)
}
