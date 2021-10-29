package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// >>> Removes a rune from a slice of rune at index s
func runeremove(slice []rune, s int) []rune {
	return append(slice[:s], slice[s+1:]...)
}

// >>> Removes a string from a slice of string at index s
func stringremove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

// >>> Goes through the string and converts numbers to int
func TrimAtoi(s string) int {
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

// >>> Parses the modification and executes the correct amount of times
func parseinst(srune []rune, index int, strarr []string) []string {
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
			strarr[index-i] = strings.ToUpper(strarr[index-i])
		}
		if tempinst == "low" {
			strarr[index-i] = strings.ToLower(strarr[index-i])
		}
		if tempinst == "cap" {
			strarr[index-i] = strings.Title(strarr[index-i])
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

func main() {
	var strarr []string
	stringFromArr := ""
	var found int = 3
	quotePair := false
	prev := ""

	// >>> Check to make sure there are 2 arguments
	args := os.Args[1:]
	if len(args) < 2 || len(args) > 2 {
		fmt.Println("invalid number of arguments")
		os.Exit(0)
	}

	// >>> Opens the file
	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// >>> Split the input by spaces, newlines and returns
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		// >>> Checks the output of scanwords then reassembles them (because of the space in  the (cap, 6), modifications are being split)
		for _, i := range scanner.Text() {
			if i == '(' || found == 1 {
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
	// >>> Checks for modifications and applies them with parseinst fuction
	for i := 0; i < len(strarr); i++ {
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
			// >>> Checks for a single quotation
			if strarr[i][dex] == '\'' && dex == 0 && len(strarr[i]) == 1 {
				if !quotePair {
					// >>> If it's the first one found
					quotePair = true
					strarr[i+1] = "'" + strarr[i+1]
					strarr = stringremove(strarr, i)
					i--
				} else {
					// >>> If it's the second one found
					quotePair = false
					strarr[i-1] += "'"
					strarr = stringremove(strarr, i)
					i--
				}
			}
			// >>> Checks for punctuation
			if (strarr[i][dex] == '.' || strarr[i][dex] == ',' || strarr[i][dex] == '!' ||
				strarr[i][dex] == '?' || strarr[i][dex] == ':' || strarr[i][dex] == ';') && dex == 0 {
				// >>> Checks to see if there is any other characters with it
				if len(strarr[i])-1 == 0 && dex == 0 {
					tempRuneSlice := []rune(strarr[i])
					prevElement = []rune(strarr[i-1])
					prevElement = append(prevElement, tempRuneSlice[dex])
					strarr = stringremove(strarr, i)
					i--
					strarr[i] = string(prevElement)
				} else if len(strarr[i])-1 > 0 && dex == 0 {
					// >>> Checks if it is multiple punctuations
					for u := 0; u < len(strarr[i]); u++ {
						if strarr[i][u] == '.' || strarr[i][u] == ',' || strarr[i][u] == '!' ||
							strarr[i][u] == '?' || strarr[i][u] == ':' || strarr[i][u] == ';' {
							count++
						}
					}
					// >>> If only one punctuation present, remove from current word and move to previous word
					if count == 1 {
						tempRuneSlice := []rune(strarr[i])
						prevElement = []rune(strarr[i-1])
						prevElement = append(prevElement, tempRuneSlice[dex])
						tempRuneSlice = runeremove(tempRuneSlice, dex)
						strarr[i] = string(tempRuneSlice)
						strarr[i-1] = string(prevElement)
					}
					// >>> If all are punctuations, then remove all and move to previous word and delete this instance
					if count == len(strarr[i]) {
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
	// >>> Check for (a) and if followed by (a,e,i,o,u or h) add (n) to the (a)
	for i := 0; i < len(strarr); i++ {
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
	// >>> Reassembles the string for writing
	for lenStr, str := range strarr {
		if lenStr > 0 && lenStr < len(strarr) {
			stringFromArr += " "
		}
		stringFromArr += str
	}
	// >>> Open and write the file
	f, err := os.Create(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(stringFromArr)
	f.Sync()
}
