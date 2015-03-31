package g8

type typ interface{}

type typBasic int

const (
	typErr typBasic = iota
	typInt
	typUint
	typInt8
	typUint8
	typFloat32
	typBool
	typString
)

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
