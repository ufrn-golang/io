// Example to demonstrate reading an input file with package [bufio].
// Method ReadString reads the file contents as strings until EOF is
// found.
//
// Author: Everton Cavalcante (everton.cavalcante@ufrn.br)
// Date: April 3, 2023

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, inputError := os.Open("test.in")
	if inputError != nil {
	 	fmt.Println("An error occurred on opening the input file")
	 	os.Exit(1)
	}

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, err := inputReader.ReadString('\n')
		fmt.Print(inputString)
		if err == io.EOF {
			break
		}
	}

	inputFile.Close()
}
