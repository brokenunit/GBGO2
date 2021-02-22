package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	a := 2
	b := 0 //для вызова паники поставь "0"
	abdivide(a, b)
	fmt.Println("все еще работает")
	err := fmt.Errorf("error occurred at: %v", time.Now())
	fmt.Println("An error happened:", err)
}

func abdivide(a int, b int) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
			f, err := os.Create("created.txt")
			if err != nil {
				panic(err)
			}
			if _, err = f.WriteString("Some text"); err != nil {
				panic(err)
			}
			f.Close()
		}
	}()
	fmt.Println(divide(a, b))
}

func divide(a, b int) int {
	return a / b
}
