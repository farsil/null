package null

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

// Uint implements a nullable uint.
type Uint struct {
	// Uint holds the underlying uint value.
	Uint uint

	// Valid holds the validity flag. If true, the underlying value is valid.
	// If false, it is invalid, and thus meaningless.
	Valid bool
}

// UintFrom creates a valid Uint from v.
func UintFrom(v uint) Uint {
	return UintFromPtr(&v)
}

// UintFromPtr creates an Uint from pointer p. If p is nil,
// the returned Uint is invalid.
func UintFromPtr(p *uint) Uint {
	if p != nil {
		return Uint{
			Uint:  *p,
			Valid: true,
		}
	}
	return Uint{}
}

// UintFromZero creates an Uint from v. If v is 0,
// the returned Uint is invalid.
func UintFromZero(v uint) Uint {
	return Uint{
		Uint:  v,
		Valid: v != 0,
	}
}

// Ptr returns a pointer to the underlying value of u if u is valid,
// otherwise returns nil.
func (u Uint) Ptr() *uint {
	if u.Valid {
		return &u.Uint
	}
	return nil
}

// Zero returns the underlying value of u if u is valid, otherwise
// returns 0.
func (u Uint) Zero() uint {
	if u.Valid {
		return u.Uint
	}
	return 0
}

// From sets the underlying value of u to v. u becomes valid.
func (u *Uint) From(v uint) {
	u.Valid = true
	u.Uint = v
}

// FromPtr invalidates u if p is nil, otherwise it sets the underlying value
// of u to the value pointed to by p, and u becomes valid.
func (u *Uint) FromPtr(p *uint) {
	u.Valid = p != nil
	if p != nil {
		u.Uint = *p
	}
}

// FromZero invalidates u if v is false,
// otherwise it sets the underlying value of u to v, and u becomes valid.
func (u *Uint) FromZero(v uint) {
	u.Valid = v != 0
	u.Uint = v
}

// String returns a string representation of u.
// If u is valid, it returns a string representation of the underlying value
// of u, otherwise it returns InvalidNullableString.
func (u Uint) String() string {
	if u.Valid {
		return strconv.FormatUint(uint64(u.Uint), 10)
	}
	return InvalidNullableString
}

// MarshalText marshals u to a byte string representation.
// If u is valid, it marshals the underlying value of u to a byte string
// representation, otherwise it returns nil. err is always nil.
func (u Uint) MarshalText() (data []byte, err error) {
	if u.Valid {
		return []byte(u.String()), nil
	}
	return nil, nil
}

// MarshalJSON encodes the underlying value of u to a JSON number if u is
// valid, otherwise it returns the JSON null value. If the underlying value of u
// cannot be marshaled, a MarshalError is returned.
func (u Uint) MarshalJSON() (data []byte, err error) {
	if u.Valid {
		bytes, err := json.Marshal(u.Uint)
		if err != nil {
			return nil, makeMarshalError("json", u)
		}
		return bytes, nil
	}
	return jNull, nil
}

// Value returns the underlying value of u converted to int64 if u is valid,
// otherwise nil. If the conversion would cause data loss,
// a ConversionError is returned.
func (u Uint) Value() (v driver.Value, err error) {
	if u.Valid {
		val := int64(u.Uint)
		if val < 0 {
			return nil, makeConversionError("sql", u.Uint, val)
		}
		return val, nil
	}
	return nil, nil
}

// Set invalidates u if str is the empty string, otherwise it parses str into
// the underlying value of u, and u becomes valid. If str is not a valid
// string representation of an unsigned integer, or if the represented integer
// is too large to be stored in an uint, u becomes invalid and a ParseError
// is returned.
func (u *Uint) Set(str string) error {
	val, err := strconv.ParseUint(str, 0, intSize)
	u.Uint = uint(val)
	u.Valid = err == nil && str != ""
	if str == "" || err == nil {
		return nil
	}

	return makeParseError("parse", str, u.Uint)
}

// UnmarshalText unmarshals from a byte string to u.
// It behaves like Set, except that it returns an UnmarshalError instead of a
// ParseError in case text cannot be parsed.
func (u *Uint) UnmarshalText(text []byte) error {
	if u.Set(string(text)) != nil {
		return makeUnmarshalError("parse", text, *u)
	}
	return nil
}

// UnmarshalJSON unmarshals from a JSON encoded byte string to u.
// If the encoded JSON data represent the JSON null value,
// or an error is produced, u becomes invalid. If the encoded JSON data
// represent a JSON number, and can be stored in an uint without data
// loss, u becomes valid, and the underlying value of u is set to the JSON
// number. If the encoded JSON data represent a JSON number,
// and cannot be stored in an uint without data loss, u becomes invalid,
// and a ConversionError is returned. Other JSON types
// produce a TypeError. Malformed JSON produces an UnmarshalError.
func (u *Uint) UnmarshalJSON(data []byte) error {
	var obj interface{}
	if json.Unmarshal(data, &obj) != nil {
		u.Valid = false
		return makeUnmarshalError("json", data, *u)
	}
	switch value := obj.(type) {
	case float64:
		u.Uint = uint(value)
		u.Valid = value == float64(u.Uint)
		if !u.Valid {
			return makeConversionError("json", value, u.Uint)
		}
		return nil
	case nil:
		u.Valid = false
		return nil
	default:
		u.Valid = false
		return makeTypeError("json", value, "float64", "nil")
	}
}

// Scan assigns a value from a database driver. If obj's type is int64,
// if the number can be stored in an uint without data loss,
// u becomes valid, and the underlying value of u becomes the value of obj.
// If obj's type is int64, and the number cannot be stored in an uint without
// data loss, u becomes invalid, and a ConversionError is returned.
// If obj is nil, u becomes invalid. If obj's type is any other type,
// u becomes invalid, and a TypeError is returned.
func (u *Uint) Scan(obj interface{}) error {
	switch value := obj.(type) {
	case int64:
		u.Uint = uint(value)
		u.Valid = value == int64(u.Uint) && value >= 0
		if !u.Valid {
			return makeConversionError("sql", value, u.Uint)
		}
		return nil
	case nil:
		u.Valid = false
		return nil
	default:
		u.Valid = false
		return makeTypeError("sql", value, "int64", "nil")
	}
}
