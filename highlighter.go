package highlighter

import "strings"
import "bytes"

func DiffLines(first, second, start, stop string) string {
	var result bytes.Buffer

	isInHighlight := false
	firstReader   := strings.NewReader(first)

	for _, b := range second {
		a, _, err := firstReader.ReadRune()
		if err == nil {
			if b != a {
				if !isInHighlight {
					isInHighlight = true
					result.WriteString(start)
				}
			} else if isInHighlight {
				isInHighlight = false
				result.WriteString(stop)
			}
		}
		result.WriteString(string(b))
	}

	if isInHighlight {
		result.WriteString(stop)
	}

	return result.String()
}
