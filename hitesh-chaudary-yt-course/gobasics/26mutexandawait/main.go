package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition, mutex and await")
	wg := &sync.WaitGroup{}
	var score = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		fmt.Println("One Routine")
		score = append(score, 1)
		wg.Done()
	}(wg)

	// wg.Add(1)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Two Routine")
		score = append(score, 2)
		wg.Done()
	}(wg)

	// wg.Add(1)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Third Routine")
		score = append(score, 3)
		wg.Done()
	}(wg)

	wg.Wait()
	fmt.Println(score)
}
