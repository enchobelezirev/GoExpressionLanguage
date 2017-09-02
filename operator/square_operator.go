package operator

import "math"

type SquareOperator struct {
}

// GetSupportedSymbols gets the supported plus symbols
func (op SquareOperator) GetSupportedSymbols() map[string]bool {
	supportedSymbols := make(map[string]bool)
	supportedSymbols["sqrt"] = true
	return supportedSymbols
}

// Calculate sums all the operands
func (op SquareOperator) Calculate(operands []int64) int64 {
	return int64(math.Sqrt(float64(operands[0])))
}
