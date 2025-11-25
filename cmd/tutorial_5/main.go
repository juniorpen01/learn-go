package main

import "strings"

func main() {
	{
		foo := "bar"
		bar := "baz"

		strBuilder := strings.Builder{}
		strBuilder.WriteString(foo)
		strBuilder.WriteString(bar)
	}
	{
		foo := "foo"
		foo += foo
	}
}
