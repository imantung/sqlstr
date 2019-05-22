package sqlstr

import (
	"regexp"
	"strings"
	"unicode"
)

type QueryString struct {
	query   string
	lowered string
}

func NewQueryString(query string) *QueryString {
	query = Clean(query)
	return &QueryString{
		query:   query,
		lowered: strings.ToLower(query),
	}
}

// After word
func (p QueryString) After(word string) string {
	iWord := strings.Index(p.lowered, strings.ToLower(word)) + len(word) + 1
	return p.after(iWord)
}

// AfterAll word
func (p QueryString) AfterAll(word string) (atAfters []string) {
	indices := regexp.MustCompile(strings.ToLower(word)).
		FindAllStringIndex(p.lowered, -1)
	for _, index := range indices {
		atAfters = append(atAfters, p.after(index[1]))
	}
	return
}

func (p QueryString) after(iWord int) (atAfter string) {
	iAfter := 0

	for i := iWord; i < len(p.lowered); i++ {
		r := rune(p.lowered[i])
		if unicode.IsLetter(r) && iAfter <= 0 {
			iAfter = i
		}

		if (unicode.IsSpace(r) || unicode.IsPunct(r)) && iAfter > 0 {
			atAfter = p.query[iAfter:i]
			break
		}
	}

	if atAfter == "" {
		atAfter = p.query[iAfter:]
	}

	return

}

func (p QueryString) TableNames() (names []string) {

	firstSyntax := p.lowered[:strings.IndexRune(p.lowered, ' ')]

	switch firstSyntax {
	case "update":
	case "insert":
	case "delete":
	}

	names = append(names, p.tableNamesByFROM()...)
	names = append(names, p.AfterAll("join")...)

	return
}

func (p QueryString) tableNamesByFROM() (names []string) {
	indices := regexp.MustCompile("from(.*?)where|from(.*?)left|from(.*?)right|from(.*?)inner|from(.*?)outer|from(.*?)full|from(.*?)join|from(.*?);|from(.*?)$").
		FindAllStringIndex(p.lowered, -1)

	for _, index := range indices {

		for _, field := range strings.Fields(p.query[index[0]:index[1]]) {
			loweredField := strings.ToLower(field)
			if loweredField == "from" || loweredField == "where" || loweredField == "left" ||
				loweredField == "right" || loweredField == "join" || loweredField == "inner" ||
				loweredField == "outer" || loweredField == "full" {
				continue
			}
			names = append(names, cleanField(field))
		}
	}

	return
}

func cleanField(field string) string {
	field = strings.TrimSpace(field)
	lastRune := field[len(field)-1]
	if lastRune == ';' || lastRune == ',' {
		field = field[:len(field)-1]
	}

	return field
}
