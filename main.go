package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	var nums []int
	for i := 1; i < len(os.Args); i++ {
		n, err := strconv.Atoi(os.Args[i])
		if err != nil {
			panic("Only numbers can be sorted.")
		}
		nums = append(nums, n)
	}

	sleep_sort(nums)
}

func sleep_sort(nums []int) {
	wg := sync.WaitGroup{}
	for _, n := range nums {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			time.Sleep(time.Duration(n) * time.Second)
			fmt.Println(n)
		}(n)
	}
	wg.Wait()
}
