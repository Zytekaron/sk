package types

type Class struct {
	// The Name of this class
	Name string

	// The Methods this class defines
	Methods map[string]*Function
}

func NewClass(name string) *Class {
	return &Class{
		Name: name,
		Methods: make(map[string]*Function),
	}
}

// todo: class members
func (c *Class) New(env *Environment, args ...Value) Value {
	constructor, ok := c.Methods[c.Name]
	obj := NewTypedObject(c)
	if ok {
		constructor.Call(env, )
	}
	return obj
}

func (c *Class) AddMethod(name string, method *Function) {
	_, ok := c.Methods[name]
	if ok {
		NewError("SyntaxError", "Duplicate method found in class " + c.Name + ": " + name).Panic()
		return
	}

	c.Methods[name] = method
}