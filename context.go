package golog

import (
	"github.com/stugotech/golog/encoder"
	"github.com/stugotech/golog/fields"
)

// Context represents a logging context
type Context interface {
	// AddStandardFields adds one or more standard fields to the context
	AddStandardFields(field ...*fields.Field)
	// Derive creates a derived context
	Derive() Context
}

// context represents a logging context
type context struct {
	// enc specifies the encoder to use for the logging
	enc encoder.LogEncoder
	// stdFields contains a list of fields to add to all logs
	stdFields []*fields.Field
	// level is the current logging level
	level LogLevel
}

// NewContext creates a new context
func NewContext(enc encoder.LogEncoder, stdFields ...*fields.Field) Context {
	ctx := &context{
		stdFields: make([]*fields.Field, len(fields)),
	}

	for _, field := range stdFields {
		ctx.stdFields = append(ctx.stdFields, field)
	}

	return ctx
}

// AddStandardFields adds one or more standard fields to the context
func (c *context) AddStandardFields(stdFields ...*Field) {
	for _, f := range stdFields {
		c.stdFields = append(c.stdFields, f)
	}
}

// Derive creates a derived context
func (c *context) Derive() Context {
	derived := &context{
		enc:       c.enc,
		stdFields: make([]*fields.Field, len(c.stdFields)),
	}

	for _, f := range c.stdFields {
		derived.stdFields = append(derived.stdFields, f)
	}

	return derived
}
