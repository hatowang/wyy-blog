package main

import "fmt"

type person struct {
	name string
	age int
	num int
}

func main() {
	p := person{
		name: "yyy",
	}
	rebuildPerson(p)
	fmt.Println(p.name)
}

func rebuildPerson(p person) {
	p.name = "xxx"
}
