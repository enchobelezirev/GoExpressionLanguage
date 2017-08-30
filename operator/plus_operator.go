package operator

const plusSymbol = "+"

// PlusOperator defines the operator '+'
type PlusOperator struct {
}

// GetSupportedSymbols gets the supported plus symbols
func (op PlusOperator) GetSupportedSymbols() map[string]bool {
	supportedSymbols := make(map[string]bool)
	supportedSymbols[plusSymbol] = true
	return supportedSymbols
}

// Calculate sums all the operands
func (op PlusOperator) Calculate(operands []int64) int64 {
	var sum int64
	for _, number := range operands {
		sum += number
	}
	return sum
}
