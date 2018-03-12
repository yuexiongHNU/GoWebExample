package main

import . "fmt"

type Skills []string

type Human struct {
	name string
	age int
	weight int
}

type Student struct {
	Skills // 匿名字段，自定义类型 string lice
	Human // 匿名字段，struct
	int // 内置类型作为匿名字段
	speciality string
}

func main()  {
	jane := Student{Human:Human{"jane", 18, 120}, speciality:"program"}
	Printf("%s %d %s \n", jane.Human.name, jane.Human.age, jane.speciality)
	tom := Student{Skills:Skills{"golang", "python", "c++"}, Human:Human{"jane", 30, 16}, int:10, speciality:"program"}
	Printf("%s %d \n", tom.Skills[1], tom.int)
	jane.Skills = append(jane.Skills, "ruby")
	Printf("%s", jane.Skills)
}
