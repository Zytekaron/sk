package types

import (
	"fmt"
	"strings"
)

type Object struct {
	// The Class that this object is an instance of
	instanceOf *Class

	// A map of data: keys/value pairs
	data map[string]Value
}

func NewObject() *Object {
	return &Object{
		instanceOf: nil,
		data:       make(map[string]Value),
	}
}

func NewTypedObject(instanceOf *Class) *Object {
	return &Object{
		instanceOf: instanceOf,
		data:       make(map[string]Value),
	}
}

func (o *Object) Get(name string) Value {
	if data, ok := o.data[name]; ok {
		return data
	}
	if o.instanceOf != nil {
		return o.instanceOf.Methods[name]
	}
	return null
}

func (o *Object) Set(name string, value Value) Value {
	o.data[name] = value
	return value
}

func (o *Object) Type() Type {
	return "object"
}

func (o *Object) String() string {
	return o.Format(0, 4, make([]Value, 0))

	var data strings.Builder
	for key, value := range o.data {
		data.WriteString(key)
		data.WriteString(": ")
		data.WriteString(value.String())
		data.WriteString(", ")
	}
	res := data.String()
	res = strings.TrimSuffix(res, ", ")
	return fmt.Sprintf("{ %v }", res)
}

// todo: check if visited can have non-cyclic types which can thus be excluded from the map
func (o *Object) Format(depth int, offset int, visited []Value) string {
	if contains(visited, o) {
		if o.instanceOf != nil {
			return fmt.Sprintf("[%s]", o.instanceOf.Name)
		}
		return "[Circular]"
	}
	visited = append(visited, o)

	var data strings.Builder
	for key, value := range o.data {
		data.WriteString(strings.Repeat(" ", depth+offset))
		data.WriteString(key)
		data.WriteString(": ")
		data.WriteString(value.Format(depth+offset, offset, visited))
		data.WriteString(",\n")
	}

	res := data.String()
	res = strings.TrimSuffix(res, ",\n")
	return fmt.Sprintf("{\n%v\n%v}", res, strings.Repeat(" ", depth))
}

func (o *Object) Interface() interface{} {
	return o.data
}

func contains(slice []Value, o Value) bool {
	for _, e := range slice {
		if e == o {
			return true
		}
	}
	return false
}
