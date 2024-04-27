package go_koans

func isPrimeNumber(possiblePrime int) bool {
	for underPrime := 2; underPrime < possiblePrime; underPrime++ {
		if possiblePrime%underPrime == 0 {
			return false
		}
	}
	return true
}

func findPrimeNumbers(channel chan int) {
	for i := 2; ; /* infinite loop */ i++ {
		// your code goes here

		Assert(i < 100) // i is afraid of heights
	}
}

func aboutConcurrency() {
	ch := make(chan int)

	Assert(__delete_me__) // concurrency can be almost trivial
	// your code goes here

	Assert(<-ch == 2)
	Assert(<-ch == 3)
	Assert(<-ch == 5)
	Assert(<-ch == 7)
	Assert(<-ch == 11)
}
