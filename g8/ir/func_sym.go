package ir

import (
	"fmt"
)

// FuncSym creates a function symbol reference to a linkable function.
// It is used to perform function call operations to functions
// from other packages (functinos not declared in the current package,
// and hence only has a symbol and function signature).
func FuncSym(pkg, sym uint32, sig *FuncSig) Ref {
	return &funcSym{pkg, sym, sig}
}

// a function symbol
type funcSym struct {
	pkg, sym uint32
	sig      *FuncSig
}

func (s *funcSym) String() string {
	return fmt.Sprintf("F[%d.%d]", s.pkg, s.sym)
}

func (s *funcSym) Addressable() bool { return false }
