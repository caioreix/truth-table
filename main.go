package main

import (
	"fmt"
	"strconv"
)

var tableSize = 2

var table4x2 = [][]bool{
	{true, true},
	{true, false},
	{false, true},
	{false, false},
}

var table8x3 = [][]bool{
	{true, true, true},
	{true, true, false},
	{true, false, true},
	{true, false, false},
	{false, true, true},
	{false, true, false},
	{false, false, true},
	{false, false, false},
}

func main() {
	// expr := []string{
	// 	"p", "->", "q", "(", "~", "q", "<->", "p", ")", "_v", "(", "q", "->", "p", ")"
	// }

	parTest := "p -> q v r (q <-> r)"
	parTestArr := []string{}

	for i := 0; i < len(parTest); i++ {
		parTestArr = append(parTestArr, string(parTest[i]))
	}

	tableSizeCalc(parTestArr)

	parTestArr = changePQR(parTestArr, 1)
	fmt.Println(parTestArr)

	if parCount(parTestArr) {
		open, close := getHigthPar(parTestArr)
		slicedP := parTestArr[open+1 : close]
		fmt.Println(slicedP)
	} else {
		fmt.Println("fail")
	}

}

func tableSizeCalc(arr []string) {
	for _, v := range arr {
		if v == "r" {
			tableSize = 3
		}
	}
}

func changePQR(arr []string, line int) []string {
	table := [][]bool{}
	if tableSize == 2 {
		table = table4x2
	} else if tableSize == 3 {
		table = table8x3
	}

	for i, v := range arr {
		switch v {
		case "p":
			arr[i] = strconv.FormatBool(table[line][0])
		case "q":
			arr[i] = strconv.FormatBool(table[line][1])
		case "r":
			arr[i] = strconv.FormatBool(table[line][2])
		}
	}

	return arr
}

func parCount(arr []string) bool {
	open, close := 0, 0
	for _, v := range arr {
		if v == "(" {
			open++
		} else if v == ")" {
			close++
		}
	}

	if open == close {
		return true
	}

	return false
}

func getHigthPar(arr []string) (int, int) {
	open := -1
	for i, v := range arr {
		if v == "(" {
			open = i
		} else if v == ")" {
			return open, i
		}
	}

	return open, -1
}

func logic(expression []string) {}
