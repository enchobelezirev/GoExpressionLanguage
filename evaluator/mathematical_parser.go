package evaluator

import (
	"strings"
)

type MathematicalParser struct {
}

func inputContainsMathematicalSymbols(input []string) bool {
	for _, symbol := range input {
		if SupportedMathematicalSymbols[symbol] {
			return true
		}
	}
	return false
}

func (parser MathematicalParser) Parse(userInput string, context Context) Evaluator {
	splittedExpression := strings.Fields(userInput)
	isMathematical := inputContainsMathematicalSymbols(splittedExpression)
	if isMathematical {
		return NewMathematicalEvaluator(splittedExpression)
	}

	return nil
}
