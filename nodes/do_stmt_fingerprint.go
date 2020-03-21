package pg_query_nodes

func (node DoStmt) Fingerprint(ctx FingerprintContext, parentNode Node, parentFieldName string) {
	ctx.WriteString("DoStmt")

	if len(node.Args.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Args.Fingerprint(&subCtx, node, "Args")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("args")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
