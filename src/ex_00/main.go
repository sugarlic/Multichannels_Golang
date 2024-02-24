package main

import (
	"fmt"
	"sync"
	"time"
)

func SleepSort(nums []int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < len(nums); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(nums[i]) * time.Second)
			out <- nums[i]
		}(i)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch := SleepSort([]int{5, 2, 3})
	for elem := range ch {
		fmt.Println(elem)
	}
	fmt.Println("main() stopped")
}
