package evaluator

import (
	"errors"
	"fmt"
	"strconv"

	util "github.com/megaboy2/GoExpressionLanguage/util"

	operatorPackage "github.com/megaboy2/GoExpressionLanguage/operator"
)

var SupportedMathematicalSymbols = map[string]bool{
	"(":   true,
	"*":   true,
	")":   true,
	"-":   true,
	"/":   true,
	"%":   true,
	"x":   true,
	"sin": true,
	"cos": true,
	"+":   true,
}

type MathematicalEvaluator struct {
	input []string
}

func NewMathematicalEvaluator(input []string) *MathematicalEvaluator {
	return &MathematicalEvaluator{input}
}

func isNumber(stringNumber string) bool {
	_, err := strconv.Atoi(stringNumber)
	if err != nil {
		return false
	}
	return true
}

func hasValidSymbols(input []string) bool {
	for _, symbol := range input {
		if !isNumber(symbol) && !SupportedMathematicalSymbols[symbol] {
			return false
		}
	}
	return true
}

func hasValidBrackets(input []string) bool {
	var bracketsCount int
	for _, symbol := range input {
		if symbol == "(" {
			bracketsCount++
		} else if symbol == ")" {
			bracketsCount--
		}
	}
	return bracketsCount == 0
}

func (evaluator MathematicalEvaluator) Validate(context Context) error {
	if !hasValidSymbols(evaluator.input) {
		return errors.New("The input contains invalid symbols. See help for more info.")
	}

	if !hasValidBrackets(evaluator.input) {
		return errors.New("The brackets are not matching. Check the closing and the opening brackets. See help for more details")

	}

	return nil
}
func (evaluator MathematicalEvaluator) Eval(context Context) { // ? should there be any Result returned
	operatorStack := operatorPackage.NewOperatorStack()
	operatorParser := operatorPackage.NewDefaultOperatorParser()
	for _, token := range evaluator.input {
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
