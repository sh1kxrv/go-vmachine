package cpu

type Object interface {
	Type() string
}

type IntObject struct {
	Value int
}
