package pg_query_nodes

import (
	"fmt"

	"github.com/lfittl/pg_query_go/deparse"
)

func (node CommonTableExpr) Deparse(ctx deparse.Context) string {
	o := deparse.Output{}
	o.Append(*node.Ctename)
	if !node.Aliascolnames.Empty() {
		o.Append(fmt.Sprintf(`(%s)`, node.Aliascolnames.Deparse(nil)))
	}
	o.Append(fmt.Sprintf(`AS (%s)`, node.Ctequery.Deparse((nil))))
	return o.String()
}
