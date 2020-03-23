package pg_query_nodes

import (
	"fmt"

	"github.com/lfittl/pg_query_go/deparse"
)

func (node TypeCast) Deparse(ctx deparse.Context) string {
	if node.TypeName.Deparse(nil) == "boolean" {
		if node.Arg.Deparse(nil) == "t" {
			return "true"
		}
		return "false"
	}
	var argCtx deparse.Context
	switch node.Arg.(type) {
	case A_Expr:
		argCtx = deparse.NewEmptyContext()
	}
	return fmt.Sprintf("%s::%s", node.Arg.Deparse(argCtx), node.TypeName.Deparse(nil))
}
