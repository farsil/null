package null_test

import (
	"null"
	"reflect"
	"testing"
)

func boolp(v bool) *bool {
	return &v
}

func TestBoolFrom(t *testing.T) {
	cases := []struct {
		literal bool
		valid   bool
	}{
		{false, true},
		{true, true},
	}

	for n, c := range cases {
		b := null.BoolFrom(c.literal)
		if c.valid != b.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, b.Valid,
			)
		}
		if !b.Valid {
			continue
		}

		if c.literal != b.Bool {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %t, got %t)",
				t.Name(), n+1, c.literal, b.Bool,
			)
		}
	}
}

func TestBoolFromPtr(t *testing.T) {
	cases := []struct {
		ptr   *bool
		valid bool
	}{
		{boolp(false), true},
		{boolp(true), true},
		{nil, false},
	}

	for n, c := range cases {
		b := null.BoolFromPtr(c.ptr)
		if c.valid != b.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, b.Valid,
			)
		}
		if !b.Valid {
			continue
		}

		if *c.ptr != b.Bool {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %t, got %t)",
				t.Name(), n+1, *c.ptr, b.Bool,
			)
		}
	}
}

func TestBoolFromZero(t *testing.T) {
	cases := []struct {
		literal bool
		valid   bool
	}{
		{false, false},
		{true, true},
	}

	for n, c := range cases {
		b := null.BoolFromZero(c.literal)
		if c.valid != b.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, b.Valid,
			)
		}
		if !b.Valid {
			continue
		}

		if c.literal != b.Bool {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %t, got %t)",
				t.Name(), n+1, c.literal, b.Bool,
			)
		}
	}
}

func TestBool_Ptr(t *testing.T) {
	cases := []struct {
		nullable null.Bool
	}{
		{null.Bool{Bool: false, Valid: true}},
		{null.Bool{Bool: true, Valid: true}},
		{null.Bool{}},
		{null.Bool{Bool: true, Valid: false}},
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
			if c.nullable.Bool != *p {
				t.Fatalf(
					"%s, case #%d: literal mismatch (expected %t, got %t)",
					t.Name(), n+1, c.nullable.Bool, *p,
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

func TestBool_Zero(t *testing.T) {
	cases := []struct {
		nullable null.Bool
	}{
		{null.Bool{Bool: false, Valid: true}},
		{null.Bool{Bool: true, Valid: true}},
		{null.Bool{}},
		{null.Bool{Bool: true, Valid: false}},
	}

	for n, c := range cases {
		b := c.nullable.Zero()
		if c.nullable.Valid {
			if c.nullable.Bool != b {
				t.Fatalf(
					"%s, case #%d: literal mismatch (expected %t, got %t)",
					t.Name(), n+1, c.nullable.Bool, b,
				)
			}
		} else {
			if b != false {
				t.Fatalf(
					"%s, case #%d: true returned, expected false",
					t.Name(), n+1,
				)
			}
		}
	}
}

func TestBool_From(t *testing.T) {
	var b null.Bool
	cases := []struct {
		literal bool
		valid   bool
	}{
		{false, true},
		{true, true},
	}

	for n, c := range cases {
		b.From(c.literal)
		if c.valid != b.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, b.Valid,
			)
		}
		if !b.Valid {
			continue
		}

		if c.literal != b.Bool {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %t, got %t)",
				t.Name(), n+1, c.literal, b.Bool,
			)
		}
	}
}

func TestBool_FromPtr(t *testing.T) {
	var b null.Bool
	cases := []struct {
		ptr   *bool
		valid bool
	}{
		{boolp(false), true},
		{boolp(true), true},
		{nil, false},
	}

	for n, c := range cases {
		b.FromPtr(c.ptr)
		if c.valid != b.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, b.Valid,
			)
		}
		if !b.Valid {
			continue
		}

		if *c.ptr != b.Bool {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %t, got %t)",
				t.Name(), n+1, *c.ptr, b.Bool,
			)
		}
	}
}

