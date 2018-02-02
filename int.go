package null

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

type Int struct {
	Int   int
	Valid bool
}

func IntFrom(v int) Int {
	return IntFromPtr(&v)
}

func IntFromPtr(v *int) Int {
	if v != nil {
		return Int{
			Int:   *v,
			Valid: true,
		}
	}
	return Int{}
}

func IntFromZero(v int) Int {
	return Int{
		Int:   v,
		Valid: v != 0,
	}
}

func (i Int) Ptr() *int {
	if i.Valid {
		return &i.Int
	}
	return nil
}

func (i Int) Zero() int {
	if i.Valid {
		return i.Int
	}
	return 0
}

func (i *Int) From(v int) {
	i.Valid = true
	i.Int = v
}

func (i *Int) FromPtr(v *int) {
	i.Valid = v != nil
	if v != nil {
		i.Int = *v
	}
}

func (i *Int) FromZero(v int) {
	i.Valid = v != 0
	i.Int = v
}

func (i Int) String() string {
	if i.Valid {
		return strconv.FormatInt(int64(i.Int), 10)
	}
	return InvalidNullableString
}

func (i Int) MarshalText() ([]byte, error) {
	if i.Valid {
		return []byte(i.String()), nil
	}
	return nil, nil
}

func (i Int) MarshalJSON() ([]byte, error) {
	if i.Valid {
		bytes, err := json.Marshal(i.Int)
		if err != nil {
			return nil, makeMarshalError("json", i)
		}
		return bytes, nil
	}
	return jNull, nil
}

func (i Int) Value() (driver.Value, error) {
	if i.Valid {
		return int64(i.Int), nil
	}
	return nil, nil
}

func (i *Int) Set(str string) error {
	val, err := strconv.ParseInt(str, 0, intSize)
	i.Int = int(val)
	i.Valid = err == nil && str != ""
	if str == "" || err == nil {
		return nil
	}

	return makeParseError("parse", str, i.Int)
}

func (i *Int) UnmarshalText(text []byte) error {
	if i.Set(string(text)) != nil {
		return makeUnmarshalError("parse", text, *i)
	}
	return nil
}

func (i *Int) UnmarshalJSON(data []byte) error {
	var obj interface{}
	if json.Unmarshal(data, &obj) != nil {
		return makeUnmarshalError("json", data, *i)
	}
	switch value := obj.(type) {
	case float64:
		i.Int = int(value)
		i.Valid = value == float64(i.Int)
		if !i.Valid {
			return makeConversionError("json", value, i.Int)
		}
		return nil
	case nil:
		i.Valid = false
		return nil
	default:
		return makeTypeError("json", value, "float64", "nil")
	}
}

func (i *Int) Scan(obj interface{}) error {
	switch value := obj.(type) {
	case int64:
		i.Int = int(value)
		i.Valid = value == int64(i.Int)
		if !i.Valid {
			return makeConversionError("sql", value, i.Int)
		}
		return nil
	case nil:
		i.Valid = false
		return nil
	default:
		return makeTypeError("sql", value, "int64", "nil")
	}
}
