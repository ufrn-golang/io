// Example to demonstrate reading user input from the console
// with the [fmt.Scanln] function.
// A sequence of values (separated with blank spaces) is read
// from the standard input and respectively assigned to the
// defined variables.
//
// Author: Everton Cavalcante (everton.cavalcante@ufrn.br)
// Date: April 3, 2023

package main

import (
	"fmt"
	"os"
)

func main() {
   var name string
   var price float32

   fmt.Println("Enter text to scan: ")
   numvalues, err := fmt.Scanln(&name, &price)
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   fmt.Printf("Scanned %v values\n", numvalues)
   fmt.Printf("Name: %v, Price: %.2f\n", name, price)
}
