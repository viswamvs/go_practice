package main

/*
	A string in Go is a sequence of bytes that represents text.
	Strings are immutable (you cannot change a character in a string).
	Strings are UTF-8 encoded, meaning they support multiple languages.

	A rune is an integer representing a Unicode character.
	Since Go strings are byte-based, special characters (like é, ₹, 😀) take more than one byte.
	Rune helps handle characters correctly, even if they take multiple bytes.

	String						Immutable sequence of bytes (UTF-8 encoded)
	Rune						Represents a single Unicode character
	len(str)					Gives number of bytes, not characters
	Iterating with range		Reads characters correctly (including multi-byte Unicode)
	String to Rune Conversion	[]rune(str)

func main() {
	str := "Hello, Go!"

	fmt.Println("String:", str)
	fmt.Println("Length:", len(str))        // Length of string in bytes
	fmt.Println("First character:", str[0]) // Returns byte value

	// Convert byte to character
	fmt.Println("First character (as string):", string(str[0]))

	str1 := "GoLang 🚀"
    
    runes := []rune(str1) // Convert string to slice of runes
    fmt.Println("Rune Slice:", runes)
    fmt.Println("First rune:", string(runes[0])) // Convert back to string
}
*/