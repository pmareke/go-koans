package go_koans

import (
	"io/ioutil"
	"strings"
)

func aboutFiles() {
	filename := "about_files.go"
	contents, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(contents), "\n")
	Assert(lines[5] == __string__) // handling files is too trivial
}
