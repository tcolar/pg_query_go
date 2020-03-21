// Auto-generated - DO NOT EDIT

package pg_query_nodes

func (node VariableShowStmt) Fingerprint(ctx FingerprintContext, parentNode Node, parentFieldName string) {
	ctx.WriteString("VariableShowStmt")

	if node.Name != nil {
		ctx.WriteString("name")
		ctx.WriteString(*node.Name)
	}
}
