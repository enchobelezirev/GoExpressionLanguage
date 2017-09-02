package operator

import "math"

type CosinusOperator struct {
}

// GetSupportedSymbols gets the supported plus symbols
func (op CosinusOperator) GetSupportedSymbols() map[string]bool {
	supportedSymbols := make(map[string]bool)
	supportedSymbols["cos"] = true
	return supportedSymbols
}

// Calculate sums all the operands
func (op CosinusOperator) Calculate(operands []int64) int64 {
	return int64(math.Cos(float64(operands[0])))
}
