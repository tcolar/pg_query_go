package pg_query_nodes

func (node UnlistenStmt) Fingerprint(ctx FingerprintContext, parentNode Node, parentFieldName string) {
	ctx.WriteString("UnlistenStmt")

	if node.Conditionname != nil {
		ctx.WriteString("conditionname")
		ctx.WriteString(*node.Conditionname)
	}
}
