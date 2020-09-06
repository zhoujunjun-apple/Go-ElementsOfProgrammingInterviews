package main

import (
	"fmt"
	"strings"
)

//NativeStrToInt function convert string to integer
func NativeStrToInt(intStr string) int {
	intNeg := false
	intVal := 0

	if intStr[0] == '-' {
		intNeg = true
		intStr = intStr[1:]
	}

	for _, b := range intStr {
		switch b {
		case '0':
			intVal = 10 * intVal
		case '1':
			intVal = 10*intVal + 1
		case '2':
			intVal = 10*intVal + 2
		case '3':
			intVal = 10*intVal + 3
		case '4':
			intVal = 10*intVal + 4
		case '5':
			intVal = 10*intVal + 5
		case '6':
			intVal = 10*intVal + 6
		case '7':
			intVal = 10*intVal + 7
		case '8':
			intVal = 10*intVal + 8
		case '9':
			intVal = 10*intVal + 9
		default:
			panic("string contains invalid integer")
		}
	}
	if intNeg {
		intVal = -intVal
	}
	return intVal
}

//SubstractStrToInt function use '0' as the base to convert each digit rune to integer
func SubstractStrToInt(intStr string) int {
	intVal := 0
	intNeg := false

	if intStr[0] == '-' {
		intNeg = true
		intStr = intStr[1:]
	}

	for _, b := range intStr {
		intVal = intVal*10 + int(b-'0') //use '0' as substract base
	}
	if intNeg {
		intVal = -intVal
	}
	return intVal
}

var DigitToString = map[int]string{
	0: "0",
	1: "1",
	2: "2",
	3: "3",
	4: "4",
	5: "5",
	6: "6",
	7: "7",
	8: "8",
	9: "9",
}

//IntToStr function convert integer into string
func IntToStr(it int) string {
	rs := []string{}

	itNeg := false
	if it < 0 {
		itNeg = true
		it = -it
	}

	remain := it
	for it >= 10 {
		remain = it % 10
		if rv, ok := DigitToString[remain]; ok {
			rs = append(rs, rv)
		}
		it = it / 10
		fmt.Printf("it=%d, remain=%d, rs=%v\n", it, remain, rs)
	}
	if rv, ok := DigitToString[it]; ok {
		rs = append(rs, rv)
	}
	if itNeg {
		rs = append(rs, "-")
	}

	for left, right := 0, len(rs)-1; left < right; {
		rs[left], rs[right] = rs[right], rs[left]
		left++
		right--
	}

	return strings.Join(rs, "")
}

func main() {
	fmt.Println(IntToStr(-123))
}
