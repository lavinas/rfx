package util

import (
	"math"
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
	if s == "--" {
		return math.NaN()
	}

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

	if s == "-" || s == "--" || s == "" || s == "0" {
		return ""
	}

	return s
}

func Round2(v float64) float64 {
	return math.Round(v*100) / 100
}

func TruncInt(s string) string {
	if idx := strings.Index(s, "."); idx != -1 {
		return s[:idx]
	}
	return s
}