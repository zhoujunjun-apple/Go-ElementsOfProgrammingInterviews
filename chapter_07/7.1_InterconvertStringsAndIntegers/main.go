package main

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

func main() {

}
