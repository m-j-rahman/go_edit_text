package gr

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func ReadFile() string {
	fileOne, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Invalid input")
	}
	return string(fileOne)
}

func ToHex(s string) string {
	result, _ := strconv.ParseInt(s, 16, 64)
	return fmt.Sprint(result)
}

func ToBin(s string) string {
	result, _ := strconv.ParseInt(s, 2, 64)
	return fmt.Sprint(result)
}

func AVowel(s string) string {
	matched, _ := regexp.MatchString("^[aeiouh]", s)
	if matched {
		return "an"
	}
	return "a"
}
