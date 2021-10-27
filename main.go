package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	GoReloaded()
}

// >>> Removes a rune from a slice of rune at index s
func runeremove(slice []rune, s int) []rune {
	return append(slice[:s], slice[s+1:]...)
}

// >>> Removes a string from a slice of string at index s
func stringremove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

// >>> Lower-cases all characters in a string
func Lower(s string) string {
	var lower []rune
	string1 := []rune(s)
	for c := range string1 {
		if string1[c] >= 'A' && string1[c] <= 'Z' {
			lower = append(lower, string1[c]+32)
		} else {
			lower = append(lower, string1[c])
		}
	}
	return string(lower)
}

// >>> Capitalises all characters in a string
func Upper(s string) string {
	var upper []rune
	string1 := []rune(s)
	for c := range string1 {
		if string1[c] >= 'a' && string1[c] <= 'z' {
			upper = append(upper, string1[c]-32)
		} else {
			upper = append(upper, string1[c])
		}
	}
	return string(upper)
}

// >>> Capitalises only 1st character of string
func Cap(s string) string {
	s = Lower(s)
	sRune := []rune(s)
	for i, r := range sRune {
		if r >= 'a' && r <= 'z' {
			sRune[i] = r - 32
			return string(sRune)
		}
	}
	return string(sRune)
}

