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
	{
		name: `select with where, order_by, and limit`,
		sql: `SELECT name AS n FROM queues
			WHERE queue_version_id IS NOT NULL
			ORDER BY created_at DESC
			LIMIT 5`,
		expected: `SELECT "name" AS n FROM "queues" WHERE "queue_version_id" IS NOT NULL ORDER BY "created_at" DESC LIMIT 5`,
	},
	{
		name: `select with a join`,
		sql: `SELECT q.id, i.id
		FROM work_items i
		JOIN queues q ON i.queue_id = q.id
		WHERE grouped_entity_type = 'Shipment'
		LIMIT 5`,
		expected: `SELECT "q"."id", "i"."id" FROM "work_items" i JOIN "queues" q ON "i"."queue_id" = "q"."id" WHERE "grouped_entity_type" = 'Shipment' LIMIT 5`,
	},
	{
		name: `select with a CTE and subquery`,
		sql: `WITH q AS (
			SELECT name, id
			FROM queues
			WHERE queue_key IS NOT NULL
			ORDER BY created_at DESC
			LIMIT 50
		)
		SELECT id, queue_id
		FROM work_items
		WHERE queue_id IN (SELECT id FROM q)`,
		expected: `TODO`,
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
