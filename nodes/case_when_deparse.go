package pg_query_nodes

import "github.com/lfittl/pg_query_go/deparse"

func (node CaseWhen) Deparse(ctx deparse.Context) string {
	o := deparse.Output{}
	o.Append("WHEN")
	o.Append(node.Expr.Deparse(nil))
	o.Append("THEN")
	o.Append(node.Result.Deparse(nil))
	return o.String()
}
