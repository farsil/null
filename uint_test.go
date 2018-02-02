package null_test

import (
	"null"
	"reflect"
	"strconv"
	"testing"
)

func uintp(v uint) *uint {
	return &v
}

func TestUintFrom(t *testing.T) {
	cases := []struct {
		literal uint
		valid   bool
	}{
		{0, true},
		{1, true},
	}

	for n, c := range cases {
		u := null.UintFrom(c.literal)
		if c.valid != u.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, u.Valid,
			)
		}
		if !u.Valid {
			continue
		}

		if c.literal != u.Uint {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %d, got %d)",
				t.Name(), n+1, c.literal, u.Uint,
			)
		}
	}
}

func TestUintFromPtr(t *testing.T) {
	cases := []struct {
		ptr   *uint
		valid bool
	}{
		{uintp(0), true},
		{uintp(1), true},
		{nil, false},
	}

	for n, c := range cases {
		u := null.UintFromPtr(c.ptr)
		if c.valid != u.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, u.Valid,
			)
		}
		if !u.Valid {
			continue
		}

		if *c.ptr != u.Uint {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %d, got %d)",
				t.Name(), n+1, *c.ptr, u.Uint,
			)
		}
	}
}

func TestUintFromZero(t *testing.T) {
	cases := []struct {
		literal uint
		valid   bool
	}{
		{0, false},
		{1, true},
	}

	for n, c := range cases {
		u := null.UintFromZero(c.literal)
		if c.valid != u.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, u.Valid,
			)
		}
		if !u.Valid {
			continue
		}

		if c.literal != u.Uint {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %d, got %d)",
				t.Name(), n+1, c.literal, u.Uint,
			)
		}
	}
}

func TestUint_Ptr(t *testing.T) {
	cases := []struct {
		nullable null.Uint
	}{
		{null.Uint{Uint: 0, Valid: true}},
		{null.Uint{Uint: 1, Valid: true}},
		{null.Uint{}},
		{null.Uint{Uint: 2, Valid: false}},
	}

	for n, c := range cases {
		p := c.nullable.Ptr()
		if c.nullable.Valid {
			if p == nil {
				t.Fatalf(
					"%s, case #%d: nil pointer returned",
					t.Name(), n+1,
				)
			}
			if c.nullable.Uint != *p {
				t.Fatalf(
					"%s, case #%d: literal mismatch (expected %d, got %d)",
					t.Name(), n+1, c.nullable.Uint, *p,
				)
			}
		} else {
			if p != nil {
				t.Fatalf(
					"%s, case #%d: non-nil pointer returned",
					t.Name(), n+1,
				)
			}
		}
	}
}

func TestUint_Zero(t *testing.T) {
	cases := []struct {
		nullable null.Uint
	}{
		{null.Uint{Uint: 0, Valid: true}},
		{null.Uint{Uint: 1, Valid: true}},
		{null.Uint{}},
		{null.Uint{Uint: 2, Valid: false}},
	}

	for n, c := range cases {
		u := c.nullable.Zero()
		if c.nullable.Valid {
			if c.nullable.Uint != u {
				t.Fatalf(
					"%s, case #%d: literal mismatch (expected %d, got %d)",
					t.Name(), n+1, c.nullable.Uint, u,
				)
			}
		} else {
			if u != 0 {
				t.Fatalf(
					"%s, case #%d: non-zero uint returned",
					t.Name(), n+1,
				)
			}
		}
	}
}

func TestUint_From(t *testing.T) {
	var u null.Uint
	cases := []struct {
		literal uint
		valid   bool
	}{
		{0, true},
		{1, true},
	}

	for n, c := range cases {
		u.From(c.literal)
		if c.valid != u.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, u.Valid,
			)
		}
		if !u.Valid {
			continue
		}

		if c.literal != u.Uint {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %d, got %d)",
				t.Name(), n+1, c.literal, u.Uint,
			)
		}
	}
}

