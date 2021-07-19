package functions

func Split(s, sep string) []string {
	var str = []string{}
	new := ""
	check := ""
	length := len(sep)
	separ := false
	count := 0
	for index, char := range s {
		if char == 'H' {
			for index1, val := range s[index:(length + index)] {
				if index1 == (length - 1) {
					check = check + string(val)
					if check == sep {
						check = ""
						separ = true
						if len(new) > 0 {
							str = append(str, new)
						}
						new = ""
						break
					} else {
						check = ""
						separ = false
						new = new + string(char)
						break
					}
					break
				} else {
					check = check + string(val)
				}
			}
		} else if separ == true {
			count++
			if count == (length - 1) {
				separ = false
				count = 0
			}
		} else if index == (len(s) - 1) {
			new = new + string(char)
			str = append(str, new)
		} else {
			new = new + string(char)

		}
	}
	return str
}
