package main

import (
	"fmt"
	"log"
	"unicode/utf8"
)

func main() {
	// Overengineered crap
	var p = func(a ...any) {
		if _, err := fmt.Println(a); err != nil {
			log.Fatalf("unable to print line: %+v", err)
		}
	}

	{
		var foo uint16 = 65535
		bar := foo + 1

		p(foo)
		p(bar)
	}
	{
		foo := 12345678.9

		p(foo)
		p(float32(foo))
	}
	{
		foo := 3
		bar := 2

		p(foo / bar)
		p(foo % bar)
	}
	{
		foo := 'Æ'
		bar := "Æ°"

		p(string(foo))
		p(bar)

		p(len(bar))
		p(utf8.RuneCountInString(bar))
		p(len([]rune(bar)))
	}
	{
		foo := false

		p(foo)
	}
}
