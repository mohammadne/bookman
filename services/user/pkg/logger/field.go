package logger

type Field struct {
	Key   string
	Value interface{}
}

// String constructs a field with the given key and value.
func String(key string, val string) Field {
	return Field{Key: key, Value: val}
}

// String constructs a field with the given key and value.
func Any(key string, val interface{}) Field {
	return Field{Key: key, Value: val}
}
