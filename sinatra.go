package reverse

import (
	"fmt"
	"regexp"
	"strings"
)

func (us *URLStore) S(name string, s string) string {
	return us.MustAdd(name, s, extractSinatraParams(s)...)
}

const bc = "/.;,"

var sinatraPatternRe = regexp.MustCompile(fmt.Sprintf("[%s](:[^%s]+)", bc, bc))

func extractSinatraParams(s string) []string {
	matches := sinatraPatternRe.FindAllStringSubmatchIndex(s, -1)
	pats := make([]string, 0, len(matches)+1)

	for _, match := range matches {
		a, b := match[2], match[3]
		pats = append(pats, s[a:b])
	}

	if strings.HasSuffix(s, "/*") {
		pats = append(pats, "*")
	}

	return pats
}
