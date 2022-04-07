package utils

import "strconv"

func ParseFloat64(s string, fallback float64) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fallback
	}
	return f
}
func ParseUint(s string, fallback uint) uint {
	f, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return fallback
	}
	return uint(f)
}
