package evaluator

type NumberBaseSystemParser struct {
}

func (parser NumberBaseSystemParser) Parse(userInput string, context Context) Evaluator {
	for _, digit := range userInput {
		number := GetRealNumberInBaseSystemRange(digit, 36)
		if number == -1 {
			return nil
		}
	}
	return NewNumberBaseSystemEvaluator(userInput)
}
