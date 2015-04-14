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
	T
}

// String returns "name T" for the named argument and "T" for an
// anonymous argument
func (a *Arg) String() string {
	if a.Name == "" {
		return a.T.String()
	}
	return fmt.Sprintf("%s %s", a.Name, a.T)
}

// Func is a function pointer type.
// It represents a particular function signature in G language.
type Func struct {
	Args     []*Arg
	Rets     []*Arg
	RetTypes []T
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

func argTypes(args []*Arg) []T {
	if args == nil {
		return nil
	}
	ret := make([]T, 0, len(args))
	for _, arg := range args {
		ret = append(ret, arg.T)
	}
	return ret
}

// NewFunc creates a new function type where all its arguments
// and return values are anonymous.
func NewFunc(args []T, rets []T) *Func {
	f := new(Func)
	for _, arg := range args {
		f.Args = append(f.Args, &Arg{T: arg})
	}
	for _, ret := range rets {
		f.Rets = append(f.Rets, &Arg{T: ret})
	}

	f.Sig = makeFuncSig(f)
	f.RetTypes = rets
	return f
}

// NewVoidFunc creates a new function that does not return anything.
func NewVoidFunc(args ...T) *Func { return NewFunc(args, nil) }

// MainFuncSig is the signature for "func main()"
var MainFuncSig = NewFunc(nil, nil)

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

func makeArg(t *Arg) *ir.FuncArg {
	return &ir.FuncArg{t.Name, t.Size(), IsBasic(t.T, Uint8)}
}

// converts a langauge function signature into a IR function signature
func makeFuncSig(f *Func) *ir.FuncSig {
	args := make([]*ir.FuncArg, len(f.Args))
	for i, t := range f.Args {
		if t.T == nil {
			panic("type missing")
		}
		args[i] = makeArg(t)
	}

	rets := make([]*ir.FuncArg, len(f.Rets))
	for i, t := range f.Rets {
		rets[i] = makeArg(t)
	}

	return ir.NewFuncSig(args, rets)
}
