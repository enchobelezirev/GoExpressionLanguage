package evaluator

type Context struct {
	executionVariables map[string]string
}

func NewContext() Context {
	return Context{executionVariables: make(map[string]string)}
}

func (context *Context) AddExecutionVariable(variableName, variableValue string) {
	context.executionVariables[variableName] = variableValue
}

func (context *Context) GetExecutionVariable(variableName string) string {
	return context.executionVariables[variableName]
}
