package go_koans

import "fmt"

func aboutStrings() {
	Assert("a"+"bc" == "abc") // string concatenation need not be difficult
	Assert(len("abc") == 3)   // and bounds are thoroughly checked

	Assert("abc"[0] == 'a') // their contents are reminiscent of C

	Assert("smith"[2:] == "ith")  // slicing may omit the end point
	Assert("smith"[:4] == "smit") // or the beginning
	Assert("smith"[2:4] == "it")  // or neither
	Assert("smith"[:] == "smith") // or both

	Assert("smith" == "smith") // they can be compared directly
	Assert("smith" < "smiths") // i suppose maybe this could be useful.. someday

	bytes := []byte{'a', 'b', 'c'}
	Assert(string(bytes) == "abc") // strings can be created from byte-slices

	bytes[0] = 'z'
	Assert(string(bytes) == "zbc") // byte-slices can be mutated, although strings cannot

	Assert(fmt.Sprintf("hello %s", "world") == "hello world")         // our old friend sprintf returns
	Assert(fmt.Sprintf("hello \"%s\"", "world") == "hello \"world\"") // quoting is familiar
	Assert(fmt.Sprintf("hello %q", "world") == "hello \"world\"")     // although it can be done more easily

	Assert(fmt.Sprintf("your balance: %d and %0.2f", 3, 4.5589) == "your balance: 3 and 4.56") // "the root of all evil" is actually a misquotation, by the way
}
