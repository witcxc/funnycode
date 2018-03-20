package main

import (
	"fmt"
)

type Person struct {
	age  int
	name string
}

func (Person p) change() int {
	p.age += 1
	fmt.Printf("\nin func %v", p)
	return 0

}
func change(t map[string]int) {
	t["hello"] = 13
	fmt.Printf("\nin func %v", t)

}
func main() {
	var p = Person{12, "hello"}
	p.change()
	fmt.Printf("\nout %v", p)
	a := make(map[string]int)
	a["hello"] = 12
	change(a)
	fmt.Printf("\nout func %v", a)
}
