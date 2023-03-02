package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}

	score := make([]int, 0)

	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		fmt.Println("One R")
		score = append(score, 1)
		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		fmt.Println("Two R")
		score = append(score, 2)
		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		go fmt.Println("Three R")
		score = append(score, 3)
		wg.Done()
	}(wg)

	wg.Wait()
	fmt.Println(score)
}
