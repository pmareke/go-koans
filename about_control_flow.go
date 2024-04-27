package go_koans

import "fmt"

func aboutControlFlow() {
	{
		a, b, c := 1, 2, 3
		Assert(a == __int__) // multiple assignment
		Assert(b == __int__) // can make
		Assert(c == __int__) // life easier
	}

	var str string

	{
		if 3.14 == 3 {
			str = "what is love?"
		} else {
			str = "baby dont hurt me"
		}
		Assert(str == __string__) // no more

		if length := len(str); length == 17 {
			str = "to be"
		} else {
			str = "or not"
		}
		Assert(str == __string__) // that is the question
	}

	{
		hola1, hola2 := "ho", "la"

		switch "hello" {
		case "hello":
			str = "hi"
		case "world":
			str = "planet"
		case fmt.Sprintf("%s%s", hola1, hola2):
			str = "senor"
		}
		Assert(str == __string__) // cases can be of any type, even arbitrary expressions

		switch {
		case false:
			str = "first"
		case true:
			str = "second"
		}
		Assert(str == __string__) // in the absence of value, there is truth
	}

	{
		n := 0
		for i := 0; i < 5; i++ {
			n += i
		}
		Assert(n == __int__) // for can have the structure with which we are all familiar
	}

	{
		n := 1
		for {
			n *= 2
			if n > 20 {
				break
			}
		}
		Assert(n == __int__) // though omitting everything creates an infinite loop
	}
}
