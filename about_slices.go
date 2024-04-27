package go_koans

func aboutSlices() {
	fruits := []string{"apple", "orange", "mango"}

	Assert(fruits[0] == "apple") // slices seem like arrays
	Assert(len(fruits) == 3)     // in nearly all respects

	tasty_fruits := fruits[1:3]         // we can even slice slices
	Assert(tasty_fruits[0] == "orange") // slices of slices also share the underlying data

	pregnancy_slots := []string{"baby", "baby", "lemon"}
	Assert(cap(pregnancy_slots) == 3) // the capacity is initially the length

	pregnancy_slots = append(pregnancy_slots, "baby!")
	Assert(len(pregnancy_slots) == 4) // slices can be extended with append(), much like realloc in C
	Assert(cap(pregnancy_slots) == 6) // but with better optimizations

	pregnancy_slots = append(pregnancy_slots, "another baby!?", "yet another, oh dear!", "they must be Catholic")

	Assert(len(pregnancy_slots) == 7)  // append() can take N arguments to append to the slice
	Assert(cap(pregnancy_slots) == 12) // the capacity optimizations have a guessable algorithm
}
