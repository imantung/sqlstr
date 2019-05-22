package sqlstr

import (
	"regexp"
)

// Obscure value in SQL Query
func Obscure(query string) string {
	query = regexp.MustCompile(`\s'(.*?)'|\s(true|TRUE)|\s(false|FALSE)|\s[0-9]+\.[0-9]+|\s[0-9]+`).ReplaceAllString(query, " ?")

	return query
}
