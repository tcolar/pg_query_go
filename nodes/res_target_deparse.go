package pg_query_nodes

import (
	"log"

	"github.com/lfittl/pg_query_go/deparse"
)

func (node ResTarget) Deparse(ctx deparse.Context) string {
	if ctx == nil && node.Name != nil {
		return *node.Name
	}
	var val string
	switch ctx.Kind() {
	case deparse.SelectContext:
		val = node.Val.Deparse(nil)
		return deparse.CompactJoin(
			" AS ",
			&val,
			deparse.DeparseString(node.Name, false),
		)
	case deparse.UpdateContext:
		val = node.Val.Deparse(nil)
		return deparse.CompactJoin(
			" = ",
			&val,
		)
	}
	log.Fatalf("Can't deparse %v in context %v", node, ctx)
	return ""
}
