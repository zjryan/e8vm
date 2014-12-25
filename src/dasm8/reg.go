package dasm8

func regStr(r uint32) string {
	switch r {
	case 0:
		return "r0"
	case 1:
		return "r1"
	case 2:
		return "r2"
	case 3:
		return "r3"
	case 4:
		return "r4"
	case 5:
		return "sp"
	case 6:
		return "ret"
	case 7:
		return "pc"
	default:
		panic("bug")
	}
}
