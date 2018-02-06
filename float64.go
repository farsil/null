package null

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

// Float64 implements a nullable float64.
type Float64 struct {
	// Float64 holds the underlying float64 value.
	Float64 float64

	// Valid holds the validity flag. If true, the underlying value is valid.
	// If false, it is invalid, and thus meaningless.
	Valid bool
}

// Float64From creates a valid Float64 from v.
func Float64From(v float64) Float64 {
	return Float64FromPtr(&v)
}

// Float64FromPtr creates a Float64 from pointer p. If p is nil,
// the returned Float64 is invalid.
func Float64FromPtr(p *float64) Float64 {
	if p != nil {
		return Float64{
			Float64: *p,
			Valid:   true,
		}
	}
	return Float64{}
}

// Float64FromZero creates a Float64 from v. If v is 0,
// the returned Float64 is invalid.
func Float64FromZero(v float64) Float64 {
	return Float64{
		Float64: v,
		Valid:   v != 0.0,
	}
}

// Ptr returns a pointer to the underlying value of f if f is valid,
// otherwise returns nil.
func (f Float64) Ptr() *float64 {
	if f.Valid {
		return &f.Float64
	}
	return nil
}

// Zero returns the underlying value of f if f is valid, otherwise
// returns 0.
func (f Float64) Zero() float64 {
	if f.Valid {
		return f.Float64
	}
	return 0.0
}

// From sets the underlying value of f to v. f becomes valid.
func (f *Float64) From(v float64) {
	f.Valid = true
	f.Float64 = v
}

// FromPtr invalidates f if p is nil, otherwise it sets the underlying value
// of f to the value pointed to by p, and f becomes valid.
func (f *Float64) FromPtr(p *float64) {
	f.Valid = p != nil
	if p != nil {
		f.Float64 = *p
	}
}

// FromZero invalidates f if v is false,
// otherwise it sets the underlying value of f to v, and f becomes valid.
func (f *Float64) FromZero(v float64) {
	f.Valid = v != 0.0
	f.Float64 = v
}

// String returns a string representation of f.
// If f is valid, it returns a string representation of the underlying value
// of f, otherwise it returns InvalidNullableString.
func (f Float64) String() string {
	if f.Valid {
		return strconv.FormatFloat(f.Float64, 'g', -1, 64)
	}
	return InvalidNullableString
}

// MarshalText marshals f to a byte string representation.
// If f is valid, it marshals the underlying value of f to a byte string
// representation, otherwise it returns nil. err is always nil.
func (f Float64) MarshalText() (data []byte, err error) {
	if f.Valid {
		return []byte(f.String()), nil
	}
	return nil, nil
}

// MarshalJSON encodes the underlying value of f to a JSON number if f is
// valid, otherwise it returns the JSON null value. If the underlying value of f
// cannot be marshaled, a MarshalError is returned.
func (f Float64) MarshalJSON() (data []byte, err error) {
	if f.Valid {
		bytes, err := json.Marshal(f.Float64)
		if err != nil {
			return nil, makeMarshalError("json", f)
		}
		return bytes, nil
	}
	return jNull, nil
}

// Value returns the underlying value of f if f is valid,
// otherwise nil. err is always nil.
func (f Float64) Value() (v driver.Value, err error) {
	if f.Valid {
		return f.Float64, nil
	}
	return nil, nil
}

// Set invalidates f if str is the empty string, otherwise it parses str into
// the underlying value of f, and f becomes valid. If str is not a valid
// string representation of a floating point number,
// f becomes invalid and a ParseError is returned.
func (f *Float64) Set(str string) error {
	var err error
	f.Float64, err = strconv.ParseFloat(str, 64)
	f.Valid = err == nil && str != ""
	if str == "" || err == nil {
		return nil
	}

	return makeParseError("parse", str, f.Float64)
}

// UnmarshalText unmarshals from a byte string to f.
// It behaves like Set, except that it returns an UnmarshalError instead of a
// ParseError in case text cannot be parsed.
func (f *Float64) UnmarshalText(text []byte) error {
	if f.Set(string(text)) != nil {
		return makeUnmarshalError("parse", text, *f)
	}
	return nil
}

// UnmarshalJSON unmarshals from a JSON encoded byte string to f.
// If the encoded JSON data represent the JSON null value,
// or an error is produced, f becomes invalid. If the encoded JSON data
// represent a JSON number, f becomes valid,
// and the underlying value of f is set to the JSON
// number. Other JSON types produce a TypeError. Malformed JSON produces
// an UnmarshalError.
func (f *Float64) UnmarshalJSON(data []byte) error {
	var obj interface{}
	if json.Unmarshal(data, &obj) != nil {
		return makeUnmarshalError("json", data, *f)
	}
	switch value := obj.(type) {
	case float64:
		f.Float64 = value
		f.Valid = true
		return nil
	case nil:
		f.Valid = false
		return nil
	default:
		return makeTypeError("json", value, "float64", "nil")
	}
}

// Scan assigns a value from a database driver. If obj's type is float64,
// f becomes valid, and the underlying value of f becomes the value of obj.
// If obj is nil, f becomes invalid. If obj's type is any other type,
// f becomes invalid, and a TypeError is returned.
func (f *Float64) Scan(obj interface{}) error {
	switch value := obj.(type) {
	case float64:
		f.Float64 = value
		f.Valid = true
		return nil
	case nil:
		f.Valid = false
		return nil
	default:
		return makeTypeError("sql", value, "float64", "nil")
	}
}
