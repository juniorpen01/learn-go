package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mu sync.RWMutex

type datum struct {
	id int
}

func (d datum) String() string {
	return fmt.Sprintf("Id: %d", d.id)
}

func (d datum) process(out *[]int) {
	delay := rand.Intn(6)
	time.Sleep(time.Duration(delay) * time.Second)

	log.Println(d)

	mu.Lock()
	*out = append(*out, d.id)
	mu.Unlock()

	// mu.RLock()
	// log.Println(*out)
	// mu.RUnlock()

	wg.Done()
}

var data = [1000]datum{{1}, {2}, {3}, {4}, {5}}

func main() {
	now := time.Now()

	out := make([]int, 0, len(data))

	for _, datum := range data {
		wg.Add(1)
		go datum.process(&out)
	}
	wg.Wait()

	log.Printf("Total execution time: %v\n", time.Since(now))
	log.Println(out)
}
