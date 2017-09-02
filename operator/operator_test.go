package operator

import "testing"

func TestOperatorParserNotSupportedOperator(t *testing.T) {
	parser := NewDefaultOperatorParser()
	operator := parser.Parse("not-supported-operator")
	if operator != nil {
		t.Error("Expected no operator to be parsed")
	}
}

func TestOperatorParserSupportedOperatorMultiply(t *testing.T) {
	parser := NewDefaultOperatorParser()
	operator := parser.Parse("x")
	_, ok := (*operator).(MultiplyOperator)
	if !ok {
		t.Error("Expected a MultiplyOperator")
	}
}
func TestOperatorParserSupportedOperatorMinus(t *testing.T) {
	parser := NewDefaultOperatorParser()
	operator := parser.Parse("-")
	_, ok := (*operator).(MinusOperator)
	if !ok {
		t.Error("Expected a MinusOperator")
	}
}
func TestOperatorParserSupportedOperatorDivision(t *testing.T) {
	parser := NewDefaultOperatorParser()
	operator := parser.Parse("÷")
	_, ok := (*operator).(DivideOperator)
	if !ok {
		t.Error("Expected a DivideOperator")
	}
}
func TestOperatorParserSupportedOperatorPlus(t *testing.T) {
	parser := NewDefaultOperatorParser()
	operator := parser.Parse("+")
	_, ok := (*operator).(PlusOperator)
	if !ok {
		t.Error("Expected a PlusOperator")
	}
}
func TestOperatorParserSupportedOperatorSqrt(t *testing.T) {
	parser := NewDefaultOperatorParser()
	operator := parser.Parse("sqrt")
	_, ok := (*operator).(SquareOperator)
	if !ok {
		t.Error("Expected a SquareOperator")
	}
}
func TestOperatorParserSupportedOperatorSinus(t *testing.T) {
	parser := NewDefaultOperatorParser()
	operator := parser.Parse("sin")
	_, ok := (*operator).(SinusOperator)
	if !ok {
		t.Error("Expected a MultiplyOperator")
	}
}
func TestOperatorParserSupportedOperatorAnd(t *testing.T) {
	parser := NewDefaultOperatorParser()
	operator := parser.Parse("&&")
	_, ok := (*operator).(AndOperator)
	if !ok {
		t.Error("Expected a AndOperator")
	}
}
