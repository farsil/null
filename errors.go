package null

import (
	"fmt"
	"reflect"
	"strings"
)

// MarshalError is returned if marshaling is not successful.
type MarshalError struct {
	// prefix indicates the error class
	prefix string

	// SrcValue holds the value that couldn't be marshaled
	SrcValue interface{}

	// SrcType holds the name of SrcValue's type
	SrcType string
}

// helper function to build a MarshalError.
func makeMarshalError(prefix string, src interface{}) error {
	return MarshalError{
		prefix:   prefix,
		SrcValue: src,
		SrcType:  reflect.TypeOf(src).Name(),
	}
}

// Error returns a string representation of e.
func (e MarshalError) Error() string {
	return fmt.Sprintf(
		"%s: cannot marshal %v of type %s",
		e.prefix, e.SrcValue, e.SrcType,
	)
}

// UnmarshalError is returned if unmarshaling is not successful.
type UnmarshalError struct {
	// prefix indicates the error class
	prefix string

	// SrcValue holds the data that couldn't be unmarshaled
	SrcValue string

	// DestType holds the destination type
	DestType string
}

// helper function to build a UnmarshalError.
func makeUnmarshalError(prefix string, src []byte, dest interface{}) error {
	return UnmarshalError{
		prefix:   prefix,
		SrcValue: string(src),
		DestType: reflect.TypeOf(dest).Name(),
	}
}

// Error returns a string representation of e.
func (e UnmarshalError) Error() string {
	return fmt.Sprintf(
		"%s: cannot unmarshal '%s' into Go value of type %s",
		e.prefix, e.SrcValue, e.DestType,
	)
}

// ConversionError is returned when a type conversion cannot happen without
// data loss.
type ConversionError struct {
	// prefix indicates the error class
	prefix string

	// SrcValue holds the value that couldn't be converted
	SrcValue interface{}

	// SrcType holds the name of SrcValue's type
	SrcType string

	// DestType holds the destination type
	DestType string
}

// helper function to build a ConversionError.
func makeConversionError(prefix string, src, dest interface{}) error {
	return ConversionError{
		prefix:   prefix,
		SrcValue: src,
		SrcType:  reflect.TypeOf(src).Name(),
		DestType: reflect.TypeOf(dest).Name(),
	}
}

// Error returns a string representation of e.
func (e ConversionError) Error() string {
	return fmt.Sprintf(
		"%s: cannot convert %v of type %s into Go value of type %s",
		e.prefix, e.SrcValue, e.SrcType, e.DestType,
	)
}

// ParseError is returned when a string cannot be parsed.
type ParseError struct {
	// prefix indicates the error class
	prefix string

	// SrcString holds the string that couldn't be parsed
	SrcString string

	// DestType holds the destination type
	DestType string
}

// helper function to build a ParseError.
func makeParseError(prefix, src string, dest interface{}) error {
	return ParseError{
		prefix:    prefix,
		SrcString: src,
		DestType:  reflect.TypeOf(dest).Name(),
	}
}

// Error returns a string representation of e.
func (e ParseError) Error() string {
	return fmt.Sprintf(
		"%s: cannot parse '%s' into Go value of type %s",
		e.prefix, e.SrcString, e.DestType,
	)
}

// TypeError is returned when a value of the wrong type is supplied.
type TypeError struct {
	// prefix indicates the error class
	prefix string

	// InvalidType holds the name of the type that is not valid
	InvalidType string

	// ExpectedTypes holds the names of accepted types
	ExpectedTypes []string
}

// helper function to build a TypeError.
func makeTypeError(prefix string, src interface{}, exp ...string) error {
	return TypeError{
		prefix:        prefix,
		InvalidType:   reflect.TypeOf(src).Name(),
		ExpectedTypes: exp,
	}
}

// Error returns a string representation of e.
func (e TypeError) Error() string {
	return fmt.Sprintf(
		"%s: invalid type (expected %s, got %s)",
		e.prefix, strings.Join(e.ExpectedTypes, " or "), e.InvalidType,
	)
}
