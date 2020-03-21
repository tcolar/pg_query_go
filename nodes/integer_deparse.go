package pg_query_nodes

import (
	"fmt"

	"github.com/lfittl/pg_query_go/deparse"
)

func (node Integer) Deparse(ctx deparse.Context) string {
	return fmt.Sprintf("%d", node.Ival)
}
