// Example to demonstrate adding contents to a file with package [bufio].
// File is opened to enable data being appended to it, and new contents are
// written with a [bufio.Writer] object and the Write function.
//
// Author: Everton Cavalcante (everton.cavalcante@ufrn.br)
// Date: April 3, 2023

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/XANi/loremipsum"
)

const numparagraphs = 3

func main() {
	file, err := os.OpenFile("test.out", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error when creating the file")
        os.Exit(1)
	}
	
	// Package loremipsum is a Lorem Ipsum generator for Go
	// https://github.com/XANi/loremipsum
	var contents []string
	liGen := loremipsum.New()
	for i := 0; i < numparagraphs; i++ {
		par := liGen.Paragraph()
		contents = append(contents, par, "\n\n")
	}

	writer := bufio.NewWriter(file)
	for _, chunk := range contents {
		if _, err := writer.Write([]byte(chunk)); err != nil {
			writer.Flush()
			fmt.Println(err)
			os.Exit(1)
		}
		writer.Flush()
	}

	file.Close()
}
