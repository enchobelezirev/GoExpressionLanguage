package evaluator

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const inputBaseSystem = "ibase"
const outputBaseSystem = "obase"

type NumberBaseSystemEvaluator struct {
	number string
}

func NewNumberBaseSystemEvaluator(number string) *NumberBaseSystemEvaluator {
	return &NumberBaseSystemEvaluator{number: number}
}

func GetRealNumberInBaseSystemRange(ascii rune, upperBound int) int64 {
	if upperBound <= 10 {
		if ascii >= rune(48) && ascii <= rune(48+upperBound-1) {
			return int64(ascii - 48)
		}
	} else if upperBound > 10 {
		if ascii >= rune(48) && ascii <= rune(57) {
			return int64(ascii - 48)
		}

		if ascii >= rune(65) && ascii <= rune(65+upperBound-11) {
			return int64(ascii - 55)
		}
	}

	return -1
}

func validateDigits(inputBaseSystem int, number string) error {
	for index, digit := range number {
		realNumber := GetRealNumberInBaseSystemRange(digit, inputBaseSystem)
		if realNumber == -1 {
			return fmt.Errorf("Unsupported digit: %c at position: %d.", digit, index)
		}

	}
	return nil
}

func (evaluator NumberBaseSystemEvaluator) Validate(context Context) error {
	inputBaseSystemVariable := context.GetExecutionVariable(inputBaseSystem)
	if inputBaseSystemVariable == "" {
		return errors.New("Undefined input base system. See help for more information.")
	}
	outputBaseSystemVariable := context.GetExecutionVariable(outputBaseSystem)
	if outputBaseSystemVariable == "" {
		return errors.New("Undefined output base system. See help for more information.")
	}

	inputBaseSystem, err := strconv.Atoi(inputBaseSystemVariable)
	if err != nil {
		return fmt.Errorf("Not supported input base system. See help for supported input base systems")
	}

	outputBaseSystem, err := strconv.Atoi(outputBaseSystemVariable)
	if err != nil {
		return fmt.Errorf("Not supported output base system. See help for supported output base systems")
	}

	if inputBaseSystem < 2 || inputBaseSystem > 36 {
		return errors.New("Unknown input base system. See help for supported input base systems.")
	}

	if outputBaseSystem < 2 || outputBaseSystem > 36 {
		return errors.New("Unknown output base system. See help for supported output base systems. ")
	}

	return validateDigits(inputBaseSystem, evaluator.number)
}

func reverse(ss []int) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

func convertToRune(number int) rune {
	if number >= 0 && number <= 9 {
		return rune(number + 48)
	}
	return rune(number + 55)
}

func convertFromTenthBaseSystemToOutputBase(numberInTenthBaseSystem int, outputBaseSystem int) []int {
	var res []int
	for numberInTenthBaseSystem > 0 {
		if numberInTenthBaseSystem < outputBaseSystem {
			res = append(res, numberInTenthBaseSystem)
			break
		}
		remainder := numberInTenthBaseSystem % outputBaseSystem
		res = append(res, remainder)
		numberInTenthBaseSystem = numberInTenthBaseSystem / outputBaseSystem
	}
	return res
}

func convertToTenthBaseSystem(number string, inputBaseSystem, outputBaseSystem int) int64 {
	number = strings.Trim(number, "\r\n")
	length := int64(len(number) - 1)
	var intResult int64
	for i, ascii := range number {
		number1 := GetRealNumberInBaseSystemRange(ascii, inputBaseSystem)
		calculatingIndex := int64(length - int64(i))
		intResult += number1 * int64(math.Pow(float64(inputBaseSystem), float64(calculatingIndex)))
	}
	return intResult
}
func (evaluator NumberBaseSystemEvaluator) Eval(context Context) { // ? should there be any Result returned
	inputBaseSystem, _ := strconv.Atoi(context.GetExecutionVariable(inputBaseSystem))
	outputBaseSystem, _ := strconv.Atoi(context.GetExecutionVariable(outputBaseSystem))
	toTenthBase := convertToTenthBaseSystem(evaluator.number, inputBaseSystem, outputBaseSystem)
	testing := convertFromTenthBaseSystemToOutputBase(int(toTenthBase), outputBaseSystem)
	reverse(testing)
	var runned []string
	for _, number := range testing {
		runned = append(runned, string(convertToRune(number)))
	}
	fmt.Printf("The result is: %v\n", strings.Join(runned, ""))
}