// <<<<<<<<<<< MAIN FUNCTION BEGINS HERE >>>>>>>>>>>
// <<<<<<<<<<< MAIN FUNCTION BEGINS HERE >>>>>>>>>>>
// <<<<<<<<<<< MAIN FUNCTION BEGINS HERE >>>>>>>>>>>
func GoReloaded() {
	var strarr []string
	stringFromArr := ""
	var found int = 3
	quotePair := false
	prev := ""
	args := os.Args[1:]
	if len(args) < 2 || len(args) > 2 { // >>> Check to make sure there are 2 arguments
		fmt.Println("too many arguments or too few arguments")
		os.Exit(0)
	}
	file, err := os.Open(args[0]) // >>> Opens the file
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords) // >>> Split the input by spaces, newlines and returns
	for scanner.Scan() {
		for _, i := range scanner.Text() {
			if i == '(' || found == 1 { // >>> This part checks the output of scanwords (because of the space in  the (cap, 6) modifications are being split in 2) then reassembles them
				found = 1
				if i == ')' {
					strarr = append(strarr, prev+scanner.Text())
					found = 2
				}
			}
		}
		if found == 3 {
			strarr = append(strarr, scanner.Text())
		}
		if found == 1 {
			prev = scanner.Text()
		}
		if found == 2 {
			found = 3
			prev = ""
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(strarr); i++ { // >>> Checks for modifications and applies them with parseinst fuction
		str := []rune(strarr[i])
		for dex := 0; dex < len(str); dex++ {
			if str[dex] == '(' && dex == 0 {
				strarr = parseinst(str, i, strarr)
				str = []rune(strarr[i])
			}
		}
	}
	for i := 0; i < len(strarr); i++ {
		var prevElement []rune
		count := 0
		for dex := 0; dex < len(strarr[i]); dex++ {
			if strarr[i][dex] == '\'' && dex == 0 && len(strarr[i]) == 1 { // >>> Checks for a single quotation
				if !quotePair { // >>> If it's the first one found
					quotePair = true
					strarr[i+1] = "'" + strarr[i+1]
					strarr = stringremove(strarr, i)
					i--
				} else { // >>> If it's the second one found
					quotePair = false
					strarr[i-1] += "'"
					strarr = stringremove(strarr, i)
					i--
				}
			}
			if (strarr[i][dex] == '.' || strarr[i][dex] == ',' || strarr[i][dex] == '!' ||
				strarr[i][dex] == '?' || strarr[i][dex] == ':' || strarr[i][dex] == ';') && dex == 0 { // >>> Checks for punctuation
				if len(strarr[i])-1 == 0 && dex == 0 { // >>> Checks to see if there is any other characters with it
					tempRuneSlice := []rune(strarr[i])
					prevElement = []rune(strarr[i-1])
					prevElement = append(prevElement, tempRuneSlice[dex])
					strarr = stringremove(strarr, i)
					i--
					strarr[i] = string(prevElement)
				} else if len(strarr[i])-1 > 0 && dex == 0 {
					for u := 0; u < len(strarr[i]); u++ { // >>> Checks if it is multiple punctuations
						if strarr[i][u] == '.' || strarr[i][u] == ',' || strarr[i][u] == '!' ||
							strarr[i][u] == '?' || strarr[i][u] == ':' || strarr[i][u] == ';' {
							count++
						}
					}
					if count == 1 { // >>> If only one punctuation present, remove from current word and move to previous word
						tempRuneSlice := []rune(strarr[i])
						prevElement = []rune(strarr[i-1])
						prevElement = append(prevElement, tempRuneSlice[dex])
						tempRuneSlice = runeremove(tempRuneSlice, dex)
						strarr[i] = string(tempRuneSlice)
						strarr[i-1] = string(prevElement)
					}
					if count == len(strarr[i]) { // >>> If all are punctuations, then remove all and move to previous word and delete this instance
						tempstr := []rune(strarr[i-1])
						tempstr2 := []rune(strarr[i])
						tempstr = append(tempstr, tempstr2...)
						strarr[i-1] = string(tempstr)
						strarr = stringremove(strarr, i)
						i--
					}
				}
			}
		}
	}
	for i := 0; i < len(strarr); i++ { // >>> Check for (a) and if followed by (a,e,i,o,u or h) add (n) to the (a)
		srune := []rune(strarr[i])
		for u := 0; u < len(srune); u++ {
			if (srune[u] == 'a' || srune[u] == 'A') && len(srune) == 1 {
				tempsrune := []rune(strarr[i+1])
				if tempsrune[0] == 'a' || tempsrune[0] == 'e' || tempsrune[0] == 'i' || tempsrune[0] == 'o' || tempsrune[0] == 'u' ||
					tempsrune[0] == 'A' || tempsrune[0] == 'E' || tempsrune[0] == 'I' || tempsrune[0] == 'O' || tempsrune[0] == 'U' ||
					tempsrune[0] == 'H' || tempsrune[0] == 'h' {
					srune = append(srune, 'n')
					strarr[i] = string(srune)
				}
			}
		}
	}
	for lenStr, str := range strarr { // >>> Reassembles the string for writing
		if lenStr > 0 && lenStr < len(strarr) {
			stringFromArr += " "
		}
		stringFromArr += str
	}
	f, err := os.Create(os.Args[2]) // >>> Open and write the file
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(stringFromArr)
	f.Sync()
}

func TrimAtoi(s string) int { // >>> Goes through the string and pulls out the numbers and converts them to int
	strrune := []rune(s)
	var ints []rune
	n := 0
	for j := 0; j < len(strrune); j++ {
		if strrune[j] >= '0' && strrune[j] <= '9' {
			ints = append(ints, strrune[j])
		}
	}
	for i := 0; i < len(ints); i++ {
		y := ints[i] - '0'
		n = n*10 + int(y)
	}
	return n
}

func parseinst(srune []rune, index int, strarr []string) []string { // >>> Parses the modification and executes the correct amount of times
	srune = runeremove(srune, 0)
	srune = runeremove(srune, len(srune)-1)
	amount := TrimAtoi(string(srune))
	if amount == 0 {
		amount = 1
	}
	tempinst := ""
	found := false
	for _, i := range srune {
		if i != ',' && !found {
			tempinst += string(i)
		} else {
			found = true
		}
	}
	for i := 0; i <= amount; i++ {
		if tempinst == "up" {
			strarr[index-i] = Upper(strarr[index-i])
		}
		if tempinst == "low" {
			strarr[index-i] = Lower(strarr[index-i])
		}
		if tempinst == "cap" {
			strarr[index-i] = Cap(strarr[index-i])
		}
	}
	if tempinst == "bin" {
		i, _ := strconv.ParseInt(strarr[index-1], 2, 64)
		strarr[index-1] = fmt.Sprint(i)
	}
	if tempinst == "hex" {
		i, _ := strconv.ParseInt(strarr[index-1], 16, 64)
		strarr[index-1] = fmt.Sprint(i)
	}
	strarr = stringremove(strarr, index)
	return strarr
}
