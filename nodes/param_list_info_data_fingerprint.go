package pg_query_nodes

import "strconv"

func (node ParamListInfoData) Fingerprint(ctx FingerprintContext, parentNode Node, parentFieldName string) {
	ctx.WriteString("ParamListInfoData")

	if node.NumParams != 0 {
		ctx.WriteString("numParams")
		ctx.WriteString(strconv.Itoa(int(node.NumParams)))
	}

	ctx.WriteString("paramMask")
	for _, val := range node.ParamMask {
		ctx.WriteString(strconv.Itoa(int(val)))
	}
}
