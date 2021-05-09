package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
)

func main() {
	////////01234567890123456789010123456789012345678901234567890
	byt,err := ioutil.ReadFile("./input.txt")
	if err!= nil {
		fmt.Println("error reading from file",err)
		os.Exit(1)
	}

	str := string(byt)
	strcp := string(byt)

	pattern := os.Args[1]
	count := 0

	//bad match table will be a map which will contain skip values for each of the runes from patter
	//for the non- matching runes, we'll skip the length directly in the logic
	//for last character, the value will be length
	bm := make(map[rune]int)

	l := len(pattern)

	for i, c := range pattern {
		bm[c] = l - i - 1
		if i == l-1 {
			bm[c] = l
		}
	}

	//cursor starts from the end of the pattern as the matching is always done from right to left
	cursor := l - 1
	index := 0

LOOP:
	if str[cursor] == pattern[cursor] {
		x := bm[rune(str[cursor])]
		for i := cursor; i >= 0; i-- {
			if str[i] == pattern[i] {
				count++
			} else {
				count = 0
				str = str[x:]
				index = index + x
				goto LOOP
			}
		}
		if count == l {
			word := color.RedString(str[:l])
			result := fmt.Sprintf("%s %s %s",strcp[:index-l],word,strcp[index +1:])
			fmt.Println(result)
			return
		}
		fmt.Println("not found")
	}

	if str[cursor] != pattern[cursor] {

		if bm[rune(str[cursor])] == 0 {
			str = str[l:]
			index = index + l

		} else {
			trim := bm[rune(str[cursor])]
			str = str[trim:]
			index = index + bm[rune(str[cursor])]

		}
		goto LOOP
	}

}
