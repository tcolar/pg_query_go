package pg_query_nodes

import "github.com/lfittl/pg_query_go/deparse"

func (node Null) Deparse(ctx deparse.Context) string {
	return "NULL"
}
