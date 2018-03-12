package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

// Return older person and age gap
func Older(p1, p2 person) (person, int){
	if p1.age > p2.age {
		return p1, p1.age-p2.age
	}
	return p2, p2.age-p1.age
}

func main() {
	var tom person
	tom.age, tom.name = 18, "tom"
	bob := person{age:27, name:"bob"}
	paul := person{"paul", 16}

	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, paul)
	bp_Older, bp_diff := Older(bob, paul)

	fmt.Printf("Of %s and %s, %s is oder by %d years\n", tom.name, bob.name, tb_Older.name, tb_diff)
	fmt.Printf("Of %s and %s, %s is oder by %d years\n", tom.name, paul.name, tp_Older.name, tp_diff)
	fmt.Printf("Of %s and %s, %s is oder by %d years\n", paul.name, bob.name, bp_Older.name, bp_diff)
}