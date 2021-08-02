package functions

func Split(s, sep string) []string {
	var str = make([]string, 0)
	var begin int
	length := len(sep)
	if sep == "" {
		for _, val := range s {
			str = append(str, string(val))
		}
		return str
	}
	for index := 0; index <= len(s)-length; index++ {
		if s[index:index+length] == sep {
			str = append(str, s[begin:index])
			begin = index + length
			index = index + length - 1
		}
	}
	if sep != s[begin:] {
		str = append(str, s[begin:])
	}
	return str
}
