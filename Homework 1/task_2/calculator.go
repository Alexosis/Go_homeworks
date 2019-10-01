package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const errWithOpening = "opening bracket is missing"
const errWithClosing = "closing bracket is missing"

func main() {
	ans, err := calc(os.Args[1])
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(ans)
	}
}

func calc(arguments string) (string, error) {
	In := strings.Split(arguments, "")
	Temple, err := GetExpression(In)
	if err != nil {
		return "", err
	} else {
		return calculating(Temple), nil
	}
}

func priority(operator string) int {
	operatorsMap := map[string]int{
		"(": 0,
		")": 1,
		"+": 2,
		"-": 3,
		"*": 4,
		"/": 5,
	}

	return operatorsMap[operator]
}

func isOperator(symbol string) bool {
	if symbol == "+" || symbol == "-" || symbol == "*" ||
		symbol == "/" || symbol == "(" || symbol == ")" {
		return true
	}
	return false
}

func GetExpression(input []string) (string, error) {
	output := ""
	operatorsStack := make([]string, 0)
	for i := 0; i < len(input); i++ {
		if _, err := strconv.Atoi(input[i]); err == nil {
			for ; i < len(input) && !isOperator(input[i]); i++ {
				output += input[i]
			}
			output += " "
			i--
		}

		if isOperator(input[i]) {
			if input[i] == "(" {
				operatorsStack = append(operatorsStack, input[i])
			} else if input[i] == ")" {
				if len(operatorsStack) == 1 && operatorsStack[len(operatorsStack)-1] != "(" {
					return "", errors.New(errWithOpening)
				}
				s := operatorsStack[len(operatorsStack)-1]
				operatorsStack = operatorsStack[:len(operatorsStack)-1]
				for s != "(" {
					output += s + " "
					s = operatorsStack[len(operatorsStack)-1]
					operatorsStack = operatorsStack[:len(operatorsStack)-1]
				}
			} else {
				if len(operatorsStack) > 0 {
					if priority(input[i]) <= priority(operatorsStack[len(operatorsStack)-1]) {
						output += operatorsStack[len(operatorsStack)-1] + " "
						operatorsStack = operatorsStack[:len(operatorsStack)-1]
					}
				}
				operatorsStack = append(operatorsStack, input[i])
			}
		}
	}

	for i := len(operatorsStack); i > 0; i-- {
		if operatorsStack[i-1] == "(" {
			return "", errors.New(errWithClosing)
		}
		output += operatorsStack[i-1] + " "
	}
	return output, nil
}

func calculating(input string) string {
	inputLine := strings.Fields(input)
	result := 0
	templeStack := make([]string, 0)
	for i := 0; i < len(inputLine); i++ {
		if _, err := strconv.Atoi(inputLine[i]); err == nil {
			a := inputLine[i]
			templeStack = append(templeStack, a)
		} else if isOperator(inputLine[i]) {
			a, _ := strconv.Atoi(templeStack[len(templeStack)-1])
			templeStack = templeStack[:len(templeStack)-1]
			b, _ := strconv.Atoi(templeStack[len(templeStack)-1])
			templeStack = templeStack[:len(templeStack)-1]

			switch inputLine[i] {
			case "+":
				{
					result = b + a
				}
			case "-":
				{
					result = b - a
				}
			case "*":
				{
					result = b * a
				}
			case "/":
				{
					result = b / a
				}
			}
			templeStack = append(templeStack, strconv.Itoa(result))
		}
	}
	return templeStack[len(templeStack)-1]
}
