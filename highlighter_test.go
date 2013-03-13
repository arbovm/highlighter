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
	{"a", "b", "<b>b</b>"},
	{"aa", "ab", "a<b>b</b>"},
	{"Hallo", "Hello", "H<b>e</b>llo"},
	{"Halli", "Hello", "H<b>e</b>ll<b>o</b>"},
	{"Hi", "Hello", "H<b>ello</b>"},
	{"Hello", "Hi", "H<b>i</b>"},
	{"pipelight", "diff", "<b>d</b>i<b>ff</b>"},
}

func TestDiffLines(t *testing.T) {

	for index, diffTest := range diffTests {
		result := DiffLines(diffTest.first, diffTest.second, "<b>", "</b>")
		if result != diffTest.wanted {
			output := fmt.Sprintf("%v \t diff of %v and %v should be %v but was %v.",
				index, diffTest.first, diffTest.second, diffTest.wanted, result)
			t.Errorf(output)
			output = strings.Replace(output, "\033", "esc", -1)
		}
		fmt.Print("\n")
	}
}
