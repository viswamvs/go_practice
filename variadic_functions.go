package main

/*
	Variadic functions can be called with any number of trailing arguments.
	For example, fmt.Println is a common variadic function.
*/
func sumOfN(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	return sum
}

/*
	func main() {
		nums := []int{1,2,3,4,5,6}

		fmt.Println(sumOfN(nums...))

		fmt.Println(sumOfN(1,2,3)) // Variadic functions can be called in the usual way with individual arguments.
	}
*/
