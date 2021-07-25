package main

import (
	"fmt"
	"io"
	"os"
	"student/functions"
)

func main() {
	arguments := os.Args
	// fmt.Println(len(arguments))
	arg, number, arr := IsValid(arguments)
	// fmt.Println(number, arr, arg)
	if !arg {
		os.Stdout.ReadFrom(os.Stdin)
	} else {
		for _, value := range arr {
			info, err := os.Open(value)
			if err != nil {
				fmt.Println("ztail:")
				// os.Stdout.Write(err)

			}
			defer info.Close()
			data := make([]byte, 64)
			for {
				values, err := info.Read(data)
				if err == io.EOF {
					break
				}
				fmt.Printf("%s", string(data[values-number:values]))
				// length := len(data)
				// fmt.Println(string(data)[length-number : length])
			}
			fmt.Printf("==> %s <==", value)
		}
	}

	// arguments := os.Args[1:]
	// arg, number, arr := IsValid(arguments)
	// if !arg {
	// 	os.Stdout.ReadFrom(os.Stdin)
	// } else {
	// 	for _, value := range arr {
	// 		info, err := ioutil.ReadFile(value)
	// 		if err != nil {
	// 			fmt.Println("ERROR:" + value)
	// 		}
	// 		length := len(info)
	// 		fmt.Println(string(info)[length-number:])
	// 	}
	// }
}

func IsValid(arguments []string) (bool, int, []string) {
	var str []string
	var result int
	if len(arguments) == 0 {
		return false, result, str
	}
	for index := 0; index < len(arguments); index++ {
		// fmt.Println(index)
		if arguments[index] == "-c" && index != len(arguments) {
			if arguments[index+1] == "0" {
				index = index + 1
				result = 0
			} else if functions.Atoi(arguments[index+1]) != 0 {
				result = functions.Atoi(arguments[index+1])
				index = index + 1
			} else {
				return false, result, str
			}
		} else {
			str = append(str, arguments[index])
		}
	}

	return true, result, str
}
