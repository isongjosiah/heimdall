package function

import "strconv"

// StringToInt converts a string to an int. If the string is empty, it returns 0
func StringToInt(s string) int {
	if s == "" {
		return 0
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
