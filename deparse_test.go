package pg_query_test

import (
	"testing"

	pg_query "github.com/lfittl/pg_query_go"
)

type testCase struct {
	name     string // test name
	sql      string // SQL query
	expected string // Expected Deparsed query (normalized)
}

var selectTests = []testCase{
	{
		name:     `simple select`,
		sql:      ` SELECT id,  name  FROM  queues WHERE   id = '00A10DC4-00B1-433E-B97E-2FE04A07E62E'  `,
		expected: `SELECT "id", "name" FROM "queues" WHERE "id" = '00A10DC4-00B1-433E-B97E-2FE04A07E62E'`,
	},
}

func TestDeparseSelects(t *testing.T) {
	for _, test := range selectTests {
		tree, err := pg_query.Parse(test.sql)
		if err != nil {
			panic(err)
		}
		deparsed := tree.Deparse(nil)
		if deparsed != test.expected {
			t.Fatalf(`Failed test %s
			Got: %s
			Expected: %s`, test.name, deparsed, test.expected)
		}
	}
}
