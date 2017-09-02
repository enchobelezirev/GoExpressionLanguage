package operator

type AndOperator struct {
}

// GetSupportedSymbols gets the supported plus symbols
func (op AndOperator) GetSupportedSymbols() map[string]bool {
	supportedSymbols := make(map[string]bool)
	supportedSymbols["&&"] = true
	supportedSymbols["and"] = true
	return supportedSymbols
}

// Calculate sums all the operands
func (op AndOperator) Calculate(operands []int64) int64 {
	var res bool = true
	for _, operand := range operands {
		res = res && (operand == 1)
	}

	if res {
		return 1
	}
	return 0
}
