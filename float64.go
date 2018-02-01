package null

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

type Float64 struct {
	Float64 float64
	Valid   bool
}

func Float64From(v float64) Float64 {
	return Float64FromPtr(&v)
}

func Float64FromPtr(v *float64) Float64 {
	if v != nil {
		return Float64{
			Float64: *v,
			Valid:   true,
		}
	}
	return Float64{}
}

func Float64FromZero(v float64) Float64 {
	return Float64{
		Float64: v,
		Valid:   v != 0.0,
	}
}

func (f Float64) Ptr() *float64 {
	if f.Valid {
		return &f.Float64
	}
	return nil
}

func (f Float64) Zero() float64 {
	if f.Valid {
		return f.Float64
	}
	return 0.0
}

func (f *Float64) From(v float64) {
	f.Valid = true
	f.Float64 = v
}

func (f *Float64) FromPtr(v *float64) {
	f.Valid = v != nil
	if v != nil {
		f.Float64 = *v
	}
}

func (f *Float64) FromZero(v float64) {
	f.Valid = v != 0.0
	f.Float64 = v
}

func (f Float64) String() string {
	if f.Valid {
		return strconv.FormatFloat(f.Float64, 'g', -1, 64)
	}
	return InvalidNullableStr
}

func (f Float64) MarshalText() ([]byte, error) {
	if f.Valid {
		return []byte(f.String()), nil
	}
	return nil, nil
}

func (f Float64) MarshalJSON() ([]byte, error) {
	if f.Valid {
		bytes, err := json.Marshal(f.Float64)
		if err != nil {
			return nil, makeMarshalError("json", f)
		}
		return bytes, nil
	}
	return jNull, nil
}

func (f Float64) Value() (driver.Value, error) {
	if f.Valid {
		return f.Float64, nil
	}
	return nil, nil
}

func (f *Float64) Set(str string) error {
	var err error
	f.Float64, err = strconv.ParseFloat(str, 64)
	f.Valid = err == nil && str != ""
	if str == "" || err == nil {
		return nil
	}

	return makeParseError("parse", str, f.Float64)
}

func (f *Float64) UnmarshalText(text []byte) error {
	if f.Set(string(text)) != nil {
		return makeUnmarshalError("parse", text, *f)
	}
	return nil
}

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
