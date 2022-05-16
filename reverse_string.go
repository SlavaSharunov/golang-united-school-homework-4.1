package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	runes := []rune(input)
	for r := len(runes) - 1; r >= 0; r-- {
		output += string(runes[r])
	}
	return output
}
