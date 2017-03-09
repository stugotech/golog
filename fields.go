package golog

import "strconv"

// the following field types are built in
const (
	int8Field fieldType = iota
	int16Field
	int32Field
	int64Field
	uint8Field
	uint16Field
	uint32Field
	uint64Field
	float32Field
	float64Field
	stringField
	boolField
)

// fieldType represents the type of a field
type fieldType int

// Field represents a named field of data to be logged
type Field interface {
	Name() string
	String() string
}

// field represents a built-in field type
type field struct {
	name      string
	value     interface{}
	valueType fieldType
}

// Int8 creates a field to represent a signed integer
func Int8(name string, value int8) Field {
	return &field{
		valueType: int8Field,
		name:      name,
		value:     value,
	}
}

// Int16 creates a field to represent a signed integer
func Int16(name string, value int16) Field {
	return &field{
		valueType: int16Field,
		name:      name,
		value:     value,
	}
}

// Int32 creates a field to represent a signed integer
func Int32(name string, value int32) Field {
	return &field{
		valueType: int32Field,
		name:      name,
		value:     value,
	}
}

// Int64 creates a field to represent a signed integer
func Int64(name string, value int64) Field {
	return &field{
		valueType: int64Field,
		name:      name,
		value:     value,
	}
}

// Int creates a field to represent a signed integer
func Int(name string, value int) Field {
	return &field{
		valueType: int64Field,
		name:      name,
		value:     int64(value),
	}
}

// Uint8 creates a field to represent an unsigned integer
func Uint8(name string, value uint8) Field {
	return &field{
		valueType: uint8Field,
		name:      name,
		value:     value,
	}
}

// Uint16 creates a field to represent an unsigned integer
func Uint16(name string, value uint16) Field {
	return &field{
		valueType: uint16Field,
		name:      name,
		value:     value,
	}
}

// Uint32 creates a field to represent an unsigned integer
func Uint32(name string, value uint32) Field {
	return &field{
		valueType: uint32Field,
		name:      name,
		value:     value,
	}
}

// Uint64 creates a field to represent an unsigned integer
func Uint64(name string, value uint64) Field {
	return &field{
		valueType: uint64Field,
		name:      name,
		value:     value,
	}
}

// Uint creates a field to represent an unsigned integer
func Uint(name string, value uint) Field {
	return &field{
		valueType: uint64Field,
		name:      name,
		value:     uint64(value),
	}
}

// Float32 creates a field to represent a real number
func Float32(name string, value float32) Field {
	return &field{
		valueType: float32Field,
		name:      name,
		value:     value,
	}
}

// Float64 creates a field to represent a real number
func Float64(name string, value float64) Field {
	return &field{
		valueType: float64Field,
		name:      name,
		value:     value,
	}
}

// Bool creates a field to represent a boolean value
func Bool(name string, value bool) Field {
	return &field{
		valueType: boolField,
		name:      name,
		value:     value,
	}
}

// Name returns the name of a field
func (f *field) Name() string {
	return f.name
}

// String returns the value of a field formatted as a string
func (f *field) String() string {
	switch f.valueType {
	case int8Field:
		i := f.value.(int8)
		return strconv.FormatInt(int64(i), 10)
	case int16Field:
		i := f.value.(int16)
		return strconv.FormatInt(int64(i), 10)
	case int32Field:
		i := f.value.(int32)
		return strconv.FormatInt(int64(i), 10)
	case int64Field:
		i := f.value.(int64)
		return strconv.FormatInt(i, 10)
	case uint8Field:
		i := f.value.(uint8)
		return strconv.FormatUint(uint64(i), 10)
	case uint16Field:
		i := f.value.(uint16)
		return strconv.FormatUint(uint64(i), 10)
	case uint32Field:
		i := f.value.(uint32)
		return strconv.FormatUint(uint64(i), 10)
	case uint64Field:
		i := f.value.(uint64)
		return strconv.FormatUint(i, 10)
	case float32Field:
		r := f.value.(float32)
		return strconv.FormatFloat(float64(r), 'g', -1, 32)
	case float64Field:
		r := f.value.(float64)
		return strconv.FormatFloat(r, 'g', -1, 64)
	case boolField:
		b := f.value.(bool)
		return strconv.FormatBool(b)
	case stringField:
		return f.value.(string)
	default:
		panic("fields: invalid value type")
	}
}
