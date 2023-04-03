package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(s)
}

func main() {

	// without a waitgroup, process is finished and newly created routines does not work.
	// there would be no output in the console.
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"phi",
		"zeta",
		"theta",
		"epsilon",
	}

	for i, w := range words {
		wg.Add(1)
		go printSomething(fmt.Sprintf("%d: %s", i, w), &wg)
	}

	wg.Wait()
}
