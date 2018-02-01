package null_test

import (
	"null"
	"reflect"
	"strconv"
	"testing"
)

func intp(v int) *int {
	return &v
}

func TestIntFrom(t *testing.T) {
	cases := []struct {
		literal int
		valid   bool
	}{
		{0, true},
		{1, true},
		{-1, true},
	}

	for n, c := range cases {
		i := null.IntFrom(c.literal)
		if c.valid != i.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, i.Valid,
			)
		}
		if !i.Valid {
			continue
		}

		if c.literal != i.Int {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %d, got %d)",
				t.Name(), n+1, c.literal, i.Int,
			)
		}
	}
}

func TestIntFromPtr(t *testing.T) {
	cases := []struct {
		ptr   *int
		valid bool
	}{
		{intp(0), true},
		{intp(1), true},
		{intp(-1), true},
		{nil, false},
	}

	for n, c := range cases {
		i := null.IntFromPtr(c.ptr)
		if c.valid != i.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, i.Valid,
			)
		}
		if !i.Valid {
			continue
		}

		if *c.ptr != i.Int {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %d, got %d)",
				t.Name(), n+1, *c.ptr, i.Int,
			)
		}
	}
}

func TestIntFromZero(t *testing.T) {
	cases := []struct {
		literal int
		valid   bool
	}{
		{0, false},
		{1, true},
		{-1, true},
	}

	for n, c := range cases {
		i := null.IntFromZero(c.literal)
		if c.valid != i.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, i.Valid,
			)
		}
		if !i.Valid {
			continue
		}

		if c.literal != i.Int {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %d, got %d)",
				t.Name(), n+1, c.literal, i.Int,
			)
		}
	}
}

func TestInt_Ptr(t *testing.T) {
	cases := []struct {
		nullable null.Int
	}{
		{null.Int{Int: 0, Valid: true}},
		{null.Int{Int: 1, Valid: true}},
		{null.Int{Int: -1, Valid: true}},
		{null.Int{}},
		{null.Int{Int: 2, Valid: false}},
		{null.Int{Int: -2, Valid: false}},
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
			if c.nullable.Int != *p {
				t.Fatalf(
					"%s, case #%d: literal mismatch (expected %d, got %d)",
					t.Name(), n+1, c.nullable.Int, *p,
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

func TestInt_Zero(t *testing.T) {
	cases := []struct {
		nullable null.Int
	}{
		{null.Int{Int: 0, Valid: true}},
		{null.Int{Int: 1, Valid: true}},
		{null.Int{Int: -1, Valid: true}},
		{null.Int{}},
		{null.Int{Int: 2, Valid: false}},
		{null.Int{Int: -2, Valid: false}},
	}

	for n, c := range cases {
		i := c.nullable.Zero()
		if c.nullable.Valid {
			if c.nullable.Int != i {
				t.Fatalf(
					"%s, case #%d: literal mismatch (expected %d, got %d)",
					t.Name(), n+1, c.nullable.Int, i,
				)
			}
		} else {
			if i != 0 {
				t.Fatalf(
					"%s, case #%d: non-zero uint returned",
					t.Name(), n+1,
				)
			}
		}
	}
}

func TestInt_From(t *testing.T) {
	var i null.Int
	cases := []struct {
		literal int
		valid   bool
	}{
		{0, true},
		{1, true},
		{-1, true},
	}

	for n, c := range cases {
		i.From(c.literal)
		if c.valid != i.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, i.Valid,
			)
		}
		if !i.Valid {
			continue
		}

		if c.literal != i.Int {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %d, got %d)",
				t.Name(), n+1, c.literal, i.Int,
			)
		}
	}
}

func TestInt_FromPtr(t *testing.T) {
	var i null.Int
	cases := []struct {
		ptr   *int
		valid bool
	}{
		{intp(0), true},
		{intp(1), true},
		{intp(-1), true},
		{nil, false},
	}

	for n, c := range cases {
		i.FromPtr(c.ptr)
		if c.valid != i.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, i.Valid,
			)
		}
		if !i.Valid {
			continue
		}

		if *c.ptr != i.Int {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %d, got %d)",
				t.Name(), n+1, *c.ptr, i.Int,
			)
		}
	}
}

func TestInt_FromZero(t *testing.T) {
	var i null.Int
	cases := []struct {
		literal int
		valid   bool
	}{
		{0, false},
		{1, true},
		{-1, true},
	}

	for n, c := range cases {
		i.FromZero(c.literal)
		if c.valid != i.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, i.Valid,
			)
		}
		if !i.Valid {
			continue
		}

		if c.literal != i.Int {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %d, got %d)",
				t.Name(), n+1, c.literal, i.Int,
			)
		}
	}
}

