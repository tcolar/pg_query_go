package pg_query_nodes

import (
	"fmt"

	"github.com/lfittl/pg_query_go/deparse"
)

func (node BoolExpr) Deparse(ctx deparse.Context) string {
	return fmt.Sprintf("NOT %s", node.Args.Items[0].Deparse(nil))
}
