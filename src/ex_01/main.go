package main

import (
	"fmt"
	"sync"
)

func Multiplex(channels ...chan interface{}) chan interface{} {
	out := make(chan interface{})

	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(ch chan interface{}) {
			defer wg.Done()
			for elem := range ch {
				out <- elem
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	input1 := make(chan interface{})
	input2 := make(chan interface{})
	input3 := make(chan interface{})

	output := Multiplex(input1, input2, input3)

	go func() {
		defer close(input1)
		defer close(input2)
		defer close(input3)
		input1 <- "Hello"
		input2 <- 42
		input3 <- true
		input2 <- "World"
		input3 <- 3.14
		input1 <- false
		input3 <- 10
		input1 <- "School21"
		input2 <- 1.618
	}()

	for elem := range output {
		fmt.Println(elem)
	}
}
