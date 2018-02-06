package null

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

// Bool implements a nullable bool.
type Bool struct {
	// Bool holds the underlying bool value.
	Bool bool

	// Valid holds the validity flag. If true, the underlying value is valid.
	// If false, it is invalid, and thus meaningless.
	Valid bool
}

// BoolFrom creates a valid Bool from v.
func BoolFrom(v bool) Bool {
	return BoolFromPtr(&v)
}

// BoolFromPtr creates a Bool from pointer p. If p is nil,
// the returned Bool is invalid.
func BoolFromPtr(p *bool) Bool {
	if p != nil {
		return Bool{
			Bool:  *p,
			Valid: true,
		}
	}
	return Bool{}
}

// BoolFromZero creates a Bool from v. If v is false,
// the returned Bool is invalid.
func BoolFromZero(v bool) Bool {
	return Bool{
		Bool:  v,
		Valid: v == true,
	}
}

// Ptr returns a pointer to the underlying value of b if b is valid,
// otherwise returns nil.
func (b Bool) Ptr() *bool {
	if b.Valid {
		return &b.Bool
	}
	return nil
}

// Zero returns the underlying value of b if b is valid, otherwise
// returns false.
func (b Bool) Zero() bool {
	if b.Valid {
		return b.Bool
	}
	return false
}

// From sets the underlying value of b to v. b becomes valid.
func (b *Bool) From(v bool) {
	b.Valid = true
	b.Bool = v
}

// FromPtr invalidates b if p is nil, otherwise it sets the underlying value
// of b to the value pointed to by p, and b becomes valid.
func (b *Bool) FromPtr(p *bool) {
	b.Valid = p != nil
	if p != nil {
		b.Bool = *p
	}
}

// FromZero invalidates b if v is false,
// otherwise it sets the underlying value of b to v, and b becomes valid.
func (b *Bool) FromZero(v bool) {
	b.Valid = v == true
	b.Bool = v
}

// String returns a string representation of b. If b is valid,
// it returns either "true" or "false", otherwise it returns
// InvalidNullableString.
func (b Bool) String() string {
	if b.Valid {
		return strconv.FormatBool(b.Bool)
	}
	return InvalidNullableString
}

// MarshalText marshals b to a byte string representation.
// If b is valid, it marshals the underlying value of b to either
// "true" or "false", otherwise it returns nil. err is always nil.
func (b Bool) MarshalText() (data []byte, err error) {
	if b.Valid {
		return []byte(b.String()), nil
	}
	return nil, nil
}

// MarshalJSON encodes the underlying value of b to a JSON boolean
// if b is valid, otherwise it returns the JSON null value.
// err is always nil.
func (b Bool) MarshalJSON() (data []byte, err error) {
	if b.Valid {
		if b.Bool {
			return jTrue, nil
		}
		return jFalse, nil
	}
	return jNull, nil
}

// Value returns the underlying value of b if b is valid,
// otherwise nil. err is always nil.
func (b Bool) Value() (v driver.Value, err error) {
	if b.Valid {
		return b.Bool, nil
	}
	return nil, nil
}

// Set invalidates b if str is the empty string, otherwise it parses str into
// the underlying value of b, and b becomes valid. If str is not a recognized
// boolean string, b becomes invalid and a ParseError is returned.
// Recognized boolean strings are 1, true, True, TRUE, T, t, 0, false, False,
// FALSE, f, F.
func (b *Bool) Set(str string) error {
	var err error
	b.Bool, err = strconv.ParseBool(str)
	b.Valid = err == nil && str != ""

	if err == nil || str == "" {
		return nil
	}
	return makeParseError("parse", str, b.Bool)
}

// UnmarshalText unmarshals from a byte string to b.
// It behaves like Set, except that it returns an UnmarshalError instead of a
// ParseError in case text cannot be parsed.
func (b *Bool) UnmarshalText(text []byte) error {
	if b.Set(string(text)) != nil {
		return makeUnmarshalError("text", text, *b)
	}
	return nil
}

// UnmarshalJSON unmarshals from a JSON encoded byte string to b.
// If the encoded JSON data represent the JSON null value,
// or an error is produced, b becomes invalid. If the encoded JSON data
// represent a JSON boolean, b becomes valid, and the underlying
// value of b is set to the JSON boolean. Other JSON types
// produce a TypeError. Malformed JSON produces an UnmarshalError.
func (b *Bool) UnmarshalJSON(data []byte) error {
	var obj interface{}
	if json.Unmarshal(data, &obj) != nil {
		b.Valid = false
		return makeUnmarshalError("json", data, *b)
	}
	switch value := obj.(type) {
	case bool:
		b.Bool = value
		b.Valid = true
		return nil
	case nil:
		b.Valid = false
		return nil
	default:
		b.Valid = false
		return makeTypeError("json", value, "bool", "nil")
	}
}

// Scan assigns a value from a database driver. If obj's type is bool,
// b becomes valid, and the underlying value of b becomes the value of obj.
// If obj is nil, b becomes invalid. If obj's type is any other type,
// b becomes invalid, and a TypeError is returned.
func (b *Bool) Scan(obj interface{}) error {
	switch value := obj.(type) {
	case bool:
		b.Bool = value
		b.Valid = true
		return nil
	case nil:
		b.Valid = false
		return nil
	default:
		b.Valid = false
		return makeTypeError("sql", value, "bool", "nil")
	}
}
