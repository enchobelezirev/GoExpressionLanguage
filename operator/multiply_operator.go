package operator

type MultiplyOperator struct {
}

func (op MultiplyOperator) GetSupportedSymbols() map[string]bool {
	supportedSymbols := make(map[string]bool)
	supportedSymbols["x"] = true
	supportedSymbols["*"] = true
	return supportedSymbols
}

func (op MultiplyOperator) Calculate(operands []int64) int64 {
	var result int64 = 1
	for _, number := range operands {
		result *= number
	}
	return result
}
