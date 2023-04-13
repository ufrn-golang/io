// Example to demonstrate reading user input from the console
// with the [bufio.Scanner] object from the [bufio] package.
// Method Scan is used to continuously read values until the
// program is told to stop reading by pressing Ctrl+D in the
// keyboard.
//
// Author: Everton Cavalcante (everton.cavalcante@ufrn.br)
// Date: April 3, 2023

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter some input: ")
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
		fmt.Print("Please enter some input: ")
	}
}
