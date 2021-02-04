package types

import "fmt"

type Float struct {
	*Object
	Value float32
}

func NewFloat(value float32) *Float {
	return &Float{
		Object: NewObject(),
		Value: value,
	}
}

func (f *Float) Type() Type {
	return "float"
}

func (f *Float) String() string {
	return fmt.Sprintf("%f", f.Value)
}

func (f *Float) Format(int, int, []Value) string {
	return f.String()
}

func (f *Float) Interface() interface{} {
	return f.String() // todo idk
}
