package pg_query_nodes

import (
	"fmt"
	"strings"

	"github.com/lfittl/pg_query_go/deparse"
)

func (node SelectStmt) Deparse(ctx deparse.Context) string {
	o := deparse.Output{}

	if node.WithClause != nil {
		o.Append(node.WithClause.Deparse(nil))
	}
	switch node.Op {
	case 1:
		if node.Larg != nil && !node.Larg.SortClause.Empty() {
			o.Append(fmt.Sprintf("(%s)", node.Larg.SortClause.Deparse(nil)))
		}
		o.Append("UNION")
		if node.All {
			o.Append("ALL")
		}
		if node.Rarg != nil && !node.Rarg.SortClause.Empty() {
			o.Append(fmt.Sprintf("(%s)", node.Rarg.SortClause.Deparse(nil)))
		}
	case 3:
		o.Append(node.Larg.Deparse(nil))
		o.Append("EXCEPT")
		o.Append(node.Rarg.Deparse(nil))
	}

	if !node.FromClause.Empty() || !node.TargetList.Empty() {
		o.Append("SELECT")
	}

	if !node.TargetList.Empty() {
		if !node.DistinctClause.Empty() {
			o.Append("DISTINCT")
			o.Append(fmt.Sprintf("ON %s", node.DistinctClause.Deparse(deparse.NewSimpleContext(deparse.SelectContext))))
		}

		o.Append(fmt.Sprintf(node.TargetList.Deparse(deparse.NewSimpleContext(deparse.SelectContext))))

		if node.IntoClause != nil {
			o.Append("INTO")
			o.Append(node.IntoClause.Deparse(nil))
		}
	}

	if !node.FromClause.Empty() {
		o.Append("FROM")
		o.Append(node.FromClause.Deparse(nil))
	}

	if node.WhereClause != nil {
		o.Append("WHERE")
		o.Append(node.WhereClause.Deparse(nil))
	}

	if node.ValuesLists != nil && len(node.ValuesLists) > 0 {
		o.Append("VALUES")
		var lists []string
		for _, list := range node.ValuesLists {
			var items []string
			for _, i := range list {
				items = append(items, i.Deparse(ctx))
			}
			lists = append(lists, fmt.Sprintf("(%s)", strings.Join(items, ", ")))
		}
		o.Append(strings.Join(lists, ", "))
	}

	if !node.GroupClause.Empty() {
		o.Append("GROUP BY")
		o.Append(node.GroupClause.Deparse(nil))
	}

	if node.HavingClause != nil {
		o.Append("HAVING")
		o.Append(node.HavingClause.Deparse(nil))
	}

	if !node.SortClause.Empty() {
		o.Append("ORDER BY")
		o.Append(node.SortClause.Deparse(nil))
	}

	if node.LimitCount != nil {
		o.Append("LIMIT")
		o.Append(node.LimitCount.Deparse(nil))
	}

	if node.LimitOffset != nil {
		o.Append("OFFSET")
		o.Append(node.LimitOffset.Deparse(nil))
	}

	if !node.LockingClause.Empty() {
		o.Append(node.LockingClause.Deparse(nil))
	}

	return o.String()
}
