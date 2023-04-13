// Example to demonstrate writing output to a file with a file descriptor.
// Method WriteString write data (as string) to the file.
//
// Author: Everton Cavalcante (everton.cavalcante@ufrn.br)
// Date: April 3, 2023

package main

import (
	"fmt"
	"os"

	"github.com/XANi/loremipsum"
)

const numparagraphs = 2

func main() {
	file, err := os.Create("test.out")
	if err != nil {
		fmt.Println("Error when creating the file")
		os.Exit(1)
	}

	// Package loremipsum is a Lorem Ipsum generator for Go
	// https://github.com/XANi/loremipsum
	liGen := loremipsum.New()
	
	for i := 0; i < numparagraphs; i++ {
		par := liGen.Paragraph()
		file.WriteString(par + "\n\n")
	}

	file.Close()
}