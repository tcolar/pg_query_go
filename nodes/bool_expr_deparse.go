package pg_query_nodes

import (
	"fmt"

	"github.com/lfittl/pg_query_go/deparse"
)

func (node BoolExpr) Deparse(ctx deparse.Context) string {
	switch node.Boolop {
	case AND_EXPR:
		return node.deparseBoolExprAnd(ctx)
	case OR_EXPR:
		return node.deparseBoolExprOr(ctx)
	default:
		panic(fmt.Sprintf("Unexpected boolop : %v", node.Boolop))
		//return fmt.Sprintf("NOT %s", node.Args.Items[0].Deparse(nil))
	}
}

func (node BoolExpr) deparseBoolExprAnd(ctx deparse.Context) string {
	return node.Args.DeparseSep(nil, " AND ")
	/*for _, arg := range node.Args.Items {
		arg.Deparse
	}*/
}

func (node BoolExpr) deparseBoolExprOr(ctx deparse.Context) string {
	return node.Args.DeparseSep(nil, " OR ")
	/*for _, arg := range node.Args.Items {
		arg.Deparse
	}*/
}

/*
def deparse_bool_expr_and(node)
# Only put parantheses around OR nodes that are inside this one
node['args'].map do |arg|
  if [BOOL_EXPR_OR].include?(arg.values[0]['boolop'])
	format('(%s)', deparse_item(arg))
  else
	deparse_item(arg)
  end
end.join(' AND ')
end

def deparse_bool_expr_or(node)
# Put parantheses around AND + OR nodes that are inside
node['args'].map do |arg|
  if [BOOL_EXPR_AND, BOOL_EXPR_OR].include?(arg.values[0]['boolop'])
	format('(%s)', deparse_item(arg))
  else
	deparse_item(arg)
  end
end.join(' OR ')
end
*/
