package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	operatorPackage "github.com/megaboy2/GoExpressionLanguage/operator"
	"github.com/megaboy2/GoExpressionLanguage/util"
)

func working() {
	expression := os.Args[1:]
	operatorStack := operatorPackage.NewOperatorStack()
	splittedExpression := strings.Fields(expression[0])
	operatorParser := operatorPackage.NewDefaultOperatorParser()
	for _, token := range splittedExpression {
		operator := operatorParser.Parse(token)
		if operator != nil {
			operatorWithOperands := operatorPackage.NewOperatorWithOperands(*operator, []int64{})
			operatorStack.Push(operatorWithOperands)
		} else if util.IsOperand(token) {
			operand, _ := util.ConvertTokenToInteger(token)
			operatorWithOperands := operatorStack.Top()
			operatorWithOperands.AddOperandToOperator(int64(operand))
		} else if token == ")" {
			operatorWithOperands := operatorStack.Pop()
			result := operatorWithOperands.CalculateOperatorWithOperands()
			if operatorStack.IsEmpty() {
				fmt.Printf("The result is %d\n", result)
				break
			}
			nextOperator := operatorStack.Top()
			nextOperator.AddOperandToOperator(result)
		}
	}
}

func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

func getRealNumber(ascii rune) int64 {
	if ascii >= rune(48) && ascii <= rune(57) {
		return int64(ascii - 48)
	}

	return int64(ascii - 55)
}

func shit(test int, o int) []int {
	var res []int
	for test > 0 {
		res = append(res, int(math.Remainder(float64(test), float64(o))))
		fmt.Println(o)
		test = test / o
	}
	return res
}

func convertToTenthBaseSystem(number string, inputBase int64) int64 {
	number = strings.Trim(number, "\r\n")
	for _, t := range number {
		fmt.Println(t)
	}
	length := int64(len(number) - 1)
	intResult := int64(0)
	for i, ascii := range number {
		number1 := getRealNumber(ascii)
		calculatingIndex := int64(length - int64(i))
		intResult += number1 * int64(math.Pow(float64(inputBase), float64(calculatingIndex)))
	}
	return intResult
}

func main() {
	fmt.Printf("%d", int64(math.Remainder(25, 8)))

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter input base: ")
	inputBase, _ := reader.ReadString('\n')
	inputBase = strings.Trim(inputBase, "\r\n")
	fmt.Println(inputBase)

	fmt.Print("Enter output base: ")
	outputBase, _ := reader.ReadString('\n')
	outputBase = strings.Trim(outputBase, "\r\n")
	fmt.Println(outputBase)

	fmt.Print("Enter number to be calculated: ")
	number, _ := reader.ReadString('\n')
	fmt.Println(number)
	inputBaseNumber, _ := strconv.Atoi(inputBase)
	fmt.Println(inputBaseNumber)

	toTenthBase := convertToTenthBaseSystem(number, int64(inputBaseNumber))

	fmt.Printf("The number in tenth base %d", toTenthBase)

	outputBaseNumber, _ := strconv.Atoi(outputBase)
	testing := shit(int(toTenthBase), outputBaseNumber)
	fmt.Println(testing)
	// operandStack := NewStack()
	// pendingOperand := false
	// fmt.Println(expression[0][0])
	// splittedExpression := strings.Fields(expression[0])
	// fmt.Printf("splitted %v", splittedExpression)
	// for _, b := range splittedExpression {
	// 	token := b
	// 	fmt.Printf("processing token %s \n", token)
	// 	if isOperator(token) {
	// 		fmt.Println("it is operator")
	// 		operatorStack.Push(token)
	// 		pendingOperand = false
	// 	} else if isOperand(token) {
	// 		fmt.Println("it is operand")
	// 		operand, _ := convertTokenToInteger(token)
	// 		if pendingOperand {
	// 			fmt.Println("there is pendingOperand")
	// 			for !operandStack.IsEmpty() {
	// 				fmt.Println("Poping from operand stack")
	// 				operand1 := operandStack.Pop()
	// 				fmt.Println("Poping from operator stack")
	// 				operator := operatorStack.Pop()
	// 				operand1String := (operand1).(int)
	// 				operatorString := (operator).(string)
	// 				fmt.Printf("processing %s, %d, %d\n", operatorString, operand1String, operand)
	// 				operand = evaluate(operatorString, operand1String, operand)
	// 			}
	// 		}
	// 		fmt.Printf("pushing operand to stack %d \n", operand)
	// 		operandStack.Push(operand)
	// 		pendingOperand = true
	// 	}
	// }
	// fmt.Println(operandStack.Pop().(int))
}
