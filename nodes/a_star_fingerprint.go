package pg_query_nodes

func (node A_Star) Fingerprint(ctx FingerprintContext, parentNode Node, parentFieldName string) {
	ctx.WriteString("A_Star")
}
