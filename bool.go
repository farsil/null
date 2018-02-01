package null

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

type Bool struct {
	Bool  bool
	Valid bool
}

func BoolFrom(v bool) Bool {
	return BoolFromPtr(&v)
}

func BoolFromPtr(v *bool) Bool {
	if v != nil {
		return Bool{
			Bool:  *v,
			Valid: true,
		}
	}
	return Bool{}
}

func BoolFromZero(v bool) Bool {
	return Bool{
		Bool:  v,
		Valid: v == true,
	}
}

func (b Bool) Ptr() *bool {
	if b.Valid {
		return &b.Bool
	}
	return nil
}

func (b Bool) Zero() bool {
	if b.Valid {
		return b.Bool
	}
	return false
}

func (b *Bool) From(v bool) {
	b.Valid = true
	b.Bool = v
}

func (b *Bool) FromPtr(v *bool) {
	b.Valid = v != nil
	if v != nil {
		b.Bool = *v
	}
}

func (b *Bool) FromZero(v bool) {
	b.Valid = v == true
	b.Bool = v
}

func (b Bool) String() string {
	if b.Valid {
		return strconv.FormatBool(b.Bool)
	}
	return sInvalid
}

func (b Bool) MarshalText() ([]byte, error) {
	if b.Valid {
		return []byte(b.String()), nil
	}
	return nil, nil
}

func (b Bool) MarshalJSON() ([]byte, error) {
	if b.Valid {
		if b.Bool {
			return jTrue, nil
		} else {
			return jFalse, nil
		}
	}
	return jNull, nil
}

func (b Bool) Value() (driver.Value, error) {
	if b.Valid {
		return b.Bool, nil
	}
	return nil, nil
}

func (b *Bool) Set(str string) error {
	var err error
	b.Bool, err = strconv.ParseBool(str)
	b.Valid = err == nil && str != ""

	if err == nil || str == "" {
		return nil
	}
	return makeParseError("parse", str, b.Bool)
}

func (b *Bool) UnmarshalText(text []byte) error {
	if b.Set(string(text)) != nil {
		return makeUnmarshalError("text", text, *b)
	}
	return nil
}

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
		return makeTypeError("json", value, "bool", "nil")
	}
}

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
		return makeTypeError("sql", value, "bool", "nil")
	}
}
