// Example to demonstrate reading a large input files with package [io/iotuil].
//
// File contents are read in chunks into a buffer (byte slice) with a predefined
// capacity. Reading is done by a [bufio.Reader] object and the Read function.
// Bytes read are further manipulated to another format (in this case, string).
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
	file, err := os.Open("test.dat")
	if err != nil {
		fmt.Println("Cannot read the file", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(file)
	size := 4 * 1024
	buffer := make([]byte, size)
	
	for {
		numbytes, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println((string(buffer[:numbytes])))
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println((string(buffer[:numbytes])))
	}

	file.Close()
}
