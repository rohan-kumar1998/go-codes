package main

import (
	"fmt"
)

type foo1 struct {
	name string
}

func (r foo1) foo2(i *int) {
	fmt.Println(r)
}

type doo struct {
	foo1
	name string
}

func (d doo) foo2(i *int) {
	fmt.Println(d)
}

type poo interface {
	foo2(*int)
}

func main() {
	// mp := map[rune]string{}
	// s := "Hello"
	// l := len(s)
	var x foo1 = foo1{"hello"}
	fmt.Println(x)
	var y doo = doo{name: "string"}
	fmt.Println(y)
	// x.foo2()
	var pop [2]poo = [2]poo{x, y}
	fmt.Println(pop)
}
