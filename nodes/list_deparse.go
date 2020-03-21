// Auto-generated - DO NOT EDIT

package pg_query

import (
	"strings"
)

func (node List) Deparse(ctx DeparseContext) string {
	var items []string
	for _, i := range node.Items {
		items = append(items, i.Deparse(ctx))
	}
	return strings.Join(items, ", ")
}