func TestUint_FromPtr(t *testing.T) {
	var u null.Uint
	cases := []struct {
		ptr   *uint
		valid bool
	}{
		{uintp(0), true},
		{uintp(1), true},
		{nil, false},
	}

	for n, c := range cases {
		u.FromPtr(c.ptr)
		if c.valid != u.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, u.Valid,
			)
		}
		if !u.Valid {
			continue
		}

		if *c.ptr != u.Uint {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %d, got %d)",
				t.Name(), n+1, *c.ptr, u.Uint,
			)
		}
	}
}

func TestUint_FromZero(t *testing.T) {
	var u null.Uint
	cases := []struct {
		literal uint
		valid   bool
	}{
		{0, false},
		{1, true},
	}

	for n, c := range cases {
		u.FromZero(c.literal)
		if c.valid != u.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, u.Valid,
			)
		}
		if !u.Valid {
			continue
		}

		if c.literal != u.Uint {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %d, got %d)",
				t.Name(), n+1, c.literal, u.Uint,
			)
		}
	}
}

func TestUint_String(t *testing.T) {
	cases := []struct {
		nullable null.Uint
		string   string
	}{
		{null.Uint{Uint: 0, Valid: true}, "0"},
		{null.Uint{Uint: 1, Valid: true}, "1"},
		{null.Uint{}, "<invalid>"},
		{null.Uint{Uint: 2, Valid: false}, "<invalid>"},
	}

	for n, c := range cases {
		if c.string != c.nullable.String() {
			t.Fatalf(
				"%s, case #%d: string mismatch (expected '%s', got '%s')",
				t.Name(), n+1, c.string, c.nullable.String(),
			)
		}
	}
}

func TestUint_MarshalText(t *testing.T) {
	nilType := reflect.TypeOf(nil)

	cases := []struct {
		nullable null.Uint
		bytes    []byte
		errType  reflect.Type
	}{
		{null.Uint{Uint: 0, Valid: true}, []byte("0"), nilType},
		{null.Uint{Uint: 1, Valid: true}, []byte("1"), nilType},
		{null.Uint{}, nil, nilType},
		{null.Uint{Uint: 2, Valid: false}, nil, nilType},
	}

	for n, c := range cases {
		u, err := c.nullable.MarshalText()
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if !reflect.DeepEqual(c.bytes, u) {
			t.Fatalf(
				"%s, case #%d: bytes mismatch (expected '%s', got '%s')",
				t.Name(), n+1, string(c.bytes), string(u),
			)
		}
	}
}

func TestUint_MarshalJSON(t *testing.T) {
	nilType := reflect.TypeOf(nil)

	cases := []struct {
		nullable null.Uint
		json     []byte
		errType  reflect.Type
	}{
		{null.Uint{Uint: 0, Valid: true}, []byte("0"), nilType},
		{null.Uint{Uint: 1, Valid: true}, []byte("1"), nilType},
		{null.Uint{}, []byte("null"), nilType},
		{null.Uint{Uint: 2, Valid: false}, []byte("null"), nilType},
	}

	for n, c := range cases {
		u, err := c.nullable.MarshalJSON()
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if !reflect.DeepEqual(c.json, u) {
			t.Fatalf(
				"%s, case #%d: json mismatch (expected '%s', got '%s')",
				t.Name(), n+1, string(c.json), string(u),
			)
		}
	}
}

func TestUint_Value64(t *testing.T) {
	u := null.Uint{
		Uint:  (^uint(0)) ^ (^uint(0) >> 1),
		Valid: true,
	}
	val64 := uint64(1 << 63)
	errType := reflect.TypeOf(null.ConversionError{})

	if val64 != uint64(u.Uint) {
		t.Skipf("%s: uint is not 64-bit, skipping", t.Name())
	}

	_, err := u.Value()
	if errType != reflect.TypeOf(err) {
		t.Fatalf(
			"%s: wrong error type (expected %v, got %v)",
			t.Name(), errType, reflect.TypeOf(err),
		)
	}
}

