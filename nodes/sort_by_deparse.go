package pg_query_nodes

import "github.com/lfittl/pg_query_go/deparse"

func (node SortBy) Deparse(ctx deparse.Context) string {
	o := deparse.Output{}
	o.Append(node.Node.Deparse(nil))
	switch node.SortbyDir {
	case SORTBY_ASC:
		o.Append("ASC")
	case SORTBY_DESC:
		o.Append("DESC")
	}
	switch node.SortbyNulls {
	case SORTBY_NULLS_FIRST:
		o.Append("NULLS FIRST")
	case SORTBY_NULLS_LAST:
		o.Append("NULLS LAST")
	}
	return o.String()
}
