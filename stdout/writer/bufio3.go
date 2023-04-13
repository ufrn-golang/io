// Example to demonstrate writing to the standard output
// with the [bufio.Writer] object from the [bufio] package.
// The values to appear in the standard output are successively
// written into the [bufio.Writer] object with the WriteString method.
// The output appears in the console when calling the Flush method
// on the [bufio.Writer] object
//
// Author: Everton Cavalcante (everton.cavalcante@ufrn.br)
// Date: April 3, 2023

package main

import (
	"bufio"
	"os"
)

func main() {
   phrases := []string{ 
      "Break the ice", 
      "No ifs, ands, or buts", 
      "Any bells ringing", 
   }
   
   writer := bufio.NewWriter(os.Stdout)
   for _, phrase := range phrases {
      writer.WriteString("- " + phrase + "\n")
   }
   writer.Flush()
}
