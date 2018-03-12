package main

import . "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	phone int
}

func main()  {
	tom := Student{Human{"tom", 18, "13123452345"}, 1111111111111}
	Printf("%s %d", tom.Human.phone, tom.phone)
}
