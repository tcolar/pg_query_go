package pg_query_nodes

import "github.com/lfittl/pg_query_go/deparse"

func (node CaseExpr) Deparse(ctx deparse.Context) string {
	o := deparse.Output{}
	o.Append("CASE")
	if node.Arg != nil {
		o.Append(node.Arg.Deparse(nil))
	}
	if !node.Args.Empty() {
		o.Append(node.Args.Deparse(nil))
	}
	if node.Defresult != nil {
		o.Append("ELSE")
		o.Append(node.Defresult.Deparse(nil))
	}
	o.Append("END")
	return o.String()
}
