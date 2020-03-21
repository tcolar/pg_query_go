// Auto-generated - DO NOT EDIT

package pg_query_nodes

func (node ReplicaIdentityStmt) Fingerprint(ctx FingerprintContext, parentNode Node, parentFieldName string) {
	ctx.WriteString("ReplicaIdentityStmt")

	if node.IdentityType != 0 {
		ctx.WriteString("identity_type")
		ctx.WriteString(string(node.IdentityType))

	}

	if node.Name != nil {
		ctx.WriteString("name")
		ctx.WriteString(*node.Name)
	}
}
