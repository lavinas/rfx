package util

import (
	"strconv"
	"strings"
)

func GetCell(row []string, idx int) string {
	if idx >= len(row) {
		return ""
	}

	return strings.TrimSpace(row[idx])
}

func ParseFloat(s string) float64 {
	s = strings.TrimSpace(s)
	if s == "" || s == "-" || s == "-.-" {
		return 0
	}

	if strings.Contains(s, ",") && strings.Contains(s, ".") {
		s = strings.ReplaceAll(s, ".", "")
		s = strings.ReplaceAll(s, ",", ".")
	} else if strings.Contains(s, ",") {
		s = strings.ReplaceAll(s, ",", ".")
	}

	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return v
}

func NormalizeParent(s string) string {

	s = strings.TrimSpace(s)

	if s == "-" {
		return ""
	}

	return s
}

func NormalizeNumericString(s string) string {

	s = strings.TrimSpace(s)

	if before, ok := strings.CutSuffix(s, ".0"); ok {
		s = before
	}

	return s
}