func TestUint_Value(t *testing.T) {
	nilType := reflect.TypeOf(nil)
	i64Type := reflect.TypeOf(int64(0))

	cases := []struct {
		nullable null.Uint
		expType  reflect.Type
		errType  reflect.Type
	}{
		{null.Uint{Uint: 0, Valid: true}, i64Type, nilType},
		{null.Uint{Uint: 1, Valid: true}, i64Type, nilType},
		{null.Uint{}, nilType, nilType},
		{null.Uint{Uint: 2, Valid: false}, nilType, nilType},
	}

	for n, c := range cases {
		v, err := c.nullable.Value()
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.expType != reflect.TypeOf(v) {
			t.Fatalf(
				"%s, case #%d: type mismatch (expected %s, got %s)",
				t.Name(), n+1, c.expType, reflect.TypeOf(v),
			)
		}
		if v == nil {
			continue
		}

		if int64(c.nullable.Uint) != reflect.ValueOf(v).Int() {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %d, got %v)",
				t.Name(), n+1, c.nullable.Uint, v,
			)
		}
	}
}

func TestUint_Set32(t *testing.T) {
	var u null.Uint
	val := (^uint(0)) ^ (^uint(0) >> 1)
	val64 := uint64(1 << 63)
	errType := reflect.TypeOf(null.ParseError{})

	if val64 == uint64(val) {
		t.Skipf("%s: uint is not 32-bit, skipping", t.Name())
	}

	err := u.Set(strconv.FormatUint(val64, 10))
	if errType != reflect.TypeOf(err) {
		t.Fatalf(
			"%s: wrong error type (expected %v, got %v)",
			t.Name(), errType, reflect.TypeOf(err),
		)
	}
}

func TestUint_Set(t *testing.T) {
	var u null.Uint
	nilType := reflect.TypeOf(nil)
	parseErrType := reflect.TypeOf(null.ParseError{})

	cases := []struct {
		string  string
		literal uint
		valid   bool
		errType reflect.Type
	}{
		{"0", 0, true, nilType},
		{"1", 1, true, nilType},
		{"010", 8, true, nilType},
		{"0x10", 16, true, nilType},
		{"", 0, false, nilType},
		{"-1", 0, false, parseErrType},
		{"0.1", 0, false, parseErrType},
		{"x", 0, false, parseErrType},
	}

	for n, c := range cases {
		err := u.Set(c.string)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != u.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, u.Valid,
			)
		}
		if !u.Valid {
			continue
		}

		if c.literal != u.Uint {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %d, got %d)",
				t.Name(), n+1, c.literal, u.Uint,
			)
		}
	}
}

func TestUint_UnmarshalText32(t *testing.T) {
	var u null.Uint
	val := (^uint(0)) ^ (^uint(0) >> 1)
	val64 := uint64(1 << 63)
	errType := reflect.TypeOf(null.UnmarshalError{})

	if val64 == uint64(val) {
		t.Skipf("%s: uint is not 32-bit, skipping", t.Name())
	}

	err := u.UnmarshalText([]byte(strconv.FormatUint(val64, 10)))
	if errType != reflect.TypeOf(err) {
		t.Fatalf(
			"%s: wrong error type (expected %v, got %v)",
			t.Name(), errType, reflect.TypeOf(err),
		)
	}
}

func TestUint_UnmarshalText(t *testing.T) {
	var u null.Uint
	nilType := reflect.TypeOf(nil)
	unmarshalErrType := reflect.TypeOf(null.UnmarshalError{})

	cases := []struct {
		bytes   []byte
		literal uint
		valid   bool
		errType reflect.Type
	}{
		{[]byte("0"), 0, true, nilType},
		{[]byte("1"), 1, true, nilType},
		{[]byte("010"), 8, true, nilType},
		{[]byte("0x10"), 16, true, nilType},
		{nil, 0, false, nilType},
		{[]byte(""), 0, false, nilType},
		{[]byte("-1"), 0, false, unmarshalErrType},
		{[]byte("0.1"), 0, false, unmarshalErrType},
		{[]byte("x"), 0, false, unmarshalErrType},
	}

	for n, c := range cases {
		err := u.UnmarshalText(c.bytes)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != u.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, u.Valid,
			)
		}
		if !u.Valid {
			continue
		}

		if c.literal != u.Uint {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %d, got %d)",
				t.Name(), n+1, c.literal, u.Uint,
			)
		}
	}
}

