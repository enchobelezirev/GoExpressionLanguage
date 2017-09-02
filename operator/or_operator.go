package operator

type OrOperator struct {
}

// GetSupportedSymbols gets the supported plus symbols
func (op OrOperator) GetSupportedSymbols() map[string]bool {
	supportedSymbols := make(map[string]bool)
	supportedSymbols["||"] = true
	supportedSymbols["or"] = true
	return supportedSymbols
}

// Calculate sums all the operands
func (op OrOperator) Calculate(operands []int64) int64 {
	var res bool = false
	for _, operand := range operands {
		res = res || (operand == 1)
	}

	if res {
		return 1
	}
	return 0
}
