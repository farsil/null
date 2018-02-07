package null

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

// Int implements a nullable int.
type Int struct {
	// Int holds the underlying int value.
	Int int

	// Valid holds the validity flag. If true, the underlying value is valid.
	// If false, it is invalid, and thus meaningless.
	Valid bool
}

// IntFrom creates a valid Int from v.
func IntFrom(v int) Int {
	return IntFromPtr(&v)
}

// IntFromPtr creates an Int from pointer p. If p is nil,
// the returned Int is invalid.
func IntFromPtr(p *int) Int {
	if p != nil {
		return Int{
			Int:   *p,
			Valid: true,
		}
	}
	return Int{}
}

// IntFromZero creates an Int from v. If v is 0,
// the returned Int is invalid.
func IntFromZero(v int) Int {
	return Int{
		Int:   v,
		Valid: v != 0,
	}
}

// Ptr returns a pointer to the underlying value of i if i is valid,
// otherwise returns nil.
func (i Int) Ptr() *int {
	if i.Valid {
		return &i.Int
	}
	return nil
}

// Zero returns the underlying value of i if i is valid, otherwise
// returns 0.
func (i Int) Zero() int {
	if i.Valid {
		return i.Int
	}
	return 0
}

// From sets the underlying value of i to v. i becomes valid.
func (i *Int) From(v int) {
	i.Valid = true
	i.Int = v
}

// FromPtr invalidates i if p is nil, otherwise it sets the underlying value
// of i to the value pointed to by p, and i becomes valid.
func (i *Int) FromPtr(p *int) {
	i.Valid = p != nil
	if p != nil {
		i.Int = *p
	}
}

// FromZero invalidates i if v is false,
// otherwise it sets the underlying value of i to v, and i becomes valid.
func (i *Int) FromZero(v int) {
	i.Valid = v != 0
	i.Int = v
}

// String returns a string representation of i.
// If i is valid, it returns a string representation of the underlying value
// of i, otherwise it returns InvalidNullableString.
func (i Int) String() string {
	if i.Valid {
		return strconv.FormatInt(int64(i.Int), 10)
	}
	return InvalidNullableString
}

// MarshalText marshals i to a byte string representation.
// If i is valid, it marshals the underlying value of i to a byte string
// representation, otherwise it returns nil. err is always nil.
func (i Int) MarshalText() (data []byte, err error) {
	if i.Valid {
		return []byte(i.String()), nil
	}
	return nil, nil
}

// MarshalJSON encodes the underlying value of i to a JSON number if i is
// valid, otherwise it returns the JSON null value. err is always nil.
func (i Int) MarshalJSON() (data []byte, err error) {
	if i.Valid {
		return []byte(i.String()), nil
	}
	return jNull, nil
}

// Value returns the underlying value of i converted to int64 if i is valid,
// otherwise nil. err is always nil.
func (i Int) Value() (v driver.Value, err error) {
	if i.Valid {
		return int64(i.Int), nil
	}
	return nil, nil
}

// Set invalidates i if str is the empty string, otherwise it parses str into
// the underlying value of i, and i becomes valid. If str is not a valid
// string representation of an integer, or if the represented integer is too
// large to be stored in an int, i becomes invalid and a ParseError is returned.
func (i *Int) Set(str string) error {
	val, err := strconv.ParseInt(str, 0, intSize)
	i.Int = int(val)
	i.Valid = err == nil && str != ""
	if str == "" || err == nil {
		return nil
	}

	return makeParseError("parse", str, i.Int)
}

// UnmarshalText unmarshals from a byte string to i.
// It behaves like Set, except that it returns an UnmarshalError instead of a
// ParseError in case text cannot be parsed.
func (i *Int) UnmarshalText(text []byte) error {
	if i.Set(string(text)) != nil {
		return makeUnmarshalError("parse", text, *i)
	}
	return nil
}

// UnmarshalJSON unmarshals from a JSON encoded byte string to i.
// If the encoded JSON data represent the JSON null value,
// or an error is produced, i becomes invalid. If the encoded JSON data
// represent a JSON number, and can be stored in an int without data
// loss, i becomes valid, and the underlying value of i is set to the JSON
// number. If the encoded JSON data represent a JSON number,
// and cannot be stored in an int without data loss, i becomes invalid,
// and a ConversionError is returned. Other JSON types
// produce a TypeError. Malformed JSON produces an UnmarshalError.
func (i *Int) UnmarshalJSON(data []byte) error {
	var obj interface{}
	if json.Unmarshal(data, &obj) != nil {
		i.Valid = false
		return makeUnmarshalError("json", data, *i)
	}
	switch value := obj.(type) {
	case float64:
		i.Int = int(value)
		i.Valid = value == float64(i.Int)
		if !i.Valid {
			return makeConversionError("json", value, i.Int)
		}
		return nil
	case nil:
		i.Valid = false
		return nil
	default:
		i.Valid = false
		return makeTypeError("json", value, "float64", "nil")
	}
}

// Scan assigns a value from a database driver. If obj's type is int64,
// if the number can be stored in an int without data loss,
// i becomes valid, and the underlying value of i becomes the value of obj.
// If obj's type is int64, and the number cannot be stored in an int without
// data loss, i becomes invalid, and a ConversionError is returned.
// If obj is nil, i becomes invalid. If obj's type is any other type,
// i becomes invalid, and a TypeError is returned.
func (i *Int) Scan(obj interface{}) error {
	switch value := obj.(type) {
	case int64:
		i.Int = int(value)
		i.Valid = value == int64(i.Int)
		if !i.Valid {
			return makeConversionError("sql", value, i.Int)
		}
		return nil
	case nil:
		i.Valid = false
		return nil
	default:
		i.Valid = false
		return makeTypeError("sql", value, "int64", "nil")
	}
}
