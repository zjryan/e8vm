package types

// Type is a general interface of a type in G language
type Type interface {
	Size() int32
	String() string
}
