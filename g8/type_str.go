package g8

import (
	"bytes"
	"fmt"
)

func typStr(t typ) string {
	switch t := t.(type) {
	case typBasic:
		switch t {
		case typVoid:
			return "void"
		case typInt:
			return "int"
		case typUint:
			return "uint"
		case typInt8:
			return "int8"
		case typUint8:
			return "uint8"
		case typBool:
			return "bool"
		case typString:
			return "string"
		default:
			panic(fmt.Errorf("invalid basic type %d", t))
		}
	case *typPtr:
		return "*" + typStr(t.t)
	case *typSlice:
		return "[]" + typStr(t.t)
	case *typArray:
		return fmt.Sprintf("[%d]%s", t.n, typStr(t.t))
	case *typFunc:
		// TODO: this is kind of ugly, need some refactor
		buf := new(bytes.Buffer)
		fmt.Fprintf(buf, "func (")
		for i, arg := range t.argTypes {
			if i > 0 {
				fmt.Fprintf(buf, ",")
			}
			fmt.Fprintf(buf, typStr(arg))
		}
		fmt.Fprintf(buf, ") ")
		if len(t.retTypes) > 1 {
			fmt.Fprintf(buf, "(")
			for i, ret := range t.retTypes {
				if i > 0 {
					fmt.Fprintf(buf, ",")
				}
				fmt.Fprintf(buf, typStr(ret))
			}
			fmt.Fprintf(buf, ")")
		} else if len(t.retTypes) == 1 {
			fmt.Fprintf(buf, typStr(t.retTypes[0]))
		}

		return buf.String()
	case []typ:
		buf := new(bytes.Buffer)
		for i, a := range t {
			if i > 0 {
				fmt.Fprintf(buf, ",")
			}
			fmt.Fprintf(buf, typStr(a))
		}
		return buf.String()
	default:
		panic(fmt.Errorf("invalid type: %T", t))
	}
}
