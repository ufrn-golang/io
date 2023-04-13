// Example to demonstrate reading input from flags provided in the
// command-line with package [flags].
//
// Author: Everton Cavalcante (everton.cavalcante@ufrn.br)
// Date: April 3, 2023

package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "", "A name to say hello to.")
var french bool

func main() {
	flag.BoolVar(&french, "french", false, "Use French language.")
	flag.BoolVar(&french, "fr", false, "Use French language.")

	flag.Parse()
	if french {
		fmt.Printf("Salut %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}
