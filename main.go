package main

import "fmt"

func main() {
	////////01234567890123456789010123456789012345678901234567890
	str := `any sentence will do apples are not oranges but bananas they aare mangoes`
	pattern := "are"
	count := 0

	//bad match table will be a map which will contain skip values for each of the runes from patter
	//for the non- matching runes, we'll skip the length directly in the logic
	//for last character, the value will be length
	bm := make(map[rune]int)

	l := len(pattern)

	for i,c := range pattern {
		bm[c] = l - i - 1
		if i == l - 1 {
			bm[c] = l
		}
	}

	//cursor starts from the end of the pattern as the matching is always done from right to left
	cursor := l - 1
	//trimmer := 0

	for j := 0; j < len(str); j++ {

		if str[cursor] != pattern[cursor] {

			if bm[rune(str[cursor])] == 0 {
				str = str[l:]
				//trimmer = trimmer + l
				//j = j + l
				continue
			}
			//trim str
			//j = j  + bm[rune(str[cursor])]
			indexToTrim  := bm[rune(str[cursor])]
			str = str[indexToTrim:]
			//trimmer = trimmer+indexToTrim
			continue
		}

		if str[cursor] == pattern[cursor] {

			for i := cursor; i >= 0; i-- {
				if str[i] == pattern[i]{
					count++
					if count == l {
						fmt.Printf("found at index: %d, result: %s pattern: %s \n", 0, str, pattern)
						return
					}
				}



			}

		}

	}

}
