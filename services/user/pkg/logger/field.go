package logger

type FieldType uint8

const (
	UnknownType FieldType = iota
	BoolType
	IntType
	Float64Type
	StringType
	ErrorType
)

type Field struct {
	Key   string
	Value interface{}
	Type  FieldType
}
