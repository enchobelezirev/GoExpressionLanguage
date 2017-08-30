package operator

// Operator defines the interface that an operator have to implement in order to be used in an expression
type Operator interface {
	GetSupportedSymbols() map[string]bool
	Calculate(operands []int64) int64
}

type OperatorWithOperands struct {
	op       Operator
	operands []int64
}

func NewOperatorWithOperands(op Operator, operands []int64) *OperatorWithOperands {
	return &OperatorWithOperands{op, operands}
}

func (opWithOperands OperatorWithOperands) CalculateOperatorWithOperands() int64 {
	return opWithOperands.op.Calculate(opWithOperands.operands)
}
func (opWithOperands *OperatorWithOperands) AddOperandToOperator(operand int64) {
	opWithOperands.operands = append(opWithOperands.operands, operand)
}
