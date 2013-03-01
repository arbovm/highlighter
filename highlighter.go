package highlighter

import "strings"
import "bytes"

func DiffLines(first string, second string) string {
	var result bytes.Buffer

	isInHighlight := false
	firstReader   := strings.NewReader(first)

	for _, b := range second {
		a, _, err := firstReader.ReadRune()
		if err == nil {
			if b != a {
				isInHighlight = true
				result.WriteString("\033[01;37m")
			} else if isInHighlight {
				isInHighlight = false
				result.WriteString("\033[0m")
			}
		}
		result.WriteString(string(b))
	}

	if isInHighlight {
		result.WriteString("\033[0m")
	}

	return result.String()
}
