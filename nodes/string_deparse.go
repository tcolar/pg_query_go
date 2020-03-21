// Auto-generated - DO NOT EDIT

package pg_query_nodes

import (
	"fmt"
	"strings"

	"github.com/lfittl/pg_query_go/deparse"
)

func (node String) Deparse(ctx deparse.Context) string {
	if ctx != nil {
		kind := ctx.Kind()
		switch kind {
		case deparse.AconstContext:
			return fmt.Sprintf("'%s'", strings.ReplaceAll(node.Str, "'", "''"))
		case deparse.ColumnOperator, deparse.DefNameAs, deparse.FuncCall, deparse.TypeName:
			return node.Str
		}
	}
	deparsed := deparse.DeparseString(&node.Str, true)
	return *deparsed
}
