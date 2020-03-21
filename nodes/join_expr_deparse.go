package pg_query_nodes

import (
	"fmt"

	"github.com/lfittl/pg_query_go/deparse"
)

func (node JoinExpr) Deparse(ctx deparse.Context) string {
	o := deparse.Output{}
	o.Append(node.Larg.Deparse(nil))
	switch node.Jointype {
	case JOIN_INNER:
		if node.IsNatural {
			o.Append("NATURAL")
		} else if node.Quals == nil && node.UsingClause.Empty() {
			o.Append("CROSS")
		}
	case JOIN_LEFT:
		o.Append("LEFT")
	case JOIN_FULL:
		o.Append("FULL")
	case JOIN_RIGHT:
		o.Append("RIGHT")
	}
	o.Append("JOIN")
	o.Append(node.Rarg.Deparse(nil))
	if node.Quals != nil {
		o.Append("ON")
		o.Append(node.Quals.Deparse(nil))
	}
	if !node.UsingClause.Empty() {
		o.Append(fmt.Sprintf(`USING (%s)`, node.UsingClause.Deparse(nil)))
	}
	return o.String()
}
