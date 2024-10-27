package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("PVC - Stellenwert converter")
	fmt.Println("---------------------")
	var inputCode = getInput("Stellenwertcode")
	var inputSystem = getInput("Input system")
	var outputSystem = getInput("Output system")
	i, _ := strconv.Atoi(inputSystem)
	o, _ := strconv.Atoi(outputSystem)
	inputCode = strings.Replace(inputCode, ",", ".", -1)

	if o == 10 {
		fmt.Println(convertToDecimal(inputCode, i, false))
	} else if i == 10 {
		dec, _ := strconv.ParseFloat(inputCode, 64)
		fmt.Println(convertDecimalToCode(dec, o, false))
	} else {
		dec := convertToDecimal(inputCode, i, true)
		fmt.Println(convertDecimalToCode(dec, o, false))
	}
}

func getInput(title string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(title + " -> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		return text
	}
}

func convertToDecimal(inputCode string, inputSystem int, isReversed bool) float64 {
	var splitInput = strings.Split(inputCode, ".")

	var inputCodes = convertStringToNumber(splitInput[0], inputSystem)
	var result float64
	for i := 0; i < len(inputCodes); i++ {
		if isReversed {
			result += float64(inputCodes[i]) * (float64(inputCodes[i]) / math.Pow(float64(inputSystem), float64(i+1)))
		} else {
			result += float64(inputCodes[i]) * math.Pow(float64(inputSystem), float64(i))
		}

	}
	if len(splitInput) > 1 {
		result += convertToDecimal(splitInput[1], inputSystem, true)
	}

	return result
}

func convertDecimalToCode(decimal float64, code int, isReverse bool) string {
	input := int(decimal)
	remainder := decimal - float64(input)

	if isReverse {
		input = 1
	}
	var result string
	round := 0
	for input != 0 {
		if len(result) > 10 {
			break
		}

		if isReverse {
			oldDecimal := decimal * float64(code)
			s := strings.Split(fmt.Sprintf("%.5f", oldDecimal), ".")
			result += s[0]
			newValue, _ := strconv.ParseFloat("0."+string(s[1]), 64)
			decimal = newValue
			if decimal == 0 {
				return result
			}
		} else {
			res := input % code
			result += strconv.Itoa(res)
			input = input / code
		}
		round++
	}
	result = reverseInputString(result)
	if remainder > 0 {
		result += "." + convertDecimalToCode(remainder, code, true)
	}
	return result
}

func convertStringToNumber(inputCode string, inputSystem int) []int {
	var result []int

	for _, el := range inputCode {
		if inputSystem > 10 && inputSystem <= 16 {
			switch string(el) {
			case "A":
				result = append(result, 10)
			case "B":
				result = append(result, 11)
			case "C":
				result = append(result, 12)
			case "D":
				result = append(result, 13)
			case "E":
				result = append(result, 14)
			case "F":
				result = append(result, 15)
			case ",":
				return reverseInput(result)
			default:
				i, _ := strconv.Atoi(string(el))
				result = append(result, i)
			}
		} else {
			i, _ := strconv.Atoi(string(el))
			result = append(result, i)
		}
	}

	return reverseInput(result)
}

func reverseInput(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInput(input[1:]), input[0])
}

func reverseInputString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; j, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
