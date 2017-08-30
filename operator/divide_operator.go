package operator

type DivideOperator struct {
}

func (op DivideOperator) GetSupportedSymbols() map[string]bool {
	supportedSymbols := make(map[string]bool)
	supportedSymbols["/"] = true
	supportedSymbols["%"] = true
	supportedSymbols["÷"] = true
	return supportedSymbols
}

func (op DivideOperator) Calculate(operands []int64) int64 {
	result := operands[0]
	for _, number := range operands[1:] {
		result /= number
	}
	return result
}
