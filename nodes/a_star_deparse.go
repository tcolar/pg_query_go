package pg_query_nodes

import "github.com/lfittl/pg_query_go/deparse"

func (node A_Star) Deparse(ctx deparse.Context) string {
	return "*"
}