func TestInt_String(t *testing.T) {
	cases := []struct {
		nullable null.Int
		string   string
	}{
		{null.Int{Int: 0, Valid: true}, "0"},
		{null.Int{Int: 1, Valid: true}, "1"},
		{null.Int{Int: -1, Valid: true}, "-1"},
		{null.Int{}, "<invalid>"},
		{null.Int{Int: 2, Valid: false}, "<invalid>"},
		{null.Int{Int: -2, Valid: false}, "<invalid>"},
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

func TestInt_MarshalText(t *testing.T) {
	nilType := reflect.TypeOf(nil)

	cases := []struct {
		nullable null.Int
		bytes    []byte
		errType  reflect.Type
	}{
		{null.Int{Int: 0, Valid: true}, []byte("0"), nilType},
		{null.Int{Int: 1, Valid: true}, []byte("1"), nilType},
		{null.Int{Int: -1, Valid: true}, []byte("-1"), nilType},
		{null.Int{}, nil, nilType},
		{null.Int{Int: 2, Valid: false}, nil, nilType},
		{null.Int{Int: -2, Valid: false}, nil, nilType},
	}

	for n, c := range cases {
		i, err := c.nullable.MarshalText()
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if !reflect.DeepEqual(c.bytes, i) {
			t.Fatalf(
				"%s, case #%d: bytes mismatch (expected '%s', got '%s')",
				t.Name(), n+1, string(c.bytes), string(i),
			)
		}

	}
}

func TestInt_MarshalJSON(t *testing.T) {
	nilType := reflect.TypeOf(nil)

	cases := []struct {
		nullable null.Int
		json     []byte
		errType  reflect.Type
	}{
		{null.Int{Int: 0, Valid: true}, []byte("0"), nilType},
		{null.Int{Int: 1, Valid: true}, []byte("1"), nilType},
		{null.Int{Int: -1, Valid: true}, []byte("-1"), nilType},
		{null.Int{}, []byte("null"), nilType},
		{null.Int{Int: 2, Valid: false}, []byte("null"), nilType},
		{null.Int{Int: -2, Valid: false}, []byte("null"), nilType},
	}

	for n, c := range cases {
		i, err := c.nullable.MarshalJSON()
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if !reflect.DeepEqual(c.json, i) {
			t.Fatalf(
				"%s, case #%d: json mismatch (expected '%s', got '%s')",
				t.Name(), n+1, string(c.json), string(i),
			)
		}
	}
}

func TestInt_Value(t *testing.T) {
	nilType := reflect.TypeOf(nil)
	i64Type := reflect.TypeOf(int64(0))

	cases := []struct {
		nullable null.Int
		expType  reflect.Type
		errType  reflect.Type
	}{
		{null.Int{Int: 0, Valid: true}, i64Type, nilType},
		{null.Int{Int: 1, Valid: true}, i64Type, nilType},
		{null.Int{Int: -1, Valid: true}, i64Type, nilType},
		{null.Int{}, nilType, nilType},
		{null.Int{Int: 2, Valid: false}, nilType, nilType},
		{null.Int{Int: -2, Valid: false}, nilType, nilType},
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

		if int64(c.nullable.Int) != reflect.ValueOf(v).Int() {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %d, got %v)",
				t.Name(), n+1, c.nullable.Int, v,
			)
		}
	}
}

func TestInt_Set32(t *testing.T) {
	var i null.Int
	val := ((^uint(0)) ^ (^uint(0) >> 1)) >> 1
	val64 := int64(1 << 62)
	errType := reflect.TypeOf(null.ParseError{})

	if val64 == int64(val) {
		t.Skipf("%s: int is not 32-bit, skipping", t.Name())
	}

	err := i.Set(strconv.FormatInt(val64, 10))
	if errType != reflect.TypeOf(err) {
		t.Fatalf(
			"%s: wrong error type (expected %v, got %v)",
			t.Name(), errType, reflect.TypeOf(err),
		)
	}
}

func TestInt_Set(t *testing.T) {
	var i null.Int
	nilType := reflect.TypeOf(nil)
	parseErrType := reflect.TypeOf(null.ParseError{})

	cases := []struct {
		string  string
		literal int
		valid   bool
		errType reflect.Type
	}{
		{"0", 0, true, nilType},
		{"1", 1, true, nilType},
		{"010", 8, true, nilType},
		{"0x10", 16, true, nilType},
		{"-1", -1, true, nilType},
		{"-010", -8, true, nilType},
		{"-0x10", -16, true, nilType},
		{"", 0, false, nilType},
		{"0.1", 0, false, parseErrType},
		{"x", 0, false, parseErrType},
	}

	for n, c := range cases {
		err := i.Set(c.string)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != i.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, i.Valid,
			)
		}
		if !i.Valid {
			continue
		}

		if c.literal != i.Int {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %d, got %d)",
				t.Name(), n+1, c.literal, i.Int,
			)
		}
	}
}

