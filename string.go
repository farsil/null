package null

import (
	"database/sql/driver"
	"encoding/json"
)

type String struct {
	// the inner value can't be named String because it conflicts with the
	// flag.Value interface implementation.
	Str   string
	Valid bool
}

func StringFrom(v string) String {
	return StringFromPtr(&v)
}

func StringFromPtr(v *string) String {
	if v != nil {
		return String{
			Str:   *v,
			Valid: true,
		}
	}
	return String{}
}

func StringFromZero(v string) String {
	return String{
		Str:   v,
		Valid: v != "",
	}
}

func (s String) Ptr() *string {
	if s.Valid {
		return &s.Str
	}
	return nil
}

func (s String) Zero() string {
	if s.Valid {
		return s.Str
	}
	return ""
}

func (s *String) From(v string) {
	s.Valid = true
	s.Str = v
}

func (s *String) FromPtr(v *string) {
	s.Valid = v != nil
	if v != nil {
		s.Str = *v
	}
}

func (s *String) FromZero(v string) {
	s.Valid = v != ""
	s.Str = v
}

func (s String) String() string {
	if s.Valid {
		return s.Str
	}
	return sInvalid
}

func (s String) MarshalText() ([]byte, error) {
	if s.Valid {
		return []byte(s.Str), nil
	}
	return nil, nil
}

func (s String) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.Str)
	}
	return jNull, nil
}

func (s String) Value() (driver.Value, error) {
	if s.Valid {
		return s.Str, nil
	}
	return nil, nil
}

func (s *String) Set(str string) error {
	s.FromZero(str)
	return nil
}

func (s *String) UnmarshalText(text []byte) error {
	s.Str = string(text)
	s.Valid = text != nil
	return nil
}

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
		return makeTypeError("json", value, "string", "nil")
	}
}

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
		return makeTypeError("sql", value, "string", "nil")
	}
}
