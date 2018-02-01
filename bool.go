package null

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

// Provides a nullable Bool. If Valid is false, the value stored in Bool is
// meaningless.
type Bool struct {
	Bool  bool
	Valid bool
}

// Creates a valid Bool from boolean v.
func BoolFrom(v bool) Bool {
	return BoolFromPtr(&v)
}

// Creates a Bool from boolean pointer v. If v is nil,
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

// Creates a Bool from boolean v. If v is false,
// the returned Bool is invalid.
func BoolFromZero(v bool) Bool {
	return Bool{
		Bool:  v,
		Valid: v == true,
	}
}

// Returns a pointer to the underlying boolean if Bool is valid, otherwise
// returns nil.
func (b Bool) Ptr() *bool {
	if b.Valid {
		return &b.Bool
	}
	return nil
}

// Returns a the underlying boolean if Bool is valid, otherwise
// returns false.
func (b Bool) Zero() bool {
	if b.Valid {
		return b.Bool
	}
	return false
}

// Sets the underlying boolean to v. Bool becomes valid.
func (b *Bool) From(v bool) {
	b.Valid = true
	b.Bool = v
}

// If v is nil, Bool becomes invalid, otherwise it sets the underlying boolean
// to the value pointed to by p, and Bool becomes valid.
func (b *Bool) FromPtr(p *bool) {
	b.Valid = p != nil
	if p != nil {
		b.Bool = *p
	}
}

// If v is false, Bool becomes invalid, otherwise it sets the underlying boolean
// to v, and Bool becomes valid.
func (b *Bool) FromZero(v bool) {
	b.Valid = v == true
	b.Bool = v
}

// Returns a string representation of Bool.
func (b Bool) String() string {
	if b.Valid {
		return strconv.FormatBool(b.Bool)
	}
	return sInvalid
}

// Marshals the boolean true to "true" and the boolean false to "false" if
// Bool is valid, otherwise it marshals to nil. err is always nil.
func (b Bool) MarshalText() (data []byte, err error) {
	if b.Valid {
		return []byte(b.String()), nil
	}
	return nil, nil
}

// Marshals to a JSON boolean if Bool is valid,
// otherwise it marshals to the JSON null value. err is always nil.
func (b Bool) MarshalJSON() (data []byte, err error) {
	if b.Valid {
		if b.Bool {
			return jTrue, nil
		} else {
			return jFalse, nil
		}
	}
	return jNull, nil
}

// Returns the underlying boolean if Bool is valid,
// otherwise nil. err is always nil.
func (b Bool) Value() (v driver.Value, err error) {
	if b.Valid {
		return b.Bool, nil
	}
	return nil, nil
}

// Parses a string representation of a boolean. If str is an empty string,
// Bool becomes invalid. If str is a valid boolean string, Bool becomes valid,
// str is parsed, and the parsed value becomes the underlying value.
// If str is an invalid boolean string, Bool becomes invalid, and a ParseError
// is returned. Recognized boolean strings are 1, true, True, TRUE, T, t, 0,
// false, False, FALSE, f, F.
func (b *Bool) Set(str string) error {
	var err error
	b.Bool, err = strconv.ParseBool(str)
	b.Valid = err == nil && str != ""

	if err == nil || str == "" {
		return nil
	}
	return makeParseError("parse", str, b.Bool)
}

// Unmarshals from a string representation of a boolean. It behaves like Set,
// except that it returns an UnmarshalError instead of a ParseError if
// text is an invalid boolean string.
func (b *Bool) UnmarshalText(text []byte) error {
	if b.Set(string(text)) != nil {
		return makeUnmarshalError("text", text, *b)
	}
	return nil
}

// Unmarshals from a JSON object. If the JSON object is the JSON null value,
// or an error is produced, Bool becomes invalid. If the JSON object is a JSON
// boolean, Bool becomes valid, the object is parsed,
// and the parsed value becomes the underlying value. Malformed JSON produces
// an UnmarshalError, while unrecognized JSON types produce a TypeError.
func (b *Bool) UnmarshalJSON(data []byte) error {
	var obj interface{}
	if json.Unmarshal(data, &obj) != nil {
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

// Assigns a value from a database driver. If obj's type is bool,
// Bool becomes valid, and the underlying boolean becomes the value of obj.
// If obj is nil, Bool becomes invalid. If obj's type is any other type,
// Bool becomes invalid, and a TypeError is returned.
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