func TestBool_FromZero(t *testing.T) {
	var b null.Bool
	cases := []struct {
		literal bool
		valid   bool
	}{
		{false, false},
		{true, true},
	}

	for n, c := range cases {
		b.FromZero(c.literal)
		if c.valid != b.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, b.Valid,
			)
		}
		if !b.Valid {
			continue
		}

		if c.literal != b.Bool {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %t, got %t)",
				t.Name(), n+1, c.literal, b.Bool,
			)
		}
	}
}

func TestBool_String(t *testing.T) {
	cases := []struct {
		nullable null.Bool
		string   string
	}{
		{null.Bool{Bool: false, Valid: true}, "false"},
		{null.Bool{Bool: true, Valid: true}, "true"},
		{null.Bool{}, "<invalid>"},
		{null.Bool{Bool: true, Valid: false}, "<invalid>"},
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

func TestBool_MarshalText(t *testing.T) {
	nilType := reflect.TypeOf(nil)

	cases := []struct {
		nullable null.Bool
		bytes    []byte
		errType  reflect.Type
	}{
		{null.Bool{Bool: false, Valid: true}, []byte("false"), nilType},
		{null.Bool{Bool: true, Valid: true}, []byte("true"), nilType},
		{null.Bool{}, nil, nilType},
		{null.Bool{Bool: true, Valid: false}, nil, nilType},
	}

	for n, c := range cases {
		b, err := c.nullable.MarshalText()
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if !reflect.DeepEqual(c.bytes, b) {
			t.Fatalf(
				"%s, case #%d: bytes mismatch (expected '%s', got '%s')",
				t.Name(), n+1, string(c.bytes), string(b),
			)
		}
	}
}

func TestBool_MarshalJSON(t *testing.T) {
	nilType := reflect.TypeOf(nil)

	cases := []struct {
		nullable null.Bool
		json    []byte
		errType  reflect.Type
	}{
		{null.Bool{Bool: false, Valid: true}, []byte("false"), nilType},
		{null.Bool{Bool: true, Valid: true}, []byte("true"), nilType},
		{null.Bool{}, []byte("null"), nilType},
		{null.Bool{Bool: true, Valid: false}, []byte("null"), nilType},
	}

	for n, c := range cases {
		b, err := c.nullable.MarshalJSON()
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if !reflect.DeepEqual(c.json, b) {
			t.Fatalf(
				"%s, case #%d: json mismatch (expected '%s', got '%s')",
				t.Name(), n+1, string(c.json), string(b),
			)
		}
	}
}

func TestBool_Value(t *testing.T) {
	nilType := reflect.TypeOf(nil)
	boolType := reflect.TypeOf(false)

	cases := []struct {
		nullable null.Bool
		expType  reflect.Type
		errType  reflect.Type
	}{
		{null.Bool{Bool: false, Valid: true}, boolType, nilType},
		{null.Bool{Bool: true, Valid: true}, boolType, nilType},
		{null.Bool{}, nilType, nilType},
		{null.Bool{Bool: true, Valid: false}, nilType, nilType},
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

		if c.nullable.Bool != reflect.ValueOf(v).Bool() {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %t, got %v)",
				t.Name(), n+1, c.nullable.Bool, v,
			)
		}
	}
}

func TestBool_Set(t *testing.T) {
	var b null.Bool
	nilType := reflect.TypeOf(nil)
	parseErrType := reflect.TypeOf(null.ParseError{})

	cases := []struct {
		string  string
		literal bool
		valid   bool
		errType reflect.Type
	}{
		{"false", false, true, nilType},
		{"true", true, true, nilType},
		{"FALSE", false, true, nilType},
		{"TRUE", true, true, nilType},
		{"False", false, true, nilType},
		{"True", true, true, nilType},
		{"F", false, true, nilType},
		{"T", true, true, nilType},
		{"f", false, true, nilType},
		{"t", true, true, nilType},
		{"0", false, true, nilType},
		{"1", true, true, nilType},
		{"", false, false, nilType},
		{"2", false, false, parseErrType},
		{"x", false, false, parseErrType},
	}

	for n, c := range cases {
		err := b.Set(c.string)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != b.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, b.Valid,
			)
		}
		if !b.Valid {
			continue
		}

		if c.literal != b.Bool {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %t, got %t)",
				t.Name(), n+1, c.literal, b.Bool,
			)
		}
	}
}

