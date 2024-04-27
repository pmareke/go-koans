package go_koans

import "fmt"

func aboutArrays() {
	fruits := [4]string{"apple", "orange", "mango"}

	Assert(fruits[0] == "apple")  // indexes begin at 0
	Assert(fruits[1] == "orange") // one is indeed the loneliest number
	Assert(fruits[2] == "mango")  // it takes two to ...tango?
	Assert(fruits[3] == "")       // there is no spoon, only an empty value

	Assert(len(fruits) == 4) // the length is what the length is
	Assert(cap(fruits) == 4) // it can hold no more

	Assert(fruits == [4]string{"apple", "orange", "mango"}) // comparing arrays is not like comparing apples and oranges

	tasty_fruits := fruits[1:3]                           // defining oneself as a variation of another
	Assert(fmt.Sprintf("%T", tasty_fruits) == "[]string") //and get not a simple array as a result
	Assert(tasty_fruits[0] == "orange")                   // slices of arrays share some data
	Assert(tasty_fruits[1] == "mango")                    // albeit slightly askewed

	Assert(len(tasty_fruits) == 2) // its length is manifest
	Assert(cap(tasty_fruits) == 3) // but its capacity is surprising!

	tasty_fruits[0] = "lemon" // are their shared roots truly identical?

	Assert(fruits[0] == "apple") // has this element remained the same?
	Assert(fruits[1] == "lemon") // how about the second?
	Assert(fruits[2] == "mango") // surely one of these must have changed
	Assert(fruits[3] == "")      // but who can know these things

	veggies := [...]string{"carrot", "pea"}

	Assert(len(veggies) == 2) // array literals need not repeat an obvious length
}
