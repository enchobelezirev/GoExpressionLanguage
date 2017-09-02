package evaluator

import "testing"

func TestMathematicalParserShouldReturnMathematicalEvaluator(t *testing.T) {
	parser := MathematicalParser{}
	evaluator := parser.Parse("( + ( 2 3 ) )", Context{})
	_, ok := evaluator.(*MathematicalEvaluator)
	if !ok {
		t.Error("Expected to get a Mathematical evaluator")
	}
}

func TestMathematicalParserShouldReturnMathematicalEvaluator1(t *testing.T) {
	parser := MathematicalParser{}
	evaluator := parser.Parse(" + 2 3", Context{})
	_, ok := evaluator.(*MathematicalEvaluator)
	if !ok {
		t.Error("Expected to get a Mathematical evaluator")
	}
}

func TestNumberBaseSystemParser(t *testing.T) {
	parser := NumberBaseSystemParser{}
	evaluator := parser.Parse("A34", Context{})
	_, ok := evaluator.(*NumberBaseSystemEvaluator)
	if !ok {
		t.Error("Expected to get a NumberBaseSystemParser evaluator")
	}
}

func TestNumberBaseSystemParser1(t *testing.T) {
	parser := NumberBaseSystemParser{}
	evaluator := parser.Parse("( )", Context{})
	_, ok := evaluator.(*NumberBaseSystemEvaluator)
	if ok {
		t.Error("Not expected Number Base System evaluator")
	}
}

func TestNumberBaseSystemEvaluator(t *testing.T) {
	context := NewContext()
	evaluator := NewNumberBaseSystemEvaluator("A")
	err := evaluator.Validate(context)
	if err == nil {
		t.Error("Expected an error from validation")
	}
	expectedErrorMessage := "Undefined input base system. See help for more information."
	if expectedErrorMessage != err.Error() {
		t.Errorf("Expected %s, got %s\n", expectedErrorMessage, err.Error())
	}
}
func TestNumberBaseSystemEvaluator1(t *testing.T) {
	context := NewContext()
	context.AddExecutionVariable("ibase", "10")
	evaluator := NewNumberBaseSystemEvaluator("A")
	err := evaluator.Validate(context)
	if err == nil {
		t.Error("Expected an error from validation")
	}
	expectedErrorMessage := "Undefined output base system. See help for more information."
	if expectedErrorMessage != err.Error() {
		t.Errorf("Expected %s, got %s\n", expectedErrorMessage, err.Error())
	}
}

func TestNumberBaseSystemEvaluatorNotSupportIbaseSystem(t *testing.T) {
	context := NewContext()
	context.AddExecutionVariable("ibase", "test")
	context.AddExecutionVariable("obase", "test")
	evaluator := NewNumberBaseSystemEvaluator("A")
	err := evaluator.Validate(context)
	if err == nil {
		t.Error("Expected an error from validation")
	}
	expectedErrorMessage := "Not supported input base system. See help for supported input base systems"
	if expectedErrorMessage != err.Error() {
		t.Errorf("Expected %s, got %s\n", expectedErrorMessage, err.Error())
	}
}

func TestNumberBaseSystemEvaluatorNotSupportObaseSystem(t *testing.T) {
	context := NewContext()
	context.AddExecutionVariable("ibase", "10")
	context.AddExecutionVariable("obase", "test")
	evaluator := NewNumberBaseSystemEvaluator("A")
	err := evaluator.Validate(context)
	if err == nil {
		t.Error("Expected an error from validation")
	}
	expectedErrorMessage := "Not supported output base system. See help for supported output base systems"
	if expectedErrorMessage != err.Error() {
		t.Errorf("Expected %s, got %s\n", expectedErrorMessage, err.Error())
	}
}

func TestNumberBaseSystemEvaluatorNotSupportIbaseSystem1(t *testing.T) {
	context := NewContext()
	context.AddExecutionVariable("ibase", "380")
	context.AddExecutionVariable("obase", "10")
	evaluator := NewNumberBaseSystemEvaluator("A")
	err := evaluator.Validate(context)
	if err == nil {
		t.Error("Expected an error from validation")
	}
	expectedErrorMessage := "Unknown input base system. See help for supported input base systems."
	if expectedErrorMessage != err.Error() {
		t.Errorf("Expected %s, got %s\n", expectedErrorMessage, err.Error())
	}
}

func TestNumberBaseSystemEvaluatorNotValidDigits(t *testing.T) {
	context := NewContext()
	context.AddExecutionVariable("ibase", "16")
	context.AddExecutionVariable("obase", "10")
	evaluator := NewNumberBaseSystemEvaluator("KA")
	err := evaluator.Validate(context)
	if err == nil {
		t.Error("Expected an error from validation")
	}
	expectedErrorMessage := "Unsupported digit: K at position: 0."
	if expectedErrorMessage != err.Error() {
		t.Errorf("Expected %s, got %s\n", expectedErrorMessage, err.Error())
	}
}

func TestNumberBaseSystemEvaluatorSuccessfulValidation(t *testing.T) {
	context := NewContext()
	context.AddExecutionVariable("ibase", "16")
	context.AddExecutionVariable("obase", "10")
	evaluator := NewNumberBaseSystemEvaluator("AB")
	err := evaluator.Validate(context)
	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}
}

func TestMathematicaEvaluatorInvalidBrackets(t *testing.T) {
	evaluator := NewMathematicalEvaluator([]string{"(", "(", "(", ")"})
	err := evaluator.Validate(Context{})
	if err == nil {
		t.Error("Expected an error from validation")
	}
	expectedErrorMessage := "The brackets are not matching. Check the closing and the opening brackets. See help for more details"
	if expectedErrorMessage != err.Error() {
		t.Errorf("Expected %s, got %s\n", expectedErrorMessage, err.Error())
	}
}

func TestMathematicaEvaluatorValidBrackets(t *testing.T) {
	evaluator := NewMathematicalEvaluator([]string{"(", "(", "(", ")", ")", ")"})
	err := evaluator.Validate(Context{})
	if err != nil {
		t.Error("Expected an error from validation")
	}

}

func TestMathematicaEvaluatorValidymbols(t *testing.T) {
	evaluator := NewMathematicalEvaluator([]string{"(", "(", "(", "pff", ")", ")"})
	err := evaluator.Validate(Context{})

	if err == nil {
		t.Error("Expected an error from validation")
	}
	expectedErrorMessage := "The input contains invalid symbols. See help for more info."
	if expectedErrorMessage != err.Error() {
		t.Errorf("Expected %s, got %s\n", expectedErrorMessage, err.Error())
	}
}

func TestMathematicaEvaluatorValidymbols1(t *testing.T) {
	evaluator := NewMathematicalEvaluator([]string{"(", "*", "(", "+", ")", ")"})
	err := evaluator.Validate(Context{})

	if err != nil {
		t.Error("Expected an error from validation")
	}

}
