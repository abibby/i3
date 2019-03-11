package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/zwzn/i3/i3gobar/modules"
)

type block func() chan string

func main() {
	run(
		schedule(modules.DateTime, time.Second),
		schedule(modules.Pamac, time.Hour),
	)
}

func run(cs ...<-chan string) {
	for s := range merge(cs...) {
		fmt.Println(s)
	}
}

func schedule(cb func() string, every time.Duration) chan string {
	cs := make(chan string)
	go func() {
		cs <- cb()
		for range time.Tick(every) {
			cs <- cb()
		}
	}()
	return cs
}

func merge(cs ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan string) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
