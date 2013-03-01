package highlighter

import "testing"
import "fmt"
import "strings"

var diffTests = []struct {
	first  string
	second string
	wanted string
}{
	{"a", "a", "a"},
	{"a", "b", "\033[01;37mb\033[0m"},
	{"aa", "ab", "a\033[01;37mb\033[0m"},
	{"Hallo", "Hello", "H\033[01;37me\033[0mllo"},
	{"Halli", "Hello", "H\033[01;37me\033[0mll\033[01;37mo\033[0m"},
	{"Hi", "Hello", "H\033[01;37mello\033[0m"},
	{"Hello", "Hi", "H\033[01;37mi\033[0m"},
}

func TestDiffLines(t *testing.T) {

	for index, diffTest := range diffTests {
		result := DiffLines(diffTest.first, diffTest.second)
		if result != diffTest.wanted {
			output := fmt.Sprintf("%v \t diff of %v and %v should be %v but was %v.",
				index, diffTest.first, diffTest.second, diffTest.wanted, result)
			t.Errorf(output)
			output = strings.Replace(output, "\033", "esc", -1)
		}
		fmt.Print("\n")
	}
}
