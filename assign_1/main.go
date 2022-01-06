package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range nums {
		if num%2 == 0 {
			fmt.Printf("%v is even\n", num)
		} else {
			fmt.Printf("%v is odd\n", num)
		}
	}
}

/*
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

 func customSorting(strArr []string) []string {
    fmt.Println(strArr)
    arrLen := len(strArr)
    oddevenArr := make([]string, arrLen)
    oddcount := 0
    lenCount := make([]int, arrLen)
    for i, str := range strArr {
        strLen := len(str)
        if strLen%2 != 0{
            oddevenArr[i] = str
            oddcount +=1
            lenCount[i] = strLen
        }
        fmt.Println(oddevenArr)
    }
    for _, str := range strArr {
        strLen := len(str)
        if strLen%2 == 0{

            oddevenArr[oddcount] = str
            lenCount[oddcount] = strLen
            oddcount+=1
        }
        fmt.Println(oddevenArr)
        fmt.Println(len(oddevenArr))
    }

    // 1 is done

    for i:=0; i<oddcount; i++{
        for j:=0; j<oddcount; j++{
            if lenCount[i]>lenCount[j]{
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
                oddevenArr[i] = tmparr[0]
                oddevenArr[j] = tmparr[1]
            }
        }
    }

    for i:=oddcount; i<arrLen; i++{
        for j:=oddcount; j<arrLen; j++{
            if lenCount[i]<lenCount[j]{
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

    return oddevenArr
}
*/
