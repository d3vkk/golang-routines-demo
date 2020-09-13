package main

import (
  "fmt"
  "time"
)

func main() {
	message := "Hello Go routine"
	// Using closures to create an
	// anonymous function
	go func() {
		fmt.Println(message)
	}()
	time.Sleep(time.Millisecond * 10)
}
