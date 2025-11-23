package main

import (
	"fmt"
	"time"
)

func main() {
	{
		foo := [...]int{1, 2, 3}
		bar := foo[:]
		fmt.Println(foo)
		fmt.Println(foo[1:3])
		bar = append(bar, bar...)
		fmt.Println(bar)
	}

	{
		foo := make(map[string]int)
		bar := map[string]int{"": 1, "john": 0, "fortnite": -1}
		fmt.Println(foo)
		fmt.Println(bar)
	}
	{
		testTimeWithVsWithoutPreallocation()
	}
}

// Idk if correct
func testTimeWithVsWithoutPreallocation() {
	const n = 1_000_000

	slice1 := []int{}
	slice2 := make([]int, 0, n)

	var testTimeToAppendNElements = func(s []int, n int) time.Duration {
		now := time.Now()

		for len(s) < n {
			s = append(s, 2)
		}

		return time.Since(now)
	}

	fmt.Printf("Total time without preallocation: %s\n", testTimeToAppendNElements(slice1, n))
	fmt.Printf("Total time with preallocation: %s\n", testTimeToAppendNElements(slice2, n))
}
