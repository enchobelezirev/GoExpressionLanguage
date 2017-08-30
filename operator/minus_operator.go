package operator

type MinusOperator struct {
}

func (op MinusOperator) GetSupportedSymbols() map[string]bool {
	supportedSymbols := make(map[string]bool)
	supportedSymbols["-"] = true
	return supportedSymbols
}

func (op MinusOperator) Calculate(operands []int64) int64 {
	result := operands[0]
	for _, number := range operands[1:] {
		result -= number
	}
	return result
}
