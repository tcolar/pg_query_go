package pg_query_nodes

import (
	"strings"

	"github.com/lfittl/pg_query_go/deparse"
)

func (node List) Deparse(ctx deparse.Context) string {
	return node.DeparseSep(ctx, ", ")
}

func (node List) DeparseSep(ctx deparse.Context, sep string) string {
	return strings.Join(node.DeparseRaw(ctx), sep)
}

func (node List) DeparseRaw(ctx deparse.Context) []string {
	var items []string
	for _, i := range node.Items {
		items = append(items, i.Deparse(ctx))
	}
	return items
}
