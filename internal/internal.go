package internal

import (
	"fmt"
	"strings"
)

// SafeNames of table、column、index
func SafeNames(names []string) []string {
	var cooked []string
	for _, v := range names {
		cooked = append(cooked, SafeName(v))
	}

	return cooked
}

// SafeName of table、column、index
func SafeName(column string) string {
	switch {
	case column == "":
		return ""
	case column == "*",
		strings.Index(column, "(") != -1,
		strings.Index(column, " ") != -1:
		return column
	default:
	}

	return fmt.Sprintf("`%s`", strings.Trim(column, "`"))
}
