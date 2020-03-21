package pg_query_nodes

import "strconv"

func (node DiscardStmt) Fingerprint(ctx FingerprintContext, parentNode Node, parentFieldName string) {
	ctx.WriteString("DiscardStmt")

	if int(node.Target) != 0 {
		ctx.WriteString("target")
		ctx.WriteString(strconv.Itoa(int(node.Target)))
	}
}
