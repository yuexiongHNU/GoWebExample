package main

import (
	. "fmt"
	"time"
)

func fibonacci2(c, quit chan int)  {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			Println("Quit!")
			return
		case <- time.After(5 * time.Second):
			Println("Time out!")
		default:
			Println("Default")
		}
	}
}

func main()  {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 15; i++ {
			Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}