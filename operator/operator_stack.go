package operator

type OperatorStack struct {
	slice []*OperatorWithOperands
}

// NewStack ..
func NewOperatorStack() *OperatorStack {
	return &OperatorStack{}
}

func (stack *OperatorStack) Push(v *OperatorWithOperands) {
	stack.slice = append(stack.slice, v)
}

func (stack *OperatorStack) Pop() *OperatorWithOperands {
	length := len(stack.slice)
	if length == 0 {
		return nil
	}

	res := stack.slice[length-1]
	stack.slice = stack.slice[:length-1]
	return res
}

func (stack *OperatorStack) IsEmpty() bool {
	return len(stack.slice) == 0
}

func (stack *OperatorStack) Top() *OperatorWithOperands {
	length := len(stack.slice)
	return stack.slice[length-1]
}
