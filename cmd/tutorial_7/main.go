package main

import (
	"fmt"
	"log"
)

func main() {
	foo := [...]int{1, 2, 3}
	log.Printf("The memory location of the array is: %p\n", &foo)
	bar := squareThreeIntegers(&foo)

	fmt.Println(foo)
	fmt.Println(bar)
}

func squareThreeIntegers(f *[3]int) [3]int {
	log.Printf("The memory location of the array is: %p\n", f)
	for i, v := range f {
		f[i] = v * v
	}
	return *f
}
