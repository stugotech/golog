package encoder

// ArrayEncoderFunc adds elements to an ArrayEncoder
type ArrayEncoderFunc func(ArrayEncoder)

// FieldEncoderFunc adds members to an object Encoder
type FieldEncoderFunc func(FieldEncoder)

// An Encodable can encode itself
type Encodable interface {
	// Encode writes the current instance to the encoder
	Encode(enc FieldEncoder)
}

// LogEncoder allows logs to be added
type LogEncoder interface {
	AddEntry(enc FieldEncoderFunc)
}

// An FieldEncoder deals with converting log information into another format
type FieldEncoder interface {
	Bool(name string, value bool)

	Complex128(name string, value complex128)
	Complex64(name string, value complex64)
	Float64(name string, value float64)
	Float32(name string, value float32)

	Int64(name string, value int64)
	Int32(name string, value int32)
	Int16(name string, value int16)
	Int8(name string, value int8)

	Uint64(name string, value uint64)
	Uint32(name string, value uint32)
	Uint16(name string, value uint16)
	Uint8(name string, value uint8)

	String(name string, value string)
	Array(name string, encoder ArrayEncoderFunc)
	Object(name string, encoder FieldEncoderFunc)
}

// An ArrayEncoder appends elements to an array
type ArrayEncoder interface {
	Array(encoder ArrayEncoderFunc)
	Object(encoder FieldEncoderFunc)
	Bool(value bool)

	Complex128(value complex128)
	Complex64(value complex64)
	Float64(value float64)
	Float32(value float32)

	Int64(value int64)
	Int32(value int32)
	Int16(value int16)
	Int8(value int8)
	Uint64(value uint64)
	Uint32(value uint32)
	Uint16(value uint16)
	Uint8(value uint8)

	String(value string)
}
