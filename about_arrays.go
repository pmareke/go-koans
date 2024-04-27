package go_koans

import "fmt"

func aboutArrays() {
	fruits := [4]string{"apple", "orange", "mango"}

	Assert(fruits[0] == __string__) // indexes begin at 0
	Assert(fruits[1] == __string__) // one is indeed the loneliest number
	Assert(fruits[2] == __string__) // it takes two to ...tango?
	Assert(fruits[3] == __string__) // there is no spoon, only an empty value

	Assert(len(fruits) == __int__) // the length is what the length is
	Assert(cap(fruits) == __int__) // it can hold no more

	Assert(fruits == [4]string{}) // comparing arrays is not like comparing apples and oranges

	tasty_fruits := fruits[1:3]                           // defining oneself as a variation of another
	Assert(fmt.Sprintf("%T", tasty_fruits) == __string__) //and get not a simple array as a result
	Assert(tasty_fruits[0] == __string__)                 // slices of arrays share some data
	Assert(tasty_fruits[1] == __string__)                 // albeit slightly askewed

	Assert(len(tasty_fruits) == __int__) // its length is manifest
	Assert(cap(tasty_fruits) == __int__) // but its capacity is surprising!

	tasty_fruits[0] = "lemon" // are their shared roots truly identical?

	Assert(fruits[0] == __string__) // has this element remained the same?
	Assert(fruits[1] == __string__) // how about the second?
	Assert(fruits[2] == __string__) // surely one of these must have changed
	Assert(fruits[3] == __string__) // but who can know these things

	veggies := [...]string{"carrot", "pea"}

	Assert(len(veggies) == __int__) // array literals need not repeat an obvious length
}
