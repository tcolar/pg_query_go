package deparse

import (
	"fmt"
	"regexp"
	"strings"
)

func DeparseString(s *string, forceEscape bool) *string {
	if s == nil {
		return nil
	}
	match, _ := regexp.MatchString(`[/^\w+$/]`, *s)
	if forceEscape || !match || isKeyword(*s) {
		escaped := fmt.Sprintf(`"%s"`, strings.ReplaceAll(*s, `"`, `""`))
		return &escaped
	}
	return s
}

func CompactJoin(joiner string, items ...*string) string {
	results := []string{}
	for _, i := range items {
		if i != nil {
			results = append(results, *i)
		}
	}
	return strings.Join(results, joiner)
}
