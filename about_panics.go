package go_koans

func divideFourBy(i int) int {
	return 4 / i
}

const __divisor__ = 0

func aboutPanics() {
	Assert(__delete_me__) // panics are exceptional errors at runtime

	n := divideFourBy(__divisor__)
	Assert(n == 2) // panics are exceptional errors at runtime
}
