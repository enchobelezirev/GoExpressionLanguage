package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/megaboy2/GoExpressionLanguage/evaluator"
)

var supportedParsers = []evaluator.UserInputParser{
	evaluator.MathematicalParser{},
	evaluator.NumberBaseSystemParser{},
}

func getEvaluator(input string, context evaluator.Context) evaluator.Evaluator {
	for _, parser := range supportedParsers {
		evaluator := parser.Parse(input, context)
		if evaluator != nil {
			return evaluator
		}
	}
	return nil
}

func getNumbers(bound int) string {
	var res []string
	if bound <= 10 {
		for index := 48; index <= 48+bound-1; index++ {
			res = append(res, string(rune(index)))
		}
	} else {
		for index := 48; index <= 57; index++ {
			res = append(res, string(rune(index)))
		}

		for index := 65; index <= 65+bound-11; index++ {
			res = append(res, string(rune(index)))
		}
	}
	return strings.Join(res, " ")
}

func printSupportedInputNumbers(inputBaseSystem string) {
	inputBaseSystemNumber, err := strconv.Atoi(inputBaseSystem)
	if err != nil {
		fmt.Printf("Unknown input base system: %s. Expected number.\n", inputBaseSystem)
	} else {
		fmt.Printf("The following symbols will be acceppted as an input base system: %s\n", getNumbers(inputBaseSystemNumber))
	}
}

func printSupportedOutputNumbers(outputBaseSystem string) {
	inputBaseSystemNumber, err := strconv.Atoi(outputBaseSystem)
	if err != nil {
		fmt.Printf("Unknown output base system: %s. Expected number.\n", outputBaseSystem)
	} else {
		fmt.Printf("The following symbols will be used as an output base system: %s\n", getNumbers(inputBaseSystemNumber))
	}
}

func readCustomInput(input string, context evaluator.Context) {
	splittedUserInput := strings.Fields(input)
	if splittedUserInput[0] == "help" {
		// print help
	} else if len(splittedUserInput) == 2 {
		if splittedUserInput[0] == "ibase" {
			printSupportedInputNumbers(splittedUserInput[1])
		}
		if splittedUserInput[0] == "obase" {
			printSupportedOutputNumbers(splittedUserInput[1])
		}
		context.AddExecutionVariable(splittedUserInput[0], splittedUserInput[1])
	}

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	context := evaluator.NewContext()
	for true {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\r\n")
		if input == "exit" {
			break
		}
		evaluator := getEvaluator(input, context)
		if evaluator == nil {
			readCustomInput(input, context)
			continue
		}
		err := evaluator.Validate(context)
		if err != nil {
			fmt.Println(err)
			continue
		}
		evaluator.Eval(context)
	}

}
