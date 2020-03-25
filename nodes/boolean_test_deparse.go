package pg_query_nodes

import (
	"fmt"

	"github.com/lfittl/pg_query_go/deparse"
)

func (node BooleanTest) Deparse(ctx deparse.Context) string {
	boolTest := node.Booltesttype
	arg := node.Arg.Deparse(nil)
	switch boolTest {
	case IS_TRUE:
		return fmt.Sprintf("%s IS TRUE", arg)
	case IS_NOT_TRUE:
		return fmt.Sprintf("%s IS NOT TRUE", arg)
	default:
		panic(fmt.Sprintf("Unexpected bool test %v", boolTest))
	}
}
