package pg_query_nodes

import "github.com/lfittl/pg_query_go/deparse"

func (node A_Const) Deparse(ctx deparse.Context) string {
	return node.Val.Deparse(deparse.NewContext(deparse.AconstContext))
}
