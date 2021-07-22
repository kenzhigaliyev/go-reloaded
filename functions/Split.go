package functions

func Split(s, sep string) []string {
	var str = []string{}
	var begin int
	length := len(sep)
	for index := 0; index < len(s)-length; index++ {
		if s[index:index+length] == sep {
			if s[begin:index] != "" {
				str = append(str, s[begin:index])
			}
			begin = index + length
		}
	}
	return str
}
