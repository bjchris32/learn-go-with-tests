package iteration

// repeat the string for given times. If no iterations given or the given iterations is less than 1, repeat just 1 time
func Repeat(character string, iterations ...int) string {
	var repeated string
	num_iterations := 1
	if len(iterations) > 0 {
		if iterations[0] > 0 {
			num_iterations = iterations[0]
		}
	}

	for i := 0; i < num_iterations; i++ {
		repeated += character
	}
	return repeated
}
