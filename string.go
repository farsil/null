package null

import (
	"database/sql/driver"
	"encoding/json"
)

// Provides a nullable string.
type String struct {
	// The underlying string value.
	// The field can't be named String because it conflicts with the
	// flag.Value interface implementation.
	Str string

	// If true, the underlying string value is valid. If false,
	// the value stored in Str is invalid, and thus meaningless.
	Valid bool
}

// Creates a valid String from v.
func StringFrom(v string) String {
	return StringFromPtr(&v)
}

// Creates a String from pointer p. If p is nil, the returned String is invalid.
func StringFromPtr(p *string) String {
	if p != nil {
		return String{
			Str:   *p,
			Valid: true,
		}
	}
	return String{}
}

// Creates a String from v. If v is the empty string,
// the returned String is invalid.
func StringFromZero(v string) String {
	return String{
		Str:   v,
		Valid: v != "",
	}
}

// Returns a pointer to the underlying string if String is valid, otherwise
// returns nil.
func (s String) Ptr() *string {
	if s.Valid {
		return &s.Str
	}
	return nil
}

// Returns the underlying string if String is valid, otherwise
// returns an empty string.
func (s String) Zero() string {
	if s.Valid {
		return s.Str
	}
	return ""
}

// Sets the underlying string to v. String becomes valid.
func (s *String) From(v string) {
	s.Valid = true
	s.Str = v
}

// If p is nil, String becomes invalid, otherwise it sets the underlying string
// to the value pointed to by p, and String becomes valid.
func (s *String) FromPtr(p *string) {
	s.Valid = p != nil
	if p != nil {
		s.Str = *p
	}
}

// If v is the empty string, String becomes invalid, otherwise it sets the
// underlying string to v, and String becomes valid.
func (s *String) FromZero(v string) {
	s.Valid = v != ""
	s.Str = v
}

// Returns the underlying string if String is valid,
// and InvalidNullableStr if not valid.
func (s String) String() string {
	if s.Valid {
		return s.Str
	}
	return InvalidNullableStr
}

// Converts the underlying value to []byte if
// String is valid, and returns nil if not valid. err is always nil.
func (s String) MarshalText() (data []byte, err error) {
	if s.Valid {
		return []byte(s.Str), nil
	}
	return nil, nil
}

// Marshals to a JSON string if String is valid,
// otherwise it marshals to the JSON null value. err is always nil.
func (s String) MarshalJSON() (data []byte, err error) {
	if s.Valid {
		return json.Marshal(s.Str)
	}
	return jNull, nil
}

// Returns the underlying string if String is valid,
// otherwise nil. err is always nil.
func (s String) Value() (v driver.Value, err error) {
	if s.Valid {
		return s.Str, nil
	}
	return nil, nil
}

// Like FromZero, if v is the empty string, String becomes invalid,
// otherwise it sets the underlying string to v, and String becomes valid.
// This function always returns nil.
func (s *String) Set(str string) error {
	s.FromZero(str)
	return nil
}

// Unmarshals from a byte string. If the byte string is nil,
// String becomes invalid, otherwise the underlying string is set to the
// byte string and String becomes valid. The returned error is always nil.
func (s *String) UnmarshalText(text []byte) error {
	s.Str = string(text)
	s.Valid = text != nil
	return nil
}

// Unmarshals from a JSON object. If the JSON object is the JSON null value,
// or an error is produced, String becomes invalid. If the JSON object is a JSON
// string, String becomes valid, and the underlying value is set to the JSON
// string. Other JSON types produce a TypeError. Malformed JSON produces an
// UnmarshalError.
func (s *String) UnmarshalJSON(data []byte) error {
	var obj interface{}
	if json.Unmarshal(data, &obj) != nil {
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

// Assigns a value from a database driver. If obj's type is string,
// String becomes valid, and the underlying string becomes the value of obj.
// If obj is nil, String becomes invalid. If obj's type is any other type,
// String becomes invalid, and a TypeError is returned.
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
