package pg_query_nodes

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lfittl/pg_query_go/deparse"
)

func (node TypeName) Deparse(ctx deparse.Context) string {
	var names []string
	for _, i := range node.Names.Items {
		names = append(names, i.Deparse(ctx))
	}
	if len(names) == 2 && names[0] == "pg_catalog" && names[1] == "interval" {
		// Intervals are tricky and should be handled in a separate method because
		// they require performing some bitmask operations.
		return node.DeparseIntervaleType()
	}
	o := deparse.Output{}
	if node.Setof {
		o.Append("SETOF")
	}
	args := ""
	if !node.Typmods.Empty() {
		args = node.Typmods.Deparse(nil)
	}
	o.Append(node.DeparseTypeNameCast(names, args))
	if !node.ArrayBounds.Empty() {
		o.Append("[]")
	}
	return o.String()
}

func (node TypeName) DeparseTypeNameCast(names []string, args string) string {
	catalog := names[0]
	type_ := names[1]

	if catalog != "pg_catalog" {
		// Just pass along any custom types.
		// (The pg_catalog types are built-in Postgres system types and are
		// handled in the case statement below)
		custom := strings.Join(names, ".")
		if len(args) > 0 {
			custom = fmt.Sprintf("%s(%s)", custom, args)
		}
		return custom
	}

	switch type_ {
	case "bpchar":
		// char(2) or char(9)
		return fmt.Sprintf("char(%s)", args)
	case "varchar":
		if len(args) == 0 {
			return "varchar"
		}
		return fmt.Sprintf("varchar(%s)", args)
	case "numeric'":
		if len(args) == 0 {
			return "numeric"
		}
		// numeric(3, 5)
		return fmt.Sprintf("numeric(%s)", args)
	case "bool":
		return "boolean"
	case "int2":
		return "smallint"
	case "int4":
		return "int"
	case "int8":
		return "bigint"
	case "real", "float4":
		return "real"
	case "float8":
		return "double precision"
	case "time":
		return "time"
	case "timetz":
		return "time with time zone"
	case "timestamp":
		return "timestamp"
	case "timestamptz":
		return "timestamp with time zone"
	default:
		panic(fmt.Sprintf("Can't deparse type: %s", type_))
	}
}

// Deparses interval type expressions like `interval year to month` or
// `interval hour to second(5)`
func (node TypeName) DeparseIntervaleType() string {
	types := []string{"interval"}
	if !node.Typmods.Empty() {
		var typmods []string
		for _, i := range node.Typmods.Items {
			typmods = append(typmods, i.Deparse(nil))
		}
		intervalIndex, _ := strconv.Atoi(typmods[0])
		intervals := []string{}
		for _, interval := range deparse.IntervalHelper.FromInt(uint(intervalIndex)) {
			// only the `second` type can take an argument.
			if interval == "second" && len(typmods) == 2 {
				intervals = append(intervals, fmt.Sprintf("second(%s)", strings.ToLower(typmods[1])))
			} else {
				intervals = append(intervals, strings.ToLower(interval))
			}
		}
		types = append(types, strings.Join(intervals, " to "))
	}

	return strings.Join(types, " ")
}
