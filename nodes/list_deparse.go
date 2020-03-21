package pg_query_nodes

import (
	"strings"

	"github.com/lfittl/pg_query_go/deparse"
)

func (node List) Deparse(ctx deparse.Context) string {
	var items []string
	for _, i := range node.Items {
		items = append(items, i.Deparse(ctx))
	}
	return strings.Join(items, ", ")
}
