// Auto-generated - DO NOT EDIT

package pg_query_nodes

import (
	"fmt"

	"github.com/lfittl/pg_query_go/deparse"
)

func (node SubLink) Deparse(ctx deparse.Context) string {
	switch node.SubLinkType {
	case ANY_SUBLINK:
		return fmt.Sprintf(`%s IN (%s)`, node.Testexpr.Deparse(nil), node.Subselect.Deparse(nil))
	case ALL_SUBLINK:
		return fmt.Sprintf(`'%s %s ALL (%s)'`,
			node.Testexpr.Deparse(nil),
			node.OperName.Items[0].Deparse(deparse.NewContext(deparse.ColumnOperator)),
			node.Subselect.Deparse(nil))
	case EXISTS_SUBLINK:
		return fmt.Sprintf(`'EXISTS(%s)'`, node.Subselect.Deparse(nil))
	default:
		return fmt.Sprintf(`(%s)`, node.Subselect.Deparse(nil))
	}
}
