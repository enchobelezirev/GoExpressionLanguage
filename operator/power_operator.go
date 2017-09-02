package operator

import "math"

type PowerOperator struct {
}

// GetSupportedSymbols gets the supported plus symbols
func (op PowerOperator) GetSupportedSymbols() map[string]bool {
	supportedSymbols := make(map[string]bool)
	supportedSymbols["^"] = true
	return supportedSymbols
}

// Calculate sums all the operands
func (op PowerOperator) Calculate(operands []int64) int64 {
	return int64(math.Pow(float64(operands[0]), float64(operands[1])))
}
