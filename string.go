package null

import (
	"database/sql/driver"
	"encoding/json"
)

// String implements a nullable string.
type String struct {
	// Str holds the underlying string value. The field can't be named String
	// because it conflicts with the flag.Value interface implementation.
	Str string

	// Valid holds the validity flag. If true, the underlying value is valid.
	// If false, it is invalid, and thus meaningless.
	Valid bool
}

// StringFrom creates a valid String from v.
func StringFrom(v string) String {
	return StringFromPtr(&v)
}

// StringFromPtr creates a String from pointer p. If p is nil,
// the returned String is invalid.
func StringFromPtr(p *string) String {
	if p != nil {
		return String{
			Str:   *p,
			Valid: true,
		}
	}
	return String{}
}

// StringFromZero creates a String from v. If v is the empty string,
// the returned String is invalid.
func StringFromZero(v string) String {
	return String{
		Str:   v,
		Valid: v != "",
	}
}

// Ptr returns a pointer to the underlying value of s if s is valid, otherwise
// returns nil.
func (s String) Ptr() *string {
	if s.Valid {
		return &s.Str
	}
	return nil
}

// Zero returns the underlying value of s if s is valid, otherwise
// returns an empty string.
func (s String) Zero() string {
	if s.Valid {
		return s.Str
	}
	return ""
}

// From sets the underlying value of s to v. s becomes valid.
func (s *String) From(v string) {
	s.Valid = true
	s.Str = v
}

// FromPtr invalidates s if p is nil, otherwise it sets the
// underlying value of s to the value pointed to by p, and s becomes valid.
func (s *String) FromPtr(p *string) {
	s.Valid = p != nil
	if p != nil {
		s.Str = *p
	}
}

// FromZero invalidates s if v is the empty string, otherwise it sets the
// underlying value of s to v, and s becomes valid.
func (s *String) FromZero(v string) {
	s.Valid = v != ""
	s.Str = v
}

// String returns the underlying value of s if s is valid,
// and InvalidNullableString if not valid.
func (s String) String() string {
	if s.Valid {
		return s.Str
	}
	return InvalidNullableString
}

// MarshalText converts the underlying value of s to []byte if
// s is valid, and returns nil if not valid. err is always nil.
func (s String) MarshalText() (data []byte, err error) {
	if s.Valid {
		return []byte(s.Str), nil
	}
	return nil, nil
}

// MarshalJSON encodes the underlying value of s to a JSON string if s is
// valid, otherwise it returns the JSON null value. err is always nil.
func (s String) MarshalJSON() (data []byte, err error) {
	if s.Valid {
		return json.Marshal(s.Str)
	}
	return jNull, nil
}

// Value returns the underlying value of s if s is valid,
// otherwise nil. err is always nil.
func (s String) Value() (v driver.Value, err error) {
	if s.Valid {
		return s.Str, nil
	}
	return nil, nil
}

// Set invalidates s if str is the empty string,
// otherwise it sets the underlying value of s to str, and s becomes valid.
// This function always returns nil.
func (s *String) Set(str string) error {
	s.FromZero(str)
	return nil
}

// UnmarshalText unmarshals from a byte string to s. If the byte string is nil,
// s becomes invalid, otherwise the underlying value of s is set to the
// converted byte string and s becomes valid.
// The returned error is always nil.
func (s *String) UnmarshalText(text []byte) error {
	s.Str = string(text)
	s.Valid = text != nil
	return nil
}

// UnmarshalJSON unmarshals from a JSON encoded byte string to s.
// If the encoded JSON data represent the JSON null value,
// or an error is produced, s becomes invalid. If the encoded JSON data
// represent a JSON string, s becomes valid, and the underlying
// value of s is set to the JSON string. Other JSON types produce a TypeError.
// Malformed JSON produces an UnmarshalError.
func (s *String) UnmarshalJSON(data []byte) error {
	var obj interface{}
	if json.Unmarshal(data, &obj) != nil {
		s.Valid = false
		return makeUnmarshalError("json", data, *s)
	}
	switch value := obj.(type) {
	case string:
		s.Str = value
		s.Valid = true
		return nil
	case nil:
		s.Valid = false
		return nil
	default:
		s.Valid = false
		return makeTypeError("json", value, "string", "nil")
	}
}

// Scan assigns a value from a database driver. If obj's type is string,
// s becomes valid, and the underlying value of s becomes the value of obj.
// If obj is nil, s becomes invalid. If obj's type is any other type,
// s becomes invalid, and a TypeError is returned.
func (s *String) Scan(obj interface{}) error {
	switch value := obj.(type) {
	case string:
		s.Str = value
		s.Valid = true
		return nil
	case nil:
		s.Valid = false
		return nil
	default:
		s.Valid = false
		return makeTypeError("sql", value, "string", "nil")
	}
}
