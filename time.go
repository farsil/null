package null

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// Time implements a nullable time.Time.
type Time struct {
	// Time holds the underlying time.Time value.
	Time time.Time

	// Valid holds the validity flag. If true, the underlying value is valid.
	// If false, it is invalid, and thus meaningless.
	Valid bool
}

// TimeFrom creates a valid Time from v.
func TimeFrom(v time.Time) Time {
	return TimeFromPtr(&v)
}

// TimeFromPtr creates a Time from pointer p. If p is nil,
// the returned Time is invalid.
func TimeFromPtr(v *time.Time) Time {
	if v != nil {
		return Time{
			Time:  *v,
			Valid: true,
		}
	}
	return Time{}
}

// TimeFromZero creates a Time from v. If v represents the zero time instant,
// the returned Time is invalid.
func TimeFromZero(v time.Time) Time {
	return Time{
		Time:  v,
		Valid: !v.IsZero(),
	}
}

// Ptr returns a pointer to the underlying value of t if t is valid,
// otherwise returns nil.
func (t Time) Ptr() *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

// Zero returns the underlying value of t if t is valid, otherwise
// returns a time.Time representing the zero time instant.
func (t Time) Zero() time.Time {
	if t.Valid {
		return t.Time
	}
	return time.Time{}
}

// From sets the underlying value of t to v. t becomes valid.
func (t *Time) From(v time.Time) {
	t.Valid = true
	t.Time = v
}

// FromPtr invalidates t if p is nil, otherwise it sets the underlying value
// of t to the value pointed to by p, and t becomes valid.
func (t *Time) FromPtr(v *time.Time) {
	t.Valid = v != nil
	if v != nil {
		t.Time = *v
	}
}

// FromZero invalidates t if v represents the zero time instant,
// otherwise it sets the underlying value of t to v, and t becomes valid.
func (t *Time) FromZero(v time.Time) {
	t.Valid = !v.IsZero()
	t.Time = v
}

// String returns a string representation of t. If t is valid,
// it formats the underlying value of t according to the RFC3339 standard with
// nanoseconds. For time instants which year is beyond 10000, not allowed by the
// standard, String tries to print a meaningful datetime regardless.
// If t is not valid, it returns InvalidNullableString.
func (t Time) String() string {
	if t.Valid {
		return t.Time.Format(time.RFC3339Nano)
	}
	return InvalidNullableString
}

// MarshalText marshals t to a byte string representation. If t is valid, it
// formats the underlying value of t according to the RFC3339 standard with
// nanoseconds, otherwise it returns nil. If the underlying value of t
// cannot be marshaled, a MarshalError is returned.
func (t Time) MarshalText() (data []byte, err error) {
	if t.Valid {
		bytes, err := t.Time.MarshalText()
		if err != nil {
			return nil, makeMarshalError("text", t)
		}
		return bytes, nil
	}
	return nil, nil
}

// MarshalJSON encodes the underlying value of t to a JSON string
// representation if t is valid, otherwise it returns the JSON null value.
// The string representation is formatted according to the RFC3339 standard
// with nanoseconds. If the underlying value of t cannot be marshaled,
// a MarshalError is returned.
func (t Time) MarshalJSON() (data []byte, err error) {
	if t.Valid {
		bytes, err := t.Time.MarshalJSON()
		if err != nil {
			return nil, makeMarshalError("json", t)
		}
		return bytes, nil
	}
	return jNull, nil
}

// Value returns the underlying value of t if t is valid,
// otherwise nil. err is always nil.
func (t Time) Value() (v driver.Value, err error) {
	if t.Valid {
		return t.Time, nil
	}
	return nil, nil
}

// Set invalidates t if str is the empty string, otherwise it parses str into
// the underlying value of t, and t becomes valid. If str is not a valid
// RFC3339 string, t becomes invalid and a ParseError is returned.
func (t *Time) Set(str string) error {
	var err error
	t.Time, err = time.Parse(time.RFC3339Nano, str)
	t.Valid = err == nil && str != ""
	if str == "" || err == nil {
		return nil
	}

	return makeParseError("parse", str, t.Time)
}

// UnmarshalText unmarshals from a byte string to t.
// It behaves like Set, except that it returns an UnmarshalError instead of a
// ParseError in case text cannot be parsed.
func (t *Time) UnmarshalText(text []byte) error {
	err := t.Time.UnmarshalText(text)
	t.Valid = err == nil && len(text) > 0

	if err == nil || len(text) == 0 {
		return nil
	}
	return makeUnmarshalError("text", text, *t)
}

// UnmarshalJSON unmarshals from a JSON encoded byte string to t.
// If the encoded JSON data represent the JSON null value,
// or an error is produced, t becomes invalid. If the encoded JSON data
// represent a valid RFC3339 JSON string, t becomes valid, and the underlying
// value of t is set to the JSON string. If the encoded JSON data represent
// a JSON string, but not formatted according to the RFC3339 standard,
// a ParseError is returned. Other JSON types produce a TypeError.
// Malformed JSON produces an UnmarshalError.
func (t *Time) UnmarshalJSON(data []byte) error {
	var obj interface{}
	var err error
	if json.Unmarshal(data, &obj) != nil {
		t.Valid = false
		return makeUnmarshalError("json", data, *t)
	}
	switch value := obj.(type) {
	case string:
		t.Time, err = time.Parse(time.RFC3339Nano, value)
		t.Valid = err == nil
		if t.Valid {
			return nil
		}
		return makeParseError("parse", value, t.Time)
	case nil:
		t.Valid = false
		return nil
	default:
		t.Valid = false
		return makeTypeError("json", value, "string", "nil")
	}
}

// Scan assigns a value from a database driver. If obj's type is time.Time,
// t becomes valid, and the underlying value of t becomes the value of obj.
// If obj is nil, t becomes invalid. If obj's type is any other type,
// t becomes invalid, and a TypeError is returned.
func (t *Time) Scan(obj interface{}) error {
	switch value := obj.(type) {
	case time.Time:
		t.Time = value
		t.Valid = true
		return nil
	case nil:
		t.Valid = false
		return nil
	default:
		t.Valid = false
		return makeTypeError("sql", value, "time.Time", "nil")
	}
}
