// Just a reminder if I ever want to come back to this, I skipped the JSON and struct crap

package main

import (
	"log"
)

func main() {
	{
		foo := [...]int{1, 2, 3}
		log.Println(sumSlice(foo[:]))
	}

	{
		foo := []int{}
		bar := []int{0}

		log.Println(isEmpty(foo))
		log.Println(isEmpty(bar))
	}
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}
