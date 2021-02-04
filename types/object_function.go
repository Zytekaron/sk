package types

import "fmt"

type Function struct {
	Name       string
	Parameters []FunctionParam

}

func NewFunction() *Function {
	return &Function{

	}
}

func (f *Function) Type() Type {
	return "function"
}

func (f *Function) String() string {
	return fmt.Sprintf("[fn %s]", f.Name)
}

func (f *Function) Format(int, int, []Value) string {
	return f.String()
}

func (f *Function) Call(env *Environment, args ...Value) Value {
	// ...
	return null
}

func (f *Function) Interface() interface{} {
	return f.String() // todo idk
}

type FunctionParam struct {
	Name    string
	Default Value
	Spread  bool
	Type    Type
}
