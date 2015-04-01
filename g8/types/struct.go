package types

// Field is a named field in a struct
type Field struct {
	Name string
	Type
}

// Method is a method function in a struct
type Method struct {
	Name string
	*Func
}

// Struct is a structure type
type Struct struct {
	Name    string
	Fields  []*Field
	Methods []*Method
}

// Size returns the overall size of the structure type
func (t *Struct) Size() int32 {
	ret := int32(0)
	for _, field := range t.Fields {
		ret += field.Size()
	}
	return ret
}

// String returns the name of the structure type
func (t *Struct) String() string { return t.Name }
