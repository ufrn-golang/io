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

func main() {
	file, err := os.OpenFile("test.out", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error when creating the file")
        os.Exit(1)
	}

	// Package loremipsum is a Lorem Ipsum generator for Go
	// https://github.com/XANi/loremipsum
	liGen := loremipsum.New()
	par := liGen.Paragraph()
	if _, err := file.Write([]byte(par + "\n")); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	file.Close()
}
