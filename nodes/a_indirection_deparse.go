package pg_query_nodes

import "github.com/lfittl/pg_query_go/deparse"

func (node A_Indirection) Deparse(ctx deparse.Context) string {

	panic("Not Implemented")

	o := deparse.Output{}
	o.Append(node.Arg.Deparse(nil))
	// How to do the ?key thing ??
	o.Append(node.Indirection.Deparse(nil))
	return o.String()
}

/*
      output = []
      arg = deparse_item(node['arg'])
      output << if node['arg'].key?(FUNC_CALL) || node['arg'].key?(SUB_LINK)
                  "(#{arg})."
                else
                  arg
                end
      node['indirection'].each do |subnode|
        output << deparse_item(subnode)
      end
	  output.join
*/
