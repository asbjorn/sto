package sto

// FizzBuzz tests the input number and returns the
// correct response according to the 'fizzbuzz' rules
func FizzBuzz(i int) (int, string) {
	if i%3 == 0 && i%5 == 0 {
		return i, "fizzbuzz"
	} else if i%3 == 0 {
		return i, "fizz"
	} else if i%5 == 0 {
		return i, "buzz"
	}

	return i, ""
}
