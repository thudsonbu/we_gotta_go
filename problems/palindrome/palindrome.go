package palindrome

func IsPalindrome(x int) bool {
	if x < 0 || x == 10 {
		return false
	}

	// determine size
	reduced, magnitude := x, 1
	for true {
		magnitude *= 10
		reduced = x

		reduced -= magnitude
		if reduced <= 0 {
			break
		}
	}

	// single digit numbers are palindromes
	if magnitude == 10 {
		return true
	}

	// subtract from highest and lowest magnitudes until zero
	// or single num apart
	var (
		out           = false
		lowMagnitude  = 10
		highMagnitude = magnitude / 10
		remainder     = x
	)

	for true {
		low := (remainder % lowMagnitude) / (lowMagnitude / 10)
		high := remainder / highMagnitude

		diff := low - high

		remainder -= low + highMagnitude*high

		if diff != 0 {
			break
		}

		lowMagnitude *= 10

		if lowMagnitude >= highMagnitude {
			out = true
			break
		}

		highMagnitude /= 10
	}

	return out
}
