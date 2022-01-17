package main

import (
	"fmt"
	"sync"
)

func Tricky3() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go printNum(i, &wg)
	}
	wg.Wait()
}

func printNum(i int, wg *sync.WaitGroup) {
	fmt.Print(i)
	wg.Done()
}

func main() {
	Tricky3()
}
