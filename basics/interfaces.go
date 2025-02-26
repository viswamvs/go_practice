package main

import "fmt"

/*
	Interfaces are named collections of method signatures.
*/

type InterfaceExample interface {
	print()
}

type Struct1 struct {
}

type Struct2 struct {
}

func (s1 *Struct1) print() {
	fmt.Println("inside struct 1")
}

func (s2 *Struct2) print() {
	fmt.Println("inside struct 2")
}

func printMethod(i InterfaceExample) {
	i.print()
}

/*
	func main() {
		s1 := &Struct1{}
		s2 := &Struct2{}

		printMethod(s1)
		printMethod(s2)
	}
*/
