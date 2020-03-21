package pg_query_nodes

import "github.com/lfittl/pg_query_go/deparse"

func (node NullTest) Deparse(ctx deparse.Context) string {
	o := deparse.Output{}
	o.Append(node.Arg.Deparse(nil))
	switch node.Nulltesttype {
	case IS_NULL:
		o.Append("IS NULL")
	case IS_NOT_NULL:
		o.Append("IS NOT NULL")
	}
	return o.String()
}
