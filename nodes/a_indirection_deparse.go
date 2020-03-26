package pg_query_nodes

import (
	"fmt"

	"github.com/lfittl/pg_query_go/deparse"
)

func (node A_Indirection) Deparse(ctx deparse.Context) string {

	panic("Not Implemented")

	o := deparse.Output{}
	switch node.Arg.(type) {
	case FuncCall, SubLink:
		o.Append(fmt.Sprintf("(%s)", node.Arg.Deparse((nil))))
	default:
		o.Append(node.Arg.Deparse(nil))
	}
	o.Append(node.Indirection.Deparse(nil))
	return o.String()
}
