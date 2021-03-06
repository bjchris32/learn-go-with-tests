package iteration

// check if the input has suffix in the string
func HasSuffix(input, suffix string) bool {
	suffixLength := len(suffix)
	inputLength := len(input)
	if inputLength < suffixLength {
		return false
	}
	inputSuffix := input[inputLength-suffixLength:inputLength]
	if inputSuffix == suffix {
		return true
	} else {
		return false
	}
}
