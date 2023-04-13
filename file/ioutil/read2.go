// Example to demonstrate reading files with package [io/iotuil].
// Function ReadFile function reads the contents of the file and
// returns it as a slice of bytes.
//
// Author: Everton Cavalcante (everton.cavalcante@ufrn.br)
// Date: April 3, 2023

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	content, err := ioutil.ReadFile("test.in")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	fmt.Println("File contents: ")
	fmt.Println(string(content))
}
