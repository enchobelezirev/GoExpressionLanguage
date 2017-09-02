package customCommands

import (
	"fmt"

	evaluator "github.com/megaboy2/GoExpressionLanguage/evaluator"
)

var supportedEvaluators = []evaluator.Evaluator{
	evaluator.MathematicalEvaluator{},
	evaluator.NumberBaseSystemEvaluator{},
}

type Help struct {
	evaluators []evaluator.Evaluator
}

func NewHelp() Help {
	return Help{evaluators: supportedEvaluators}
}

func NewHelpWithCustomEvaluators(customEvaluators []evaluator.Evaluator) Help {
	return Help{evaluators: customEvaluators}
}

func (help Help) PrintHelp() {
	for _, evaluator := range help.evaluators {
		fmt.Printf("%s\n\t%s:  %s\n\n", evaluator.GetUsage().Description, evaluator.GetUsage().ExampleUsage, evaluator.GetUsage().DetailedDescription)
	}
}
