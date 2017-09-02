package operator

var operators = []Operator{
	PlusOperator{},
	MinusOperator{},
	MultiplyOperator{},
	DivideOperator{},
	SinusOperator{},
	CosinusOperator{},
	TangensOperator{},
	SquareOperator{},
	PowerOperator{},
	AndOperator{},
	OrOperator{},
}

type OperatorParser struct {
	supportedOperators []Operator
}

func NewDefaultOperatorParser() *OperatorParser {
	return &OperatorParser{supportedOperators: operators}
}

func NewOperatorParser(operators []Operator) *OperatorParser {
	return &OperatorParser{supportedOperators: operators}
}

func (operatorParser OperatorParser) Parse(token string) *Operator {
	for _, supportedOperator := range operatorParser.supportedOperators {
		supportedSymbols := supportedOperator.GetSupportedSymbols()
		if supportedSymbols[token] {
			return &supportedOperator
		}
	}
	return nil
}
