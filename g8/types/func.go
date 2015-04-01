package types

import (
	"bytes"
	"fmt"

	"lonnie.io/e8vm/fmt8"
	"lonnie.io/e8vm/g8/ir"
)

// Arg is a function argument or return value
type Arg struct {
	Name string // optional
	Type
}

// String returns "name T" for the named argument and "T" for an
// anonymous argument
func (a *Arg) String() string {
	if a.Name == "" {
		return a.Type.String()
	}
	return fmt.Sprintf("%s %s", a.Name, a.Type)
}

// Func is a function pointer type.
// It represents a particular function signature in G language.
type Func struct {
	Args     []*Arg
	Rets     []*Arg
	RetTypes []Type
	Sig      *ir.FuncSig // the signature for IR
}

// NewFuncNamed creates a new function type
func NewFuncNamed(args []*Arg, rets []*Arg) *Func {
	ret := new(Func)
	ret.Args = args
	ret.Rets = rets

	ret.Sig = makeFuncSig(ret)
	ret.RetTypes = argTypes(ret.Rets)
	return ret
}

func argTypes(args []*Arg) []Type {
	if args == nil {
		return nil
	}
	ret := make([]Type, 0, len(args))
	for _, arg := range args {
		ret = append(ret, arg.Type)
	}
	return ret
}

// NewFunc creates a new function type where all its arguments
// and return values are anonymous.
func NewFunc(args []Type, rets []Type) *Func {
	f := new(Func)
	for _, arg := range args {
		f.Args = append(f.Args, &Arg{Type: arg})
	}
	for _, ret := range rets {
		f.Rets = append(f.Rets, &Arg{Type: ret})
	}

	f.Sig = makeFuncSig(f)
	f.RetTypes = rets
	return f
}

// NewVoidFunc creates a new function that does not return anything.
func NewVoidFunc(args ...Type) *Func { return NewFunc(args, nil) }

// String returns the function signature (without the argument names).
func (t *Func) String() string {
	// TODO: this is kind of ugly, need some refactor
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "func (%s) ", fmt8.Join(t.Args, ","))
	if len(t.Rets) > 1 {
		fmt.Fprintf(buf, "(")
		for i, ret := range t.Rets {
			if i > 0 {
				fmt.Fprint(buf, ",")
			}
			fmt.Fprint(buf, ret)
		}
		fmt.Fprint(buf, ")")
	} else if len(t.Rets) == 1 {
		fmt.Fprint(buf, t.Rets[0])
	}

	return buf.String()
}

// Size returns the size of a function pointer, which is always the same
// as the size of a PC register.
func (t *Func) Size() int32 { return 4 }

// converts a langauge function signature into a IR function signature
func makeFuncSig(f *Func) *ir.FuncSig {
	ret := new(ir.FuncSig)
	for _, t := range f.Args {
		ret.AddArg(t.Size(), t.Name)
	}
	for _, t := range f.Rets {
		ret.AddRet(t.Size(), t.Name)
	}
	return ret
}
