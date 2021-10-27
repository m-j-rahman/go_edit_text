package main

import (
	"fmt"
	"gr"
	"strings"
)

func main() {
	sliceStr := strings.Split(gr.ReadFile(), " ")

	for i, word := range sliceStr {
		switch word {
		case "(hex)":
			sliceStr[i-1] = gr.ToHex(sliceStr[i-1])
			sliceStr = append(sliceStr[:i], sliceStr[i+1:]...)
		case "(bin)":
			sliceStr[i-1] = gr.ToBin(sliceStr[i-1])
			sliceStr = append(sliceStr[:i], sliceStr[i+1:]...)
		case "(up)":
			sliceStr[i-1] = strings.ToUpper(sliceStr[i-1])
			sliceStr = append(sliceStr[:i], sliceStr[i+1:]...)
		case "(low)":
			sliceStr[i-1] = strings.ToLower(sliceStr[i-1])
			sliceStr = append(sliceStr[:i], sliceStr[i+1:]...)
		case "(cap)":
			sliceStr[i-1] = strings.Title(strings.ToLower(sliceStr[i-1]))
			sliceStr = append(sliceStr[:i], sliceStr[i+1:]...)
		case "a":
			sliceStr[i-1] = gr.AVowel(sliceStr[i+1])
		case "A":
			sliceStr[i-1] = gr.AVowel(sliceStr[i+1])
		}
	}
	fmt.Println(sliceStr)
}
