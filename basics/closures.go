package main

/*
	Go supports anonymous functions, which can form closures. 
	Anonymous functions are useful when you want to define a function inline without having to name it.

	This function closureExample returns another function, which we define anonymously in the body of closureExample. 
	The returned function closes over the variable count to form a closure.

	Anonymous functions can also be recursive, 
	but this requires explicitly declaring a variable with var to store the function before it’s defined.

	var fibo func(n int) int

	fibo = func(n int) int {
		if n < 2 {
			return n
		}
		return fibo(n-1) + fibo(n-2)
	}
*/

func closureExample() func() int {
	count := 0

	return func() int{
		count++;
		return count
	}
}

/*
	func main() {

		counter1 := closureExample()
		fmt.Println(counter1())
		fmt.Println(counter1())

		counter2 := closureExample() // the state is unique to the particular function
		fmt.Println(counter2())
		fmt.Println(counter2())
	}
*/