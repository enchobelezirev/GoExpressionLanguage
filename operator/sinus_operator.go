package operator

import "math"

type SinusOperator struct {
}

// GetSupportedSymbols gets the supported plus symbols
func (op SinusOperator) GetSupportedSymbols() map[string]bool {
	supportedSymbols := make(map[string]bool)
	supportedSymbols["sin"] = true
	return supportedSymbols
}

// Calculate sums all the operands
func (op SinusOperator) Calculate(operands []int64) int64 {
	return int64(math.Sin(float64(operands[0])))
}
