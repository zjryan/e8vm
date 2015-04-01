package g8

type structField struct {
	name string
	t    typ
}

type structMethod struct {
	name string
	t    *typFunc
}

type typStruct struct {
	name    string
	fields  []*structField
	methods []*structMethod
}

func (t *typStruct) Size() int32 {
	ret := int32(0)
	for _, field := range t.fields {
		ret += field.t.Size()
	}
	return ret
}

func (t *typStruct) String() string { return t.name }
