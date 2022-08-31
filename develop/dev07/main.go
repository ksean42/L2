package main

import (
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, c := range channels {
		go func(c <-chan interface{}) {
			for v := range c {
				out <- v
			}
			wg.Done()
			fmt.Println("sig chan closed")
		}(c)
	}
	go func() {
		wg.Wait()
		fmt.Println("or chan closed")
		close(out)
	}()
	return out
}

func main() {
	sig := func(after time.Duration, id int) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
			c <- id
		}()
		return c
	}

	start := time.Now()
	ch := or(
		sig(5*time.Second, 1),
		sig(1*time.Second, 2),
		sig(1*time.Second/2, 3),
	)
	for {
		if res, ok := <-ch; ok {
			fmt.Println("output from", res, "sig channel")
		} else {
			break
		}
	}

	fmt.Printf("fone after %v", time.Since(start))
}
