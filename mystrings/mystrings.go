package mystrings

func Reverse (s *string) string {
	str := *s
	result := ""

	if str != "" {
	for _,v := range str {
		result = string(v) + result
	}
	*s = result
		return result
	}
	return "Nothing present to reverse"
}
