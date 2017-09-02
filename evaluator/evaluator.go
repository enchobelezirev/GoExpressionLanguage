package evaluator

type Evaluator interface {
	Validate(context Context) error
	Eval(context Context)
}
