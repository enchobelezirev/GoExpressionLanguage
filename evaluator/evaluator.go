package evaluator

import usage "github.com/megaboy2/GoExpressionLanguage/usage"

type Evaluator interface {
	Validate(context Context) error
	Eval(context Context)
	GetUsage() usage.Usage
}
