package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var carry byte

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
	exp := ""
	fmt.Print("Digite sua expressão: ")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	exp = strings.TrimSpace(input)
	exp = symbolChange(exp)
	fmt.Println(exp, "\n ")

	//p -> q ^ (~q v p <->  (~p <-> ~q v r) ) <-> (p _v q)

	//the main table
	mainExp := exprSplit(exp)
	_, tableY := getTableSize(mainExp)
	oneZero := ""
	trueFalse := ""

	for i := 0; i < tableY; i++ {
		mainExp := exprSplit(exp)
		// fmt.Println("Line ", i+1)
		boolRest, err := strconv.ParseBool(start(mainExp, i))
		// fmt.Println()
		if err != nil {
			log.Fatal(err)
		}

		switch boolRest {
		case true:
			oneZero += "1"
			trueFalse += "T"
		case false:
			oneZero += "0"
			trueFalse += "F"
		}
	}

	var tableType string
	if oneZero == "1111" {
		tableType = "tautology"
	} else if oneZero == "0000" {
		tableType = "contradiction"
	} else {
		tableType = "contingency"
	}

	fmt.Print(oneZero, "\n", trueFalse, "\n", "type = ", tableType, "\n")
}

func start(mainExp []string, line int) string {
	tableX, _ := getTableSize(mainExp)
	boolExp := changePQR(mainExp, line, tableX)
	// fmt.Println("mainExp = ", mainExp)

	usdExp := boolExp
	// fmt.Println("boolExp = ", usdExp)

	_, qtd := parCount(mainExp)
	for qtd != 0 {
		countsIsCorrect, qtd := parCount(boolExp)

		if countsIsCorrect && qtd > 0 {
			open, close := getHightPar(usdExp)
			usdExp = usdExp[open+1 : close]
			//start resolving the arr
			usdExp = startResolution(usdExp)
			result := usdExp[0]
			boolExp = removeFromToArr(open, close, boolExp)
			boolExp[open] = result
			usdExp = boolExp
			// fmt.Println("usdExp  = ", usdExp)
			//end
		} else if !countsIsCorrect {
			log.Fatal("The number of parentheses is not correct!")
		}

		if len(usdExp) == 1 {
			break
		}
	}

	return usdExp[0]
}

func symbolChange(exprStr string) string {
	newExp := ""
	for _, v := range exprStr {
		switch string(v) {
		case "→":
			newExp += "->"
		case "↔":
			newExp += "<->"
		case "∧":
			newExp += "^"
		case "∼":
			newExp += "~"
		case "∨":
			newExp += "v"
		default:
			newExp += string(v)
		}
	}

	return newExp
}

func startResolution(arr []string) []string {
	operator := true
	for operator {
		arr, operator = hasNegOperator(arr)
	}

	operator = true
	for operator {
		arr, operator = hasOperator("^", arr, conjunction)
	}

	operator = true
	for operator {
		arr, operator = hasOperator("v", arr, inclusiveDisjunction)
	}

	operator = true
	for operator {
		arr, operator = hasOperator("_v", arr, exclusiveDisjunction)
	}

	operator = true
	for operator {
		arr, operator = hasOperator("->", arr, conditional)
	}

	operator = true
	for operator {
		arr, operator = hasOperator("<->", arr, biConditional)
	}

	return arr
}

func hasOperator(operator string, arr []string, resFunc func(bool, bool) bool) ([]string, bool) {
	if pos := findInArr(operator, arr); pos != -1 {
		boolV1, err := strconv.ParseBool(arr[pos-1])
		if err != nil {
			log.Fatalf("fail converting %s of the array %v", arr[pos+1], arr)
		}
		boolV2, err := strconv.ParseBool(arr[pos+1])
		if err != nil {
			log.Fatalf("fail converting %s of the array %v", arr[pos+1], arr)
		}

		arr[pos] = strconv.FormatBool(resFunc(boolV1, boolV2))
		arr = removeArrPos(arr, pos+1, pos-1)
		return arr, true
	}

	return arr, false
}

func hasNegOperator(arr []string) ([]string, bool) {
	if pos := findInArr("~", arr); pos != -1 {
		boolV, err := strconv.ParseBool(arr[pos+1])
		if err != nil {
			log.Fatalf("fail converting %s of the array %v", arr[pos+1], arr)
		}
		arr[pos+1] = strconv.FormatBool(negation(boolV))
		arr = removeArrPos(arr, pos)

		return arr, true
	}

	return arr, false
}

// removeArrPos positions want start with
func removeArrPos(arr []string, positions ...int) []string {
	for _, pos := range positions {
		arr = append(arr[:pos], arr[pos+1:]...)
	}

	return arr
}

func removeFromToArr(first, last int, arr []string) []string {
	for ; last > first; last-- {
		arr = append(arr[:last-1], arr[last:]...)
	}

	return arr
}

func findInArr(value string, arr []string) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}

	return -1
}

func changePQR(arr []string, line, tableSize int) []string {
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

func getTableSize(arr []string) (int, int) {
	var tableX, tableY, p, q, r int
	for _, v := range arr {
		switch v {
		case "p":
			p++
		case "q":
			q++
		case "r":
			r++
		}
	}

	if p != 0 {
		tableX++
	}
	if q != 0 {
		tableX++
	}
	if r != 0 {
		tableX++
	}

	if tableX == 2 {
		tableY = 4
	} else if tableX == 3 {
		tableY = 8
	}

	return tableX, tableY
}

func exprSplit(expression string) []string {
	truthExArr := []string{"("}

	for i := 0; i < len(expression); i++ {
		switch string(expression[i]) {
		case "_":
			carry = '_'
		case "-":
			if carry == '<' {
				carry = '='
			} else {
				carry = '-'
			}
		case "<":
			carry = '<'
		case ">":
			truthExArr = append(truthExArr, carryCheck(carry, expression[i]))
			carry = '0'
		case "v":
			truthExArr = append(truthExArr, carryCheck(carry, expression[i]))
			carry = '0'
		case " ":
			// Just ignore
		default:
			truthExArr = append(truthExArr, string(expression[i]))
		}
	}

	truthExArr = append(truthExArr, ")")

	return truthExArr
}

func parCount(arr []string) (bool, int) {
	open, close := 0, 0
	for _, v := range arr {
		if v == "(" {
			open++
		} else if v == ")" {
			close++
		}
	}

	if open == close {
		return true, close
	}

	return false, -1
}

func getHightPar(exprArr []string) (int, int) {
	open := -1
	for i, v := range exprArr {
		if v == "(" {
			open = i
		} else if v == ")" {
			return open, i
		}
	}

	return open, -1
}

func carryCheck(carry, value byte) (stash string) {
	switch carry {
	case '-':
		stash = "->"
	case '_':
		stash = "_v"
	case '=':
		stash = "<->"
	default:
		stash = string(value)
	}

	return
}

// ~
func negation(p bool) bool {
	return !p
}

// ^
func conjunction(p, q bool) bool {
	return p && q
}

// v
func inclusiveDisjunction(p, q bool) bool {
	return p || q
}

// _v
func exclusiveDisjunction(p, q bool) bool {
	return p != q
}

// Conditional  é igual a "->"
func conditional(p, q bool) bool {
	return !(p == true && q == false)
}

// <->
func biConditional(p, q bool) bool {
	return p == q
}
