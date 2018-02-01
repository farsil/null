package null

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// Provides a nullable time.Time.
type Time struct {
	// The underlying time.Time value.
	Time time.Time

	// If true, the underlying time.Time value is valid. If false,
	// the value stored in Time is invalid, and thus meaningless.
	Valid bool
}

// Creates a valid Time from v.
func TimeFrom(v time.Time) Time {
	return TimeFromPtr(&v)
}

// Creates a Time from pointer p. If p is nil, the returned Time is invalid.
func TimeFromPtr(v *time.Time) Time {
	if v != nil {
		return Time{
			Time:  *v,
			Valid: true,
		}
	}
	return Time{}
}

// Creates a Time from v. If v is a zero time.Time,
// the returned Time is invalid.
func TimeFromZero(v time.Time) Time {
	return Time{
		Time:  v,
		Valid: !v.IsZero(),
	}
}

// Returns a pointer to the underlying time.Time if Time is valid, otherwise
// returns nil.
func (t Time) Ptr() *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

// Returns the underlying time.Time if Bool is valid, otherwise
// returns a zero time.Time.
func (t Time) Zero() time.Time {
	if t.Valid {
		return t.Time
	}
	return time.Time{}
}

// Sets the underlying time.Time to v. Time becomes valid.
func (t *Time) From(v time.Time) {
	t.Valid = true
	t.Time = v
}

// If p is nil, Time becomes invalid, otherwise it sets the underlying time.Time
// to the value pointed to by p, and Time becomes valid.
func (t *Time) FromPtr(v *time.Time) {
	t.Valid = v != nil
	if v != nil {
		t.Time = *v
	}
}

// If v is a zero time.Time, Time becomes invalid,
// otherwise it sets the underlying time.Time to v, and Bool becomes valid.
func (t *Time) FromZero(v time.Time) {
	t.Valid = !v.IsZero()
	t.Time = v
}

// Returns a string representation of Time if valid, and InvalidNullableStr
// if not valid. Valid Time objects are formatted following the RFC3339 standard
// with nanoseconds. For dates which year is beyond 10000, not allowed by the
// standard, String tries to print a meaningful datetime regardless.
func (t Time) String() string {
	if t.Valid {
		return t.Time.Format(time.RFC3339Nano)
	}
	return InvalidNullableStr
}

// Marshals the underlying value to a byte string representation if Time is
// valid, otherwise it marshals to nil. The string representation is formatted
// following the RFC3339 standard with nanoseconds. If the underlying value
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

// Marshals to a JSON string representation if Time is valid,
// otherwise it marshals to the JSON null value. The string representation
// is formatted following the RFC3339 standard with nanoseconds.
// If the underlying value cannot be marshaled, a MarshalError is returned.
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

// Returns the underlying time.Time if Time is valid,
// otherwise nil. err is always nil.
func (t Time) Value() (v driver.Value, err error) {
	if t.Valid {
		return t.Time, nil
	}
	return nil, nil
}

// Parses a string representation of a time.Time. If str is an empty string,
// Time becomes invalid. If str is a valid RFC3339 string, Time becomes valid,
// str is parsed, and the parsed value becomes the underlying value.
// If str is an invalid RFC3339 string, Time becomes invalid, and a ParseError
// is returned.
func (t *Time) Set(str string) error {
	var err error
	t.Time, err = time.Parse(time.RFC3339Nano, str)
	t.Valid = err == nil && str != ""
	if str == "" || err == nil {
		return nil
	}

	return makeParseError("parse", str, t.Time)
}

// Unmarshals from a byte string representation of a time.Time.
// It behaves like Set,
// except that it returns an UnmarshalError instead of a ParseError if
// text is an invalid RFC3339 string.
func (t *Time) UnmarshalText(text []byte) error {
	err := t.Time.UnmarshalText(text)
	t.Valid = err == nil && len(text) > 0

	if err == nil || len(text) == 0 {
		return nil
	}
	return makeUnmarshalError("text", text, *t)
}

// Unmarshals from a JSON object. If the JSON object is the JSON null value,
// or an error is produced, Time becomes invalid.
// If the JSON object is a JSON string, UnmarshalJSON attempts to parse the
// object. If the JSON string is a valid RFC3339 datetime,
// the underlying value is set to the parsed value and Time becomes valid,
// otherwise Time becomes invalid and a ParseError is returned. Other JSON types
// produce a TypeError. Malformed JSON produces an UnmarshalError.
func (t *Time) UnmarshalJSON(data []byte) error {
	var obj interface{}
	var err error
	if json.Unmarshal(data, &obj) != nil {
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

// Assigns a value from a database driver. If obj's type is time.Time,
// Time becomes valid, and the underlying time.Time becomes the value of obj.
// If obj is nil, Time becomes invalid. If obj's type is any other type,
// Time becomes invalid, and a TypeError is returned.
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
