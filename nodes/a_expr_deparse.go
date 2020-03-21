package pg_query_nodes

import (
	"fmt"

	"github.com/lfittl/pg_query_go/deparse"
)

func (node A_Expr) Deparse(ctx deparse.Context) string {
	o := deparse.Output{}
	o.Append(node.Lexpr.Deparse(deparse.PassOrEmptyContext(ctx)))
	o.Append(node.Rexpr.Deparse(deparse.PassOrEmptyContext(ctx)))
	sep := node.Name.Items[0].Deparse(deparse.NewContext(deparse.ColumnOperator))
	output := o.Join(fmt.Sprintf(" %s ", sep))
	if ctx != nil {
		// This is a nested expression, add parentheses.
		output = fmt.Sprintf("(%s)", output)
	}
	return output
}
