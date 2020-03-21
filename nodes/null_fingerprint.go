package pg_query_nodes

func (node Null) Fingerprint(ctx FingerprintContext, parentNode Node, parentFieldName string) {
	// Intentionally ignoring all fields for fingerprinting
}
