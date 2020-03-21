// Auto-generated - DO NOT EDIT

package pg_query

import (
	"fmt"
	"strings"

	"github.com/lfittl/pg_query_go/util"
)

func (node SelectStmt) Deparse(ctx DeparseContext) string {
	var o *util.Output

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
			o.Append(fmt.Sprintf("ON %s", node.DistinctClause.Deparse(NewSimpleContext(SelectContext))))
		}

		o.Append(fmt.Sprintf(node.TargetList.Deparse(NewSimpleContext(SelectContext))))

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
	/*

	   if node['groupClause']
	     output << 'GROUP BY'
	     output << node['groupClause'].map do |item|
	       deparse_item(item)
	     end.join(', ')
	   end

	   if node['havingClause']
	     output << 'HAVING'
	     output << deparse_item(node['havingClause'])
	   end

	   if node['sortClause']
	     output << 'ORDER BY'
	     output << node['sortClause'].map do |item|
	       deparse_item(item)
	     end.join(', ')
	   end

	   if node['limitCount']
	     output << 'LIMIT'
	     output << deparse_item(node['limitCount'])
	   end

	   if node['limitOffset']
	     output << 'OFFSET'
	     output << deparse_item(node['limitOffset'])
	   end

	   if node['lockingClause']
	     node['lockingClause'].map do |item|
	       output << deparse_item(item)
	     end
	   end

	   output.join(' ')
	*/
	return o.String()
}
