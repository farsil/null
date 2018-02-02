package null

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

type Uint struct {
	Uint  uint
	Valid bool
}

func UintFrom(v uint) Uint {
	return UintFromPtr(&v)
}

func UintFromPtr(v *uint) Uint {
	if v != nil {
		return Uint{
			Uint:  *v,
			Valid: true,
		}
	}
	return Uint{}
}

func UintFromZero(v uint) Uint {
	return Uint{
		Uint:  v,
		Valid: v != 0,
	}
}

func (u Uint) Ptr() *uint {
	if u.Valid {
		return &u.Uint
	}
	return nil
}

func (u Uint) Zero() uint {
	if u.Valid {
		return u.Uint
	}
	return 0
}

func (u *Uint) From(v uint) {
	u.Valid = true
	u.Uint = v
}

func (u *Uint) FromPtr(v *uint) {
	u.Valid = v != nil
	if v != nil {
		u.Uint = *v
	}
}

func (u *Uint) FromZero(v uint) {
	u.Valid = v != 0
	u.Uint = v
}

func (u Uint) String() string {
	if u.Valid {
		return strconv.FormatUint(uint64(u.Uint), 10)
	}
	return InvalidNullableString
}

func (u Uint) MarshalText() ([]byte, error) {
	if u.Valid {
		return []byte(u.String()), nil
	}
	return nil, nil
}

func (u Uint) MarshalJSON() ([]byte, error) {
	if u.Valid {
		bytes, err := json.Marshal(u.Uint)
		if err != nil {
			return nil, makeMarshalError("json", u)
		}
		return bytes, nil
	}
	return jNull, nil
}

func (u Uint) Value() (driver.Value, error) {
	if u.Valid {
		return int64(u.Uint), nil
	}
	return nil, nil
}

func (u *Uint) Set(str string) error {
	val, err := strconv.ParseUint(str, 0, intSize)
	u.Uint = uint(val)
	u.Valid = err == nil && str != ""
	if str == "" || err == nil {
		return nil
	}

	return makeParseError("parse", str, u.Uint)
}

func (u *Uint) UnmarshalText(text []byte) error {
	if u.Set(string(text)) != nil {
		return makeUnmarshalError("parse", text, *u)
	}
	return nil
}

func (u *Uint) UnmarshalJSON(data []byte) error {
	var obj interface{}
	if json.Unmarshal(data, &obj) != nil {
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
		return makeTypeError("json", value, "float64", "nil")
	}
}

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
		return makeTypeError("sql", value, "int64", "nil")
	}
}