func TestUint_UnmarshalJSON32(t *testing.T) {
	var u null.Uint
	val := (^uint(0)) ^ (^uint(0) >> 1)
	val64 := uint64(1 << 63)
	errType := reflect.TypeOf(null.ConversionError{})

	if val64 == uint64(val) {
		t.Skipf("%s: uint is not 32-bit, skipping", t.Name())
	}

	err := u.UnmarshalJSON([]byte(strconv.FormatUint(val64, 10)))
	if errType != reflect.TypeOf(err) {
		t.Fatalf(
			"%s: wrong error type (expected %v, got %v)",
			t.Name(), errType, reflect.TypeOf(err),
		)
	}
}

func TestUint_UnmarshalJSON(t *testing.T) {
	var u null.Uint
	nilType := reflect.TypeOf(nil)
	cnvErrType := reflect.TypeOf(null.ConversionError{})
	typeErrType := reflect.TypeOf(null.TypeError{})
	unmarshalErrType := reflect.TypeOf(null.UnmarshalError{})

	cases := []struct {
		json    []byte
		literal uint
		valid   bool
		errType reflect.Type
	}{
		{[]byte("0"), 0, true, nilType},
		{[]byte("1"), 1, true, nilType},
		{[]byte("null"), 0, false, nilType},
		{[]byte("0.1"), 0, false, cnvErrType},
		{[]byte("-1"), 0, false, cnvErrType},
		{[]byte(`"x"`), 0, false, typeErrType},
		{nil, 0, false, unmarshalErrType},
		{[]byte("x"), 0, false, unmarshalErrType},
	}

	for n, c := range cases {
		err := u.UnmarshalJSON(c.json)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != u.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, u.Valid,
			)
		}
		if !u.Valid {
			continue
		}

		if c.literal != u.Uint {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %d, got %d)",
				t.Name(), n+1, c.literal, u.Uint,
			)
		}
	}
}

func TestUint_Scan32(t *testing.T) {
	var u null.Uint
	errType := reflect.TypeOf(null.ConversionError{})
	val := ((^uint(0)) ^ (^uint(0) >> 1)) >> 1
	val64 := int64(1 << 62)

	if val64 == int64(val) {
		t.Skipf("%s: uint is not 32-bit, skipping", t.Name())
	}

	err := u.Scan(val64)
	if errType != reflect.TypeOf(err) {
		t.Fatalf(
			"%s: wrong error type (expected %v, got %v)",
			t.Name(), errType, reflect.TypeOf(err),
		)
	}
}

func TestUint_Scan(t *testing.T) {
	var u null.Uint
	nilType := reflect.TypeOf(nil)
	conversionErrType := reflect.TypeOf(null.ConversionError{})
	typeErrType := reflect.TypeOf(null.TypeError{})

	cases := []struct {
		source  interface{}
		valid   bool
		errType reflect.Type
	}{
		{int64(0), true, nilType},
		{int64(1), true, nilType},
		{nil, false, nilType},
		{int64(-1), false, conversionErrType},
		{float64(0.1), false, typeErrType},
		{"x", false, typeErrType},
	}

	for n, c := range cases {
		err := u.Scan(c.source)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != u.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, u.Valid,
			)
		}
		if !u.Valid {
			continue
		}

		if c.source.(int64) != int64(u.Uint) {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %d, got %d)",
				t.Name(), n+1, c.source, u.Uint,
			)
		}
	}
}
