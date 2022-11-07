package numpalindrom

func IsPalindrome(x int) bool {
	out := false

	// determine size
	reduced, magnitude := x, 1
	for true {
		magnitude *= 10
		reduced = x

		if x >= 0 {
			reduced -= magnitude
			if reduced <= 0 {
				break
			}
		} else {
			reduced += magnitude
			if reduced >= 0 {
				break
			}
		}
	}

	// single digit numbers are palindromes
	if magnitude == 10 {
		return true
	}

	// subtract from highest and lowest magnitudes until zero
	// or single num apart
	lowMagnitude := 10
	highMagnitude := magnitude / 10
	for true {
		lowNum := x % lowMagnitude
		highNum := x / highMagnitude

		diff := lowNum - highNum

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
