package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

type datum struct {
	id int
}

func (d datum) String() string {
	return fmt.Sprintf("Id: %d", d.id)
}

func (d datum) process() {
	delay := rand.Intn(6)
	time.Sleep(time.Duration(delay) * time.Second)
	log.Println(d)
	wg.Done()
}

var data = [...]datum{{1}, {2}, {3}, {4}, {5}}

func main() {
	now := time.Now()

	for _, datum := range data {
		wg.Add(1)
		go datum.process()
	}
	wg.Wait()

	log.Printf("Total execution time: %v\n", time.Since(now))
}
