package pg_query_nodes

import (
	"fmt"

	"github.com/lfittl/pg_query_go/deparse"
)

func (node A_Indices) Deparse(ctx deparse.Context) string {
	return fmt.Sprintf(`[%s]`, node.Uidx.Deparse(nil))
}
