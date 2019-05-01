package main

import "fmt"

func main() {
	var x uint64

	fmt.Print("Enter a large int number (>255): ")
	fmt.Scanln(&x)

	fmt.Printf("You inputed number: %d\n", x)
	fmt.Printf("If we convert the number to byte type we have: %d\n", byte(x))
}
