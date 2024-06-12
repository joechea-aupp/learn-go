package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generator(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for range [5]int{} {
		time.Sleep(time.Millisecond * 200)
		ch <- rand.Int()
	}

	close(ch)
	fmt.Println("generator done")
}

func consumer(id int, sleep time.Duration, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range ch {
		time.Sleep(time.Millisecond * sleep)
		fmt.Printf("consumer %d: %d\n", id, task)
	}
	fmt.Printf("consumer %d done\n", id)
}

func main() {
	// rand.Seed(42)
	ch := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(3)

	go generator(ch, &wg)
	go consumer(1, 100, ch, &wg)
	go consumer(2, 200, ch, &wg)

	wg.Wait()
}
