package pg_query_nodes

import (
	"fmt"

	"github.com/lfittl/pg_query_go/deparse"
)

func (node A_Expr) Deparse(ctx deparse.Context) string {
	switch node.Kind {
	case AEXPR_OP:
		return node.DeparseAexpr(ctx)
	case AEXPR_OP_ALL:
		return node.DeparseAexprAll(ctx)
	case AEXPR_OP_ANY:
		return node.DeparseAexprAny(ctx)
	case AEXPR_IN:
		return node.DeparseAexprIn(ctx)
	case AEXPR_ILIKE:
		return node.DeparseAexprIlike(ctx)
	// case CONSTR_TYPE_FOREIGN: ???
	case AEXPR_BETWEEN, AEXPR_NOT_BETWEEN, AEXPR_BETWEEN_SYM, AEXPR_NOT_BETWEEN_SYM:
		return node.DeparseAexprBetween(ctx)
	case AEXPR_NULLIF:
		return node.DeparseAexprNullIf(ctx)
	default:
		panic(fmt.Sprintf("Unexpected aexpr type: %v", node.Kind))
	}
}

func (node A_Expr) DeparseAexpr(ctx deparse.Context) string {
	o := deparse.Output{}
	o.Append(node.Lexpr.Deparse(deparse.PassOrEmptyContext(ctx)))
	o.Append(node.Rexpr.Deparse(deparse.PassOrEmptyContext(ctx)))
	sep := node.Name.Items[0].Deparse(deparse.NewContext(deparse.ColumnOperator))
	output := o.Join(fmt.Sprintf(" %s ", sep))
	if ctx != nil {
		// This is a nested expression, add parentheses.
		output = fmt.Sprintf("(%s)", output)
	}
	return output
}

func (node A_Expr) DeparseAexprIn(ctx deparse.Context) string {
	value := node.Rexpr.Deparse(deparse.PassOrEmptyContext(ctx))
	operator := "IN"
	if node.Name.Deparse(nil) != `"="` {
		operator = "NOT IN"
	}
	return fmt.Sprintf("%s %s (%s)", node.Lexpr.Deparse(nil), operator, value)
}

func (node A_Expr) DeparseAexprLike(ctx deparse.Context) string {
	value := node.Rexpr.Deparse(deparse.PassOrEmptyContext(ctx))
	operator := "LIKE"
	if node.Name.Deparse(nil) != `"~~"` {
		operator = "NOT LIKE"
	}
	return fmt.Sprintf("%s %s (%s)", node.Lexpr.Deparse(nil), operator, value)
}

func (node A_Expr) DeparseAexprIlike(ctx deparse.Context) string {
	value := node.Rexpr.Deparse(deparse.PassOrEmptyContext(ctx))
	operator := "ILIKE"
	if node.Name.Deparse(nil) != `"~~*"` {
		operator = "NOT ILIKE"
	}
	return fmt.Sprintf("%s %s (%s)", node.Lexpr.Deparse(nil), operator, value)
}

func (node A_Expr) DeparseAexprAny(ctx deparse.Context) string {
	o := deparse.Output{}
	o.Append(node.Lexpr.Deparse(nil))
	o.Append(fmt.Sprintf("ANY(%s)", node.Rexpr.Deparse(nil)))
	sep := node.Name.Items[0].Deparse(deparse.NewContext(deparse.ColumnOperator))
	return o.Join(fmt.Sprintf(" %s ", sep))
}

func (node A_Expr) DeparseAexprAll(ctx deparse.Context) string {
	o := deparse.Output{}
	o.Append(node.Lexpr.Deparse(nil))
	o.Append(fmt.Sprintf("ALL(%s)", node.Rexpr.Deparse(nil)))
	sep := node.Name.Items[0].Deparse(deparse.NewContext(deparse.ColumnOperator))
	return o.Join(fmt.Sprintf(" %s ", sep))
}

func (node A_Expr) DeparseAexprBetween(ctx deparse.Context) string {
	comparator := ""
	switch node.Kind {
	case AEXPR_BETWEEN:
		comparator = "BETWEEN"
	case AEXPR_NOT_BETWEEN:
		comparator = "NOT BETWEEN"
	case AEXPR_BETWEEN_SYM:
		comparator = "BETWEEN SYMMETRIC"
	case AEXPR_NOT_BETWEEN_SYM:
		comparator = "NOT BETWEEN SYMMETRIC"
	default:
		panic(fmt.Sprintf("Unexpected Between node type: %v", node.Kind))
	}
	name := node.Lexpr.Deparse(nil)
	output := node.Rexpr.(*List).DeparseSep(nil, " AND ")
	return fmt.Sprintf("%s %s %s", name, comparator, output)
}

func (node A_Expr) DeparseAexprNullIf(ctx deparse.Context) string {
	return fmt.Sprintf("NULLIF%s, %s)",
		node.Lexpr.Deparse((nil)),
		node.Rexpr.Deparse((nil)),
	)
}
