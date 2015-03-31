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
	fields  []*structField
	methods []*structMethod
}
