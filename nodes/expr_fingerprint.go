package pg_query_nodes

func (node Expr) Fingerprint(ctx FingerprintContext, parentNode Node, parentFieldName string) {
	ctx.WriteString("Expr")
}
