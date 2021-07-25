package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) == 0 {
		os.Stdout.ReadFrom(os.Stdin)
	} else {
		for _, value := range arguments {
			info, err := ioutil.ReadFile(value)
			if err != nil {
				fmt.Println("ERROR:" + value)
			}
			fmt.Print(string(info))
		}
	}
}
