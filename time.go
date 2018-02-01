package null

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Time struct {
	Time  time.Time
	Valid bool
}

func TimeFrom(v time.Time) Time {
	return TimeFromPtr(&v)
}

func TimeFromPtr(v *time.Time) Time {
	if v != nil {
		return Time{
			Time:  *v,
			Valid: true,
		}
	}
	return Time{}
}

func TimeFromZero(v time.Time) Time {
	return Time{
		Time:  v,
		Valid: v != tEmpty,
	}
}

func (t Time) Ptr() *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

func (t Time) Zero() time.Time {
	if t.Valid {
		return t.Time
	}
	return tEmpty
}

func (t *Time) From(v time.Time) {
	t.Valid = true
	t.Time = v
}

func (t *Time) FromPtr(v *time.Time) {
	t.Valid = v != nil
	if v != nil {
		t.Time = *v
	}
}

func (t *Time) FromZero(v time.Time) {
	t.Valid = v != tEmpty
	t.Time = v
}

// RFC3339 does not allow year > 10000, but this function will print a
// meaningful date regardless
func (t Time) String() string {
	if t.Valid {
		return t.Time.Format(time.RFC3339Nano)
	}
	return sInvalid
}

func (t Time) MarshalText() ([]byte, error) {
	if t.Valid {
		bytes, err := t.Time.MarshalText()
		if err != nil {
			return nil, makeMarshalError("text", t)
		}
		return bytes, nil
	}
	return nil, nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	if t.Valid {
		bytes, err := t.Time.MarshalJSON()
		if err != nil {
			return nil, makeMarshalError("json", t)
		}
		return bytes, nil
	}
	return jNull, nil
}

func (t Time) Value() (driver.Value, error) {
	if t.Valid {
		return t.Time, nil
	}
	return nil, nil
}

func (t *Time) Set(str string) error {
	var err error
	t.Time, err = time.Parse(time.RFC3339Nano, str)
	t.Valid = err == nil && str != ""
	if str == "" || err == nil {
		return nil
	}

	return makeParseError("parse", str, t.Time)
}

func (t *Time) UnmarshalText(text []byte) error {
	err := t.Time.UnmarshalText(text)
	t.Valid = err == nil && len(text) > 0

	if err == nil || len(text) == 0 {
		return nil
	}
	return makeUnmarshalError("text", text, *t)
}

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
		return makeTypeError("json", value, "string", "nil")
	}
}

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
		return makeTypeError("sql", value, "time.Time", "nil")
	}
}
