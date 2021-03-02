package main

import (
	"fmt"
	"os"
)

func fiboMap(n int) int {
	// m := make(map[int]int)
	m := map[int]int{
		0: 0,
		1: 1,
		2: 1,
	}
	var j int
	if n == 2 {
		return 1
	}
	if n > 93 {
		fmt.Println("неверный индекс")
		os.Exit(1)
	}
	if n < 0 {
		fmt.Println("неверный индекс")
		os.Exit(1)
	}
	for i := 2; i < n; i++ {
		m[i] = m[i-1] + m[i-2]
		j = i
	}
	return m[j]
}

func main() {
	fmt.Println("введите индекс числа")
	var n int
	fmt.Scanln(&n)
	fmt.Println(fiboMap(n))
}
