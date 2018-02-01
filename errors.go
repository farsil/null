package null

import (
	"fmt"
	"reflect"
	"strings"
)

type MarshalError struct {
	prefix   string
	SrcValue interface{}
	SrcType  string
}

type UnmarshalError struct {
	prefix   string
	SrcValue string
	DestType string
}

type ConversionError struct {
	prefix   string
	SrcValue interface{}
	SrcType  string
	DestType string
}

type ParseError struct {
	prefix    string
	SrcString string
	DestType  string
}

type TypeError struct {
	prefix        string
	InvalidType   string
	ExpectedTypes []string
}

func makeConversionError(prefix string, src, dest interface{}) error {
	return ConversionError{
		prefix:   prefix,
		SrcValue: src,
		SrcType:  reflect.TypeOf(src).Name(),
		DestType: reflect.TypeOf(dest).Name(),
	}
}

func makeTypeError(prefix string, src interface{}, exp ...string) error {
	return TypeError{
		prefix:        prefix,
		InvalidType:   reflect.TypeOf(src).Name(),
		ExpectedTypes: exp,
	}
}

func makeParseError(prefix, src string, dest interface{}) error {
	return ParseError{
		prefix:    prefix,
		SrcString: src,
		DestType:  reflect.TypeOf(dest).Name(),
	}
}

func makeMarshalError(prefix string, src interface{}) error {
	return MarshalError{
		prefix:   prefix,
		SrcValue: src,
		SrcType:  reflect.TypeOf(src).Name(),
	}
}

func makeUnmarshalError(prefix string, src []byte, dest interface{}) error {
	return UnmarshalError{
		prefix:   prefix,
		SrcValue: string(src),
		DestType: reflect.TypeOf(dest).Name(),
	}
}

func (e ConversionError) Error() string {
	return fmt.Sprintf(
		"%s: cannot convert %v of type %s into Go value of type %s",
		e.prefix, e.SrcValue, e.SrcType, e.DestType,
	)
}

func (e TypeError) Error() string {
	return fmt.Sprintf(
		"%s: invalid type (expected %s, got %s)",
		e.prefix, strings.Join(e.ExpectedTypes, " or "), e.InvalidType,
	)
}

func (e ParseError) Error() string {
	return fmt.Sprintf(
		"%s: cannot parse '%s' into Go value of type %s",
		e.prefix, e.SrcString, e.DestType,
	)
}

func (e MarshalError) Error() string {
	return fmt.Sprintf(
		"%s: cannot marshal %v of type %s",
		e.prefix, e.SrcValue, e.SrcType,
	)
}

func (e UnmarshalError) Error() string {
	return fmt.Sprintf(
		"%s: cannot unmarshal '%s' into Go value of type %s",
		e.prefix, e.SrcValue, e.DestType,
	)
}
