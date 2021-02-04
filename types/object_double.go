package types

import (
	"fmt"
)

type Double struct {
	*Object
	Value float64
}

func NewDouble(value float64) *Double {
	return &Double{
		Object: NewObject(),
		Value: value,
	}
}

func (d *Double) Type() Type {
	return "double"
}

func (d *Double) String() string {
	return fmt.Sprintf("%f", d.Value)
}

func (d *Double) Format(int, int, []Value) string {
	return d.String()
}

func (d *Double) Interface() interface{} {
	return d.String() // todo idk
}