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

// TableNames of Query
func (p QueryString) TableNames() (names []string) {
	firstSyntax := p.lowered[:strings.IndexRune(p.lowered, ' ')]

	switch firstSyntax {
	case "update":
		names = append(names, cleanName(p.After("update")))
		return
	case "insert":
		index := regexp.MustCompile("insert(.*?)into").FindStringIndex(p.lowered)
		names = append(names, cleanName(p.after(index[1])))
		return
	case "delete":
		index := regexp.MustCompile("delete(.*?)from").FindStringIndex(p.lowered)
		names = append(names, cleanName(p.after(index[1])))
		return
	}

	names = append(names, p.tableNamesByFROM()...)
	names = append(names, p.AfterAll("join")...)

	return
}

func (p QueryString) tableNamesByFROM() (names []string) {
	indices := regexp.MustCompile("from(.*?)(left|inner|right|outer|full)|from(.*?)join|from(.*?)where|from(.*?);|from(.*?)$").
		FindAllStringIndex(p.lowered, -1)

	for _, index := range indices {
		fromStmt := p.lowered[index[0]:index[1]]
		lastSyntax := fromStmt[strings.LastIndex(fromStmt, " ")+1:]

		var tableStmt string
		if lastSyntax == "from" || lastSyntax == "where" || lastSyntax == "left" ||
			lastSyntax == "right" || lastSyntax == "join" || lastSyntax == "inner" ||
			lastSyntax == "outer" || lastSyntax == "full" {
			tableStmt = p.query[index[0]+len("from")+1 : index[1]-len(lastSyntax)-1]
		} else {
			tableStmt = p.query[index[0]+len("from")+1:]
		}

		for _, name := range strings.Split(tableStmt, ",") {
			names = append(names, cleanName(name))
		}
	}
	return
}

func cleanName(name string) string {
	name = strings.Fields(name)[0]
	name = strings.TrimSpace(name)
	lastRune := name[len(name)-1]
	if lastRune == ';' {
		name = name[:len(name)-1]
	}
	return name
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
