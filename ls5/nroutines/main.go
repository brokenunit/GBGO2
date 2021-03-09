package main

import (
	"fmt"
	"sync"
	"time"
)

func oneMoreRoutine(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("routine %d start\n", i)
	time.Sleep(time.Second)
	fmt.Printf("routine %d end\n", i)
}
func main() {
	var n int
	fmt.Println("vvedite N")
	fmt.Scanln(&n)
	var wg sync.WaitGroup
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go oneMoreRoutine(i, &wg)
	}
	wg.Wait()
}
