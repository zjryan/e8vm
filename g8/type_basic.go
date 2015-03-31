package g8

type typ interface{}

type typBasic int

const (
	typVoid typBasic = iota
	typInt
	typUint
	typInt8
	typUint8
	typFloat32
	typBool
	typString
)

func isVoid(a typ) bool { return isBasic(a, typVoid) }

func isBasic(a typ, t typBasic) bool {
	code, ok := a.(typBasic)
	if !ok {
		return false
	}
	return code == t
}

func bothBasic(a, b typ, t typBasic) bool {
	return isBasic(a, t) && isBasic(b, t)
}
