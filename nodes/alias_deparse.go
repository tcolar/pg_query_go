package pg_query_nodes

import (
	"fmt"

	"github.com/lfittl/pg_query_go/deparse"
)

func (node Alias) Deparse(ctx deparse.Context) string {
	n := deparse.DeparseString(node.Aliasname, false)
	if node.Colnames.Empty() {
		return *n
	}
	return fmt.Sprintf("%s (%s)", *n, node.Colnames.Deparse(nil))
}
