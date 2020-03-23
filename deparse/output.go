package deparse

import (
	"strings"
)

// output string building helper
type Output struct {
	data []string
}

func (o *Output) Append(s string) {
	o.data = append(o.data, s)
}

func (o Output) String() string {
	return o.Join(" ")
}

func (o Output) Join(sep string) string {
	return strings.Join(o.data, sep)
}
