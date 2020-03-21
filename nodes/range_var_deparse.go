package pg_query_nodes

import (
	"fmt"

	"github.com/lfittl/pg_query_go/deparse"
)

func (node RangeVar) Deparse(ctx deparse.Context) string {
	o := deparse.Output{}
	if !node.Inh {
		o.Append("ONLY")
	}
	schema := ""
	if node.Schemaname != nil {
		schema = fmt.Sprintf(`"%s".`, *node.Schemaname)
	}
	o.Append(fmt.Sprintf(`%s"%s"`, schema, *node.Relname))
	if node.Alias != nil {
		o.Append(node.Alias.Deparse(nil))
	}
	return o.String()
}
