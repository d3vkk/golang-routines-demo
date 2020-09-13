package main

import (
  "fmt"
  "time"
)

func routine() {
	fmt.Println("Golang Goroutine")
}

func main() {
	go routine()
  /*
    main() function sleeps for 10ms
    in which time routine() is assigned

    Both main() and routine() are then
    executed concurrently
  */
	time.Sleep(time.Millisecond * 10)
}
