package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxChickenPrice = 5
const maxPossibleChickenPrice = 20

const maxTofuPrice = 3
const maxPossibleTofuPrice = 9

type website string

func (w website) checkChickenPrices(ch chan website) {
	for {
		time.Sleep(time.Second)

		chickenPrice := rand.Intn(maxPossibleChickenPrice + 1)
		if chickenPrice <= maxChickenPrice {
			ch <- w
			break
		}
	}
}

func (w website) checkTofuPrices(ch chan website) {
	for {
		time.Sleep(time.Second)

		tofuPrice := rand.Intn(maxPossibleTofuPrice + 1)
		if tofuPrice <= maxTofuPrice {
			ch <- w
			break
		}
	}
}

var websites = [...]website{"walmart.com", "costco.com", "wholefoods.com"}

func main() {
	chickenCh := make(chan website)
	tofuCh := make(chan website)

	for _, website := range websites {
		go website.checkChickenPrices(chickenCh)
		go website.checkTofuPrices(tofuCh)
	}

	select {
	case website := <-chickenCh:
		fmt.Printf("Found a deal on chicken at %s", website)

	case website := <-tofuCh:
		fmt.Printf("Found a deal on tofu at %s", website)
	}
}

// import (
// 	"log"
// 	"time"
// )

// func main() {
// 	c := make(chan int, 5)
// 	go func() {
// 		defer close(c)
// 		defer log.Println("Exiting process...")

// 		for i := range 5 {
// 			c <- i
// 		}

// 	}()
// 	for i := range c {
// 		log.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }
