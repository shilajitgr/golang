/*
sort the string slice with the following conditions:
1. even lenght string precedes odd length strings
2. for odd length strings, shorter length string precedes longer
3. for even length strings, longer length string precedes shorter
4. for equal length strings, the appear in alphabetical order
*/

package main

import (
	"fmt"
	"sort"
)

func customSorting(strArr []string) []string {
	fmt.Println(strArr)
	arrLen := len(strArr)
	oddevenArr := make([]string, arrLen)
	oddcount := 0
	evencount := 1
	lenCount := make([]int, arrLen)
	for _, str := range strArr {
		strLen := len(str)
		if strLen%2 != 0 {
			oddevenArr[oddcount] = str
			lenCount[oddcount] = strLen
			oddcount += 1
		} else {
			oddevenArr[arrLen-evencount] = str
			lenCount[arrLen-evencount] = strLen
			evencount += 1
		}
	}

	fmt.Println(oddevenArr)
	// 1 is done

	for i := 0; i < oddcount; i++ {
		for j := 0; j < oddcount; j++ {
			if lenCount[i] < lenCount[j] {
				temp := lenCount[i]
				lenCount[i] = lenCount[j]
				lenCount[j] = temp
				temp_str := oddevenArr[i]
				oddevenArr[i] = oddevenArr[j]
				oddevenArr[j] = temp_str
			}
			if lenCount[i] == lenCount[j] {
				tmparr := []string{oddevenArr[i], oddevenArr[j]}
				sort.Strings(tmparr)
				oddevenArr[i] = tmparr[1]
				oddevenArr[j] = tmparr[0]
			}
		}
	}

	fmt.Println(oddevenArr)

	for i := oddcount; i < arrLen; i++ {
		for j := oddcount; j < arrLen; j++ {
			if lenCount[i] > lenCount[j] {
				temp := lenCount[i]
				lenCount[i] = lenCount[j]
				lenCount[j] = temp
				temp_str := oddevenArr[i]
				oddevenArr[i] = oddevenArr[j]
				oddevenArr[j] = temp_str
			}
			if lenCount[i] == lenCount[j] {
				tmparr := []string{oddevenArr[i], oddevenArr[j]}
				sort.Strings(tmparr)
				oddevenArr[i] = tmparr[1]
				oddevenArr[j] = tmparr[0]
			}
		}
	}

	fmt.Println(oddevenArr)

	return oddevenArr
}

func main() {
	fmt.Println(customSorting([]string{"a", "ab", "wbc", "abc", "abcd", "abcde"}))
}
