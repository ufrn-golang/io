// Example to demonstrate reading user input from the console
// with the [fmt.Sscanf] function. Input considers a template
// against the input source to extract values.
//
// Author: Everton Cavalcante (everton.cavalcante@ufrn.br)
// Date: April 3, 2023

package main

import "fmt"

func main() {
	var name string
	var price float64

	source := "Product: Milk 5.95"
	template := "%s %f"
	
	fmt.Sscanf(source, template, &name, &price)
	fmt.Printf("Name: %s, Price: %.2f\n", name, price)
}
