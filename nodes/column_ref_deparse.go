package pg_query_nodes

import "github.com/lfittl/pg_query_go/deparse"

func (node ColumnRef) Deparse(ctx deparse.Context) string {
	return node.Fields.DeparseSep(nil, ".")
}
