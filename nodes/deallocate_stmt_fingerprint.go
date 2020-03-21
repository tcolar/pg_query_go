package pg_query_nodes

func (node DeallocateStmt) Fingerprint(ctx FingerprintContext, parentNode Node, parentFieldName string) {
	ctx.WriteString("DeallocateStmt")
	// Intentionally ignoring node.Name for fingerprinting
}
