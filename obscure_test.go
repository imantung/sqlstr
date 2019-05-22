package sqlstr_test

import (
	"testing"

	"github.com/imantung/sqlstr"
	"github.com/stretchr/testify/require"
)

func TestObscure(t *testing.T) {
	testcase := []struct {
		query    string
		obscured string
	}{
		{
			"SELECT * FROM table WHERE s = 'text' AND i = 12345",
			"SELECT * FROM table WHERE s = ? AND i = ?",
		},
		{
			"SELECT * FROM table WHERE b1 = true AND b2 = false",
			"SELECT * FROM table WHERE b1 = ? AND b2 = ?",
		},
		{
			"SELECT * FROM table WHERE disc = 0.8",
			"SELECT * FROM table WHERE disc = ?",
		},
	}

	for _, tt := range testcase {
		require.Equal(t, tt.obscured, sqlstr.Obscure(tt.query))
	}
}
