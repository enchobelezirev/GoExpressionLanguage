package evaluator

type UserInputParser interface {
	Parse(userInput string, context Context) Evaluator
}
