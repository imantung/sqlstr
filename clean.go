package sqlstr

import (
	"regexp"
	"strings"
)

// Clean SQL Query from double white space, comment, etc.
func Clean(query string) string {

	// remove comment
	query = regexp.MustCompile(`/\*(.*)\*/|\-\-(.*)`).ReplaceAllString(query, "")

	// remove double white space
	query = regexp.MustCompile(`\s+`).ReplaceAllString(query, " ")

	// remove leading/trailing space
	query = strings.TrimSpace(query)

	return query
}
