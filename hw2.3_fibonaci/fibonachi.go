package main

import (
	"fmt"
	"math/big"
)

 // реализовал по примеру как из разбора домашнего задания

func main() {
	// Initialize two big ints with the first two numbers in the sequence.
	a := big.NewInt(0)
	b := big.NewInt(1)

	for i := 1; i < 100; i++ {
		a.Add(a, b)
		a, b = b, a
		fmt.Println(a)
	}


/*   изначально фибоначи реализовал наверное так же как и большинство
func main() {
	a := 1
	b := 2

	for i := 1; i < 100; i++ {

		a, b = b, a + b

		fmt.Println(a)
	} */

}