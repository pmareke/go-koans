package go_koans

func aboutMaps() {
	ages := map[string]int{
		"bob": 10,
		"joe": 20,
		"dan": 30,
	}

	age := ages["bob"]
	Assert(age == __int__) // map syntax is warmly familiar

	age, ok := ages["bob"]
	Assert(ok == __bool__) // with a handy multiple-assignment variation

	age, ok = ages["steven"]
	Assert(age == __int__)    // the zero value is used when absent
	Assert(ok == __boolean__) // though there are better ways to check for presence

	Assert(len(ages) == __int__) // length is based on keys

	ages["bob"] = 99
	Assert(ages["bob"] == __int__) // values can be changed for keys

	ages["steven"] = 77
	Assert(ages[__string__] == 77) // new ones can be added

	delete(ages, "steven")
	age, ok = ages["steven"]
	Assert(ok == __boolean__) // key/value pairs can be removed
}
