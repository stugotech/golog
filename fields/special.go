package fields

import (
	"runtime"

	"github.com/stugotech/golog/encoder"
)

// LineNumber creates a field containing the line number it was called from
func LineNumber(name string) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			_, _, line, ok := runtime.Caller(2)
			if !ok {
				return
			}
			enc.Int64("line", int64(line))
		},
	}
}

// PostponeLineNumber creates a field containing the line number Log was called from
func PostponeLineNumber(name string) *Field {
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			_, _, line, ok := runtime.Caller(2)
			if !ok {
				return
			}
			enc.Int64(name, int64(line))
		},
	}
}

// FileName creates a field containing the filename that this func was called from
func FileName(name string) *Field {
	_, caller, _, ok := runtime.Caller(2)
	return &Field{
		Name: name,
		Encode: func(enc encoder.FieldEncoder) {
			if !ok {
				return
			}
			enc.String(name, caller)
		},
	}
}