func TestBool_UnmarshalText(t *testing.T) {
	var b null.Bool
	nilType := reflect.TypeOf(nil)
	unmarshalErrType := reflect.TypeOf(null.UnmarshalError{})

	cases := []struct {
		bytes   []byte
		literal bool
		valid   bool
		errType reflect.Type
	}{
		{[]byte("false"), false, true, nilType},
		{[]byte("true"), true, true, nilType},
		{[]byte("FALSE"), false, true, nilType},
		{[]byte("TRUE"), true, true, nilType},
		{[]byte("False"), false, true, nilType},
		{[]byte("True"), true, true, nilType},
		{[]byte("F"), false, true, nilType},
		{[]byte("T"), true, true, nilType},
		{[]byte("f"), false, true, nilType},
		{[]byte("t"), true, true, nilType},
		{[]byte("0"), false, true, nilType},
		{[]byte("1"), true, true, nilType},
		{nil, false, false, nilType},
		{[]byte(""), false, false, nilType},
		{[]byte("2"), false, false, unmarshalErrType},
		{[]byte("x"), false, false, unmarshalErrType},
	}

	for n, c := range cases {
		err := b.UnmarshalText(c.bytes)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != b.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, b.Valid,
			)
		}
		if !b.Valid {
			continue
		}

		if c.literal != b.Bool {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %t, got %t)",
				t.Name(), n+1, c.literal, b.Bool,
			)
		}
	}
}

func TestBool_UnmarshalJSON(t *testing.T) {
	var b null.Bool
	nilType := reflect.TypeOf(nil)
	typeErrType := reflect.TypeOf(null.TypeError{})
	unmarshalErrType := reflect.TypeOf(null.UnmarshalError{})

	cases := []struct {
		json   []byte
		literal bool
		valid   bool
		errType reflect.Type
	}{
		{[]byte("false"), false, true, nilType},
		{[]byte("true"), true, true, nilType},
		{[]byte("FALSE"), false, false, unmarshalErrType},
		{[]byte("TRUE"), true, false, unmarshalErrType},
		{[]byte("False"), false, false, unmarshalErrType},
		{[]byte("True"), true, false, unmarshalErrType},
		{[]byte("F"), false, false, unmarshalErrType},
		{[]byte("T"), true, false, unmarshalErrType},
		{[]byte("f"), false, false, unmarshalErrType},
		{[]byte("t"), true, false, unmarshalErrType},
		{[]byte("0"), false, false, typeErrType},
		{[]byte("1"), true, false, typeErrType},
		{nil, false, false, unmarshalErrType},
		{[]byte(""), false, false, unmarshalErrType},
		{[]byte("2"), false, false, typeErrType},
		{[]byte("x"), false, false, unmarshalErrType},
	}

	for n, c := range cases {
		err := b.UnmarshalJSON(c.json)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != b.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, b.Valid,
			)
		}
		if !b.Valid {
			continue
		}

		if c.literal != b.Bool {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %t, got %t)",
				t.Name(), n+1, c.literal, b.Bool,
			)
		}
	}
}

func TestBool_Scan(t *testing.T) {
	var b null.Bool
	nilType := reflect.TypeOf(nil)
	typeErrType := reflect.TypeOf(null.TypeError{})

	cases := []struct {
		source  interface{}
		valid   bool
		errType reflect.Type
	}{
		{false, true, nilType},
		{true, true, nilType},
		{nil, false, nilType},
		{float64(0.1), false, typeErrType},
		{"x", false, typeErrType},
	}

	for n, c := range cases {
		err := b.Scan(c.source)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != b.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, b.Valid,
			)
		}
		if !b.Valid {
			continue
		}

		if c.source.(bool) != b.Bool {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %t, got %t)",
				t.Name(), n+1, c.source, b.Bool,
			)
		}
	}
}
