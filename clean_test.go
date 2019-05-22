package sqlstr_test

import (
	"testing"

	"github.com/imantung/sqlstr"
	"github.com/stretchr/testify/require"
)

func TestClean(t *testing.T) {
	testcase := []struct {
		query   string
		cleaned string
	}{
		{
			"SELECT * \r\nFROM table\r\n",
			"SELECT * FROM table",
		},
		{
			"SELECT *                   FROM               table",
			"SELECT * FROM table",
		},
		{
			`SELECT * /* some comment*/ 
      FROM table; /* another comment */`,
			"SELECT * FROM table;",
		},
		{
			`SELECT * -- some comment 
        FROM table; -- another comment`,
			"SELECT * FROM table;",
		},
	}

	for _, tt := range testcase {
		require.Equal(t, tt.cleaned, sqlstr.Clean(tt.query))
	}

}
