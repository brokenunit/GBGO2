package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var workers = make(chan int, 1000)
	var j int
	var m sync.Mutex

	for i := 1; i <= 1000; i++ {
		// workers <- i

		go func() {
			defer func() {
				<-workers
			}()
			m.Lock()
			j++
			m.Unlock()
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(j)

}
