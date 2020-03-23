package pg_query_nodes

import (
	"fmt"
	"strings"

	"github.com/lfittl/pg_query_go/deparse"
)

func (node FuncCall) Deparse(ctx deparse.Context) string {
	o := deparse.Output{}
	// SUM(a, b)
	args := node.Args.DeparseRaw(nil)
	// COUNT(*)
	if node.AggStar {
		args = append(args, "*")
	}
	var names []string
	for _, i := range node.Funcname.Items {
		name := i.Deparse(deparse.NewContext(deparse.FuncCall))
		if name != "pg_catalog" {
			names = append(names, i.Deparse(ctx))
		}
	}
	name := strings.Join(names, ".")
	distinct := ""
	if node.AggDistinct {
		distinct = "DISTINCT"
	}
	o.Append(fmt.Sprintf("%s(%s%s)", name, distinct, strings.Join(args, ", ")))
	return o.String()
}

/*
      output = []

      # SUM(a, b)
      args = Array(node['args']).map { |arg| deparse_item(arg) }
      # COUNT(*)
      args << '*' if node['agg_star']

      name = (node['funcname'].map { |n| deparse_item(n, FUNC_CALL) } - ['pg_catalog']).join('.')
      distinct = node['agg_distinct'] ? 'DISTINCT ' : ''
      output << format('%s(%s%s)', name, distinct, args.join(', '))
      output << format('OVER %s', deparse_item(node['over'])) if node['over']

	  output.join(' ')
*/