func TestInt_UnmarshalText32(t *testing.T) {
	var i null.Int
	val := ((^uint(0)) ^ (^uint(0) >> 1)) >> 1
	val64 := int64(1 << 62)
	errType := reflect.TypeOf(null.UnmarshalError{})

	if val64 == int64(val) {
		t.Skipf("%s: int is not 32-bit, skipping", t.Name())
	}

	err := i.UnmarshalText([]byte(strconv.FormatInt(val64, 10)))
	if errType != reflect.TypeOf(err) {
		t.Fatalf(
			"%s: wrong error type (expected %v, got %v)",
			t.Name(), errType, reflect.TypeOf(err),
		)
	}
}

func TestInt_UnmarshalText(t *testing.T) {
	var i null.Int
	nilType := reflect.TypeOf(nil)
	unmarshalErrType := reflect.TypeOf(null.UnmarshalError{})

	cases := []struct {
		bytes   []byte
		literal int
		valid   bool
		errType reflect.Type
	}{
		{[]byte("0"), 0, true, nilType},
		{[]byte("1"), 1, true, nilType},
		{[]byte("010"), 8, true, nilType},
		{[]byte("0x10"), 16, true, nilType},
		{[]byte("-1"), -1, true, nilType},
		{[]byte("-010"), -8, true, nilType},
		{[]byte("-0x10"), -16, true, nilType},
		{nil, 0, false, nilType},
		{[]byte(""), 0, false, nilType},
		{[]byte("0.1"), 0, false, unmarshalErrType},
		{[]byte("x"), 0, false, unmarshalErrType},
	}

	for n, c := range cases {
		err := i.UnmarshalText(c.bytes)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != i.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, i.Valid,
			)
		}
		if !i.Valid {
			continue
		}

		if c.literal != i.Int {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %d, got %d)",
				t.Name(), n+1, c.literal, i.Int,
			)
		}
	}
}

func TestInt_UnmarshalJSON32(t *testing.T) {
	var i null.Int
	val := ((^uint(0)) ^ (^uint(0) >> 1)) >> 1
	val64 := int64(1 << 62)
	errType := reflect.TypeOf(null.ConversionError{})

	if val64 == int64(val) {
		t.Skipf("%s: int is not 32-bit, skipping", t.Name())
	}

	err := i.UnmarshalJSON([]byte(strconv.FormatInt(val64, 10)))
	if errType != reflect.TypeOf(err) {
		t.Fatalf(
			"%s: wrong error type (expected %v, got %v)",
			t.Name(), errType, reflect.TypeOf(err),
		)
	}
}

func TestInt_UnmarshalJSON(t *testing.T) {
	var i null.Int
	nilType := reflect.TypeOf(nil)
	cnvErrType := reflect.TypeOf(null.ConversionError{})
	typeErrType := reflect.TypeOf(null.TypeError{})
	unmarshalErrType := reflect.TypeOf(null.UnmarshalError{})

	cases := []struct {
		json    []byte
		literal int
		valid   bool
		errType reflect.Type
	}{
		{[]byte("0"), 0, true, nilType},
		{[]byte("1"), 1, true, nilType},
		{[]byte("-1"), -1, true, nilType},
		{[]byte("null"), 0, false, nilType},
		{[]byte("0.1"), 0, false, cnvErrType},
		{[]byte(`"x"`), 0, false, typeErrType},
		{nil, 0, false, unmarshalErrType},
		{[]byte("x"), 0, false, unmarshalErrType},
	}

	for n, c := range cases {
		err := i.UnmarshalJSON(c.json)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != i.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, i.Valid,
			)
		}
		if !i.Valid {
			continue
		}

		if c.literal != i.Int {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %d, got %d)",
				t.Name(), n+1, c.literal, i.Int,
			)
		}
	}
}

func TestInt_Scan32(t *testing.T) {
	var i null.Int
	val := ((^uint(0)) ^ (^uint(0) >> 1)) >> 1
	val64 := int64(1 << 62)
	errType := reflect.TypeOf(null.ConversionError{})

	if val64 == int64(val) {
		t.Skipf("%s: int is not 32-bit, skipping", t.Name())
	}

	err := i.Scan(val64)
	if errType != reflect.TypeOf(err) {
		t.Fatalf(
			"%s: wrong error type (expected %v, got %v)",
			t.Name(), errType, reflect.TypeOf(err),
		)
	}
}

func TestInt_Scan(t *testing.T) {
	var i null.Int
	nilType := reflect.TypeOf(nil)
	typeErrType := reflect.TypeOf(null.TypeError{})

	cases := []struct {
		source  interface{}
		valid   bool
		errType reflect.Type
	}{
		{int64(0), true, nilType},
		{int64(1), true, nilType},
		{int64(-1), true, nilType},
		{nil, false, nilType},
		{float64(0.1), false, typeErrType},
		{"x", false, typeErrType},
	}

	for n, c := range cases {
		err := i.Scan(c.source)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != i.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, i.Valid,
			)
		}
		if !i.Valid {
			continue
		}

		if c.source.(int64) != int64(i.Int) {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %d, got %d)",
				t.Name(), n+1, c.source, i.Int,
			)
		}
	}
}
