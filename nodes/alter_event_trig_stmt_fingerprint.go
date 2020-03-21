package pg_query_nodes

func (node AlterEventTrigStmt) Fingerprint(ctx FingerprintContext, parentNode Node, parentFieldName string) {
	ctx.WriteString("AlterEventTrigStmt")

	if node.Tgenabled != 0 {
		ctx.WriteString("tgenabled")
		ctx.WriteString(string(node.Tgenabled))

	}

	if node.Trigname != nil {
		ctx.WriteString("trigname")
		ctx.WriteString(*node.Trigname)
	}
}
