package go_koans

type coolNumber int

func (cn coolNumber) multiplyByTwo() int {
	return int(cn) * 2
}

func aboutTypes() {
	i := coolNumber(4)
	Assert(i == coolNumber(__int__))     // values can be converted between compatible types
	Assert(i.multiplyByTwo() == __int__) // you can add methods on any type you define
}
