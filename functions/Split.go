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

// func Split(s, sep string) []string {
// 	var res []string
// 	var n int
// 	for i := 0; i < len(s)-len(sep)+1; i++ {
// 		if s[i:len(sep)+i] == sep {
// 			res = append(res, s[n:i])
// 			n = i + len(sep)
// 			if sep != "" { // chto by ne dobavlyzt besconechno pustotu
// 				i = n - 1 // -1 -2 -3
// 			}
// 		}
// 	}
// 	res = append(res, s[n:len(s)])
// 	if sep == "" {
// 		res = res[1 : len(res)-1]
// 	}
// 	return res
// }

// func Split(s, sep string) []string {
// 	var str = []string{}
// 	var begin int
// 	// strLen := len(s)
// 	length := len(sep)
// 	first := false
// 	last := false
// 	fmt.Println(len(s), length)
// 	if len(sep) == 0 || len(sep) > len(s) {
// 		str = append(str, s)
// 		return str
// 	}
// 	for index := 0; index < len(s)-length+1; index++ {
// 		fmt.Println(s[index:index+length], index)
// 		if s[index:index+length] == sep {
// 			if index+length == len(s) {
// 				last = true
// 			}
// 			fmt.Println(s[begin:index])
// 			if s[begin:index] != "" {
// 				if first {
// 					str = append(str, " "+s[begin:index])
// 					first = false
// 				} else if last {
// 					str = append(str, s[begin:index]+" ")
// 				} else {
// 					str = append(str, s[begin:index])
// 				}
// 			} else if last {
// 				str = append(str, " ")
// 			}
// 			if begin == 0 {
// 				first = true
// 			}
// 			begin = index + length
// 		}
// 	}
// 	return str
// }
