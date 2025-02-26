package main

/*
	Go supports pointers, allowing you to pass references to values and records within your program.

	passByValue -> will get a copy of variable's value
	passByReference -> will get a memory address of variable, it can acces to change the variable's value
*/

func passByValue(n int) {
	n += 10
}

func passByReference(n *int) {
	*n += 10
}

/*
	func main() {
		n := 10

		fmt.Println("initial value:", n) // 10

		passByValue(n)
		fmt.Println("after pass by value:", n) // 10

		passByReference(&n)
		fmt.Println("after pass by reference:", n) // 20
	}
*/
