package fields

import "github.com/stugotech/golog/encoder"

// FieldType describes the type in a field
type FieldType int

// The following constants define the available field types
const (
	UnknownType FieldType = iota
	BoolType

	Complex128Type
	Complex64Type
	Float64Type
	Float32Type

	Int64Type
	Int32Type
	Int16Type
	Int8Type

	Uint64Type
	Uint32Type
	Uint16Type
	Uint8Type

	StringType
	ArrayType
	ObjectType
)

// Field represents a field in a log or error value
type Field struct {
	Name   string
	Encode encoder.FieldEncoderFunc
}

// Postpone adds a field later if it is called
func Postpone(name string, f encoder.FieldEncoderFunc) *Field {
	return &Field{
		Name:   name,
		Encode: f,
	}
}

// Bool creates a boolean field
func Bool(name string, value bool) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Bool(name, value)
		},
	}
}

// Complex128 creates a complex128 field
func Complex128(name string, value complex128) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Complex128(name, value)
		},
	}
}

// Complex64 creates a complex64 field
func Complex64(name string, value complex64) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Complex64(name, value)
		},
	}
}

// Float64 creates a float64 field
func Float64(name string, value float64) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Float64(name, value)
		},
	}
}

// Float32 creates a float32 field
func Float32(name string, value float32) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Float32(name, value)
		},
	}
}

// Int64 creates a int64 field
func Int64(name string, value int64) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Int64(name, value)
		},
	}
}

// Int32 creates a int32 field
func Int32(name string, value int32) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Int32(name, value)
		},
	}
}

// Int16 creates a int16 field
func Int16(name string, value int16) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Int16(name, value)
		},
	}
}

// Int8 creates a int8 field
func Int8(name string, value int8) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Int8(name, value)
		},
	}
}

// Uint64 creates a uint64 field
func Uint64(name string, value uint64) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Uint64(name, value)
		},
	}
}

// Uint32 creates a uint32 field
func Uint32(name string, value uint32) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Uint32(name, value)
		},
	}
}

// Uint16 creates a uint16 field
func Uint16(name string, value uint16) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Uint16(name, value)
		},
	}
}

// Uint8 creates a uint8 field
func Uint8(name string, value uint8) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Uint8(name, value)
		},
	}
}

// String creates a string field
func String(name string, value string) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.String(name, value)
		},
	}
}

// Object creates an object field
func Object(name string, fieldEnc encoder.FieldEncoderFunc) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Object(name, fieldEnc)
		},
	}
}

// Bools creates a []bool field
func Bools(name string, values []bool) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Array(name, func(enc encoder.ArrayEncoder) {
				for _, v := range values {
					enc.Bool(v)
				}
			})
		},
	}
}

// Complex128s creates a []complex128 field
func Complex128s(name string, values []complex128) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Array(name, func(enc encoder.ArrayEncoder) {
				for _, v := range values {
					enc.Complex128(v)
				}
			})
		},
	}
}

// Complex64s creates a []complex64 field
func Complex64s(name string, values []complex64) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Array(name, func(enc encoder.ArrayEncoder) {
				for _, v := range values {
					enc.Complex64(v)
				}
			})
		},
	}
}

// Float64s creates a []float64 field
func Float64s(name string, values []float64) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Array(name, func(enc encoder.ArrayEncoder) {
				for _, v := range values {
					enc.Float64(v)
				}
			})
		},
	}
}

// Float32s creates a []float32 field
func Float32s(name string, values []float32) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Array(name, func(enc encoder.ArrayEncoder) {
				for _, v := range values {
					enc.Float32(v)
				}
			})
		},
	}
}

// Int64s creates a []Int64 field
func Int64s(name string, values []int64) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Array(name, func(enc encoder.ArrayEncoder) {
				for _, v := range values {
					enc.Int64(v)
				}
			})
		},
	}
}

// Int32s creates a []int32 field
func Int32s(name string, values []int32) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Array(name, func(enc encoder.ArrayEncoder) {
				for _, v := range values {
					enc.Int32(v)
				}
			})
		},
	}
}

// Int16s creates a []int16 field
func Int16s(name string, values []int16) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Array(name, func(enc encoder.ArrayEncoder) {
				for _, v := range values {
					enc.Int16(v)
				}
			})
		},
	}
}

// Int8s creates a []int8 field
func Int8s(name string, values []int8) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Array(name, func(enc encoder.ArrayEncoder) {
				for _, v := range values {
					enc.Int8(v)
				}
			})
		},
	}
}

// Uint64s creates a []uint64 field
func Uint64s(name string, values []uint64) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Array(name, func(enc encoder.ArrayEncoder) {
				for _, v := range values {
					enc.Uint64(v)
				}
			})
		},
	}
}

// Uint32s creates a []uint32 field
func Uint32s(name string, values []uint32) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Array(name, func(enc encoder.ArrayEncoder) {
				for _, v := range values {
					enc.Uint32(v)
				}
			})
		},
	}
}

// Uint16s creates a []uint16 field
func Uint16s(name string, values []uint16) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Array(name, func(enc encoder.ArrayEncoder) {
				for _, v := range values {
					enc.Uint16(v)
				}
			})
		},
	}
}

// Uint8s creates a []uint8 field
func Uint8s(name string, values []uint8) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Array(name, func(enc encoder.ArrayEncoder) {
				for _, v := range values {
					enc.Uint8(v)
				}
			})
		},
	}
}

// Strings creates a []string field
func Strings(name string, values []string) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Array(name, func(enc encoder.ArrayEncoder) {
				for _, v := range values {
					enc.String(v)
				}
			})
		},
	}
}

// Error creates an error field
func Error(name string, err error) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			encodable, ok := err.(encoder.Encodable)
			if ok {
				enc.Object(name, encodable.Encode)
			} else {
				enc.String(name, err.Error())
			}
		},
	}
}

// Encodable creates a field that encodes itself
func Encodable(name string, value encoder.Encodable) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			enc.Object(name, value.Encode)
		},
	}
}
