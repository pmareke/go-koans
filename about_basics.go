package go_koans

func aboutBasics() {
	Assert(true == true)  // what is truth?
	Assert(true != false) // in it there is nothing false

	var i int = 1
	Assert(i == 1.0000000000000000000000000000000000000) // precision is in the eye of the beholder

	k := 1 //short assignment can be used, as well
	Assert(k == 1.0000000000000000000000000000000000000)

	Assert(5%2 == 1)
	Assert(5*2 == 10)
	Assert(5^2 == 7)

	var x int
	Assert(x == 0) // zero values are valued in Go

	var f float32
	Assert(f == 0.0) // for types of all types

	var s string
	Assert(s == "") // both typical or atypical types

	var c struct {
		x int
		f float32
		s string
	}
	Assert(c.x == 0)   // and types within composite types
	Assert(c.f == 0.0) // which match the other types
	Assert(c.s == "")  // in a typical way
}
