package sqlstr_test

import (
	"testing"

	"github.com/imantung/sqlstr"
	"github.com/stretchr/testify/require"
)

func TestQueryString_After(t *testing.T) {
	testcase := []struct {
		query   string
		word    string
		atAfter string
	}{
		{
			"SELECT * FROM table",
			"FROM",
			"table",
		},
		{
			"SELECT * FROM        table",
			"FROM",
			"table",
		},
		{
			"select * from table;",
			"FROM",
			"table",
		},
		{
			"SELECT * FROM table WHERE true;",
			"FROM",
			"table",
		},
		{
			"SELECT * FROM table1 WHERE true;",
			"FROM",
			"table1",
		},
		{
			"select * from table1, table2, table3;",
			"FROM",
			"table1",
		},
	}

	for _, tt := range testcase {
		queryString := sqlstr.NewQueryString(tt.query)
		require.Equal(t, tt.atAfter, queryString.After(tt.word))
	}
}

func TestQueryString_AfterAll(t *testing.T) {
	testcase := []struct {
		query    string
		word     string
		atAfters []string
	}{
		{
			`SELECT column_name(s) FROM table1
			UNION
			SELECT column_name(s) FROM table2;`,
			"FROM",
			[]string{"table1", "table2"},
		},
	}

	for _, tt := range testcase {
		queryString := sqlstr.NewQueryString(tt.query)
		require.Equal(t, tt.atAfters, queryString.AfterAll(tt.word))
	}

}

func TestQueryString_TableNames(t *testing.T) {
	testcase := []struct {
		query      string
		tableNames []string
	}{
		{
			"SELECT * FROM table1, table2, table3 WHERE true;",
			[]string{"table1", "table2", "table3"},
		},
		{
			`SELECT * FROM table1, table2, table3 WHERE true
			UNION
			SELECT * FROM table4, table5, table6 WHERE true`,
			[]string{"table1", "table2", "table3", "table4", "table5", "table6"},
		},
		{
			`SELECT * FROM table1, table2, table3;`,
			[]string{"table1", "table2", "table3"},
		},
		{
			`SELECT * FROM table1, table2, table3`,
			[]string{"table1", "table2", "table3"},
		},
		{
			`SELECT column_name(s)
		FROM table1
		LEFT JOIN table2
		ON table1.column_name = table2.column_name;`,
			[]string{"table1", "table2"},
		},
	}

	for _, tt := range testcase {
		queryString := sqlstr.NewQueryString(tt.query)
		require.Equal(t, tt.tableNames, queryString.TableNames())
	}
}
