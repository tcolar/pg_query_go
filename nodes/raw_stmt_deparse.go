package pg_query_nodes

import "github.com/lfittl/pg_query_go/deparse"

func (node RawStmt) Deparse(ctx deparse.Context) string {
	return node.Stmt.Deparse(ctx)
}
