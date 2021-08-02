package functions

func Split(s, sep string) []string {
	var str = []string{}
	var begin int
	length := len(sep)
	for index := 0; index < len(s)-length; index++ {
		if s[index:index+length] == sep {
			// if s[begin:index] != "" {
			str = append(str, s[begin:index])
			// }
			begin = index + length
			index = index + length - 1
		}
	}
	if sep != s[begin:] {
		str = append(str, s[begin:])
	}
	return str
}
