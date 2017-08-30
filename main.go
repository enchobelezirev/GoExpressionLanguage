package main

import (
	"fmt"
	"os"
	"strings"

	operatorPackage "github.com/megaboy2/GoExpressionLanguage/operator"
	"github.com/megaboy2/GoExpressionLanguage/util"
)

func main() {
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
