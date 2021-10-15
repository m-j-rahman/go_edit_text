package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]

	if args[0] == "sample.txt" {
		file, _ := ioutil.ReadFile(args[0])
		str := strings.Split(string(file), " ")

		for i, word := range str {
			if word == "(up)" {
				str[i-1] = strings.ToUpper(str[i-1])
				str = append(str[:i], str[i+1:]...)
			} else if word == "(low)" {
				str[i-1] = strings.ToLower(str[i-1])
				str = append(str[:i], str[i+1:]...)
			} else if word == "(cap)" {
				str[i-1] = strings.Title(str[i-1])
				str = append(str[:i], str[i+1:]...)
			} else if word == "(hex)" {
				hexnum := str[i-1]
				numb1, err := strconv.ParseInt(hexnum, 16, 64)

				if err != nil {
					panic(err)
				} else {
					hexnum = string(numb1)
					str = append(str[:i], str[i+1:]...)
				}
			} else if word == "(bin)" {
				binnum := str[i-1]
				numb2, err := strconv.ParseInt(binnum, 2, 64)
				if err != nil {
					panic(err)
				} else {
					binnum = string(numb2)
					str = append(str[:i], str[i+1:]...)
				}
			}
		}
		fmt.Println(str)
	}
}
