package main

import "fmt"

type foo struct {
	bar, baz int
	int
	qux
}

func (q *foo) waldo() int {
	return 4
}

type qux struct {
	quux string
}

func (q qux) waldo() int {
	return 3
}

type garply interface {
	waldo() int
}

func fred(g garply) {
	fmt.Println(g.waldo())
}

func main() {
	{
		corge := foo{qux: qux{"grault"}}

		fmt.Println(corge)
		fmt.Println(corge.int)
		fmt.Println(corge.qux)
		fmt.Println(corge.qux.quux)
		fmt.Println(corge.quux)
	}

	{
		fred(&foo{}) // Apparently escapes to heap, haven't confirmed it yet tho
		fred(qux{})
	}
}
