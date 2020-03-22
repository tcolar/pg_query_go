package pg_query_nodes

import "github.com/lfittl/pg_query_go/deparse"

func (node WithClause) Deparse(ctx deparse.Context) string {
	o := deparse.Output{}
	o.Append("WITH")
	o.Append(node.Ctes.Deparse(nil))
	return o.String()
}
