package operator

import "math"

type TangensOperator struct {
}

// GetSupportedSymbols gets the supported plus symbols
func (op TangensOperator) GetSupportedSymbols() map[string]bool {
	supportedSymbols := make(map[string]bool)
	supportedSymbols["tan"] = true
	return supportedSymbols
}

// Calculate sums all the operands
func (op TangensOperator) Calculate(operands []int64) int64 {
	return int64(math.Tan(float64(operands[0])))
}
