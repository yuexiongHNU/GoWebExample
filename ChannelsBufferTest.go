package main

import (
	. "fmt"
)

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	Println(<-c)
	Println(<-c)
	close(c)
}
