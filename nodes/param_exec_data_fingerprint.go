// Auto-generated - DO NOT EDIT

package pg_query_nodes

import "strconv"

func (node ParamExecData) Fingerprint(ctx FingerprintContext, parentNode Node, parentFieldName string) {
	ctx.WriteString("ParamExecData")

	if node.Isnull {
		ctx.WriteString("isnull")
		ctx.WriteString(strconv.FormatBool(node.Isnull))
	}
}
