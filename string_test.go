package null_test

import (
	"null"
	"reflect"
	"testing"
)

func stringp(v string) *string {
	return &v
}

func TestStringFrom(t *testing.T) {
	cases := []struct {
		literal string
		valid   bool
	}{
		{"", true},
		{"foo", true},
		{"bar", true},
	}

	for n, c := range cases {
		s := null.StringFrom(c.literal)
		if c.valid != s.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, s.Valid,
			)
		}
		if !s.Valid {
			continue
		}

		if c.literal != s.Str {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected '%s', got '%s')",
				t.Name(), n+1, c.literal, s.Str,
			)
		}
	}
}

func TestStringFromPtr(t *testing.T) {
	cases := []struct {
		ptr   *string
		valid bool
	}{
		{stringp("foo"), true},
		{stringp("bar"), true},
		{stringp(""), true},
		{nil, false},
	}

	for n, c := range cases {
		s := null.StringFromPtr(c.ptr)
		if c.valid != s.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, s.Valid,
			)
		}
		if !s.Valid {
			continue
		}

		if *c.ptr != s.Str {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected '%s', got '%s')",
				t.Name(), n+1, *c.ptr, s.Str,
			)
		}
	}
}

func TestStringFromZero(t *testing.T) {
	cases := []struct {
		literal string
		valid   bool
	}{
		{"foo", true},
		{"bar", true},
		{"", false},
	}

	for n, c := range cases {
		s := null.StringFromZero(c.literal)
		if c.valid != s.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, s.Valid,
			)
		}
		if !s.Valid {
			continue
		}

		if c.literal != s.Str {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected '%s', got '%s')",
				t.Name(), n+1, c.literal, s.Str,
			)
		}
	}
}

func TestString_Ptr(t *testing.T) {
	cases := []struct {
		nullable null.String
	}{
		{null.String{Str: "foo", Valid: true}},
		{null.String{Str: "bar", Valid: true}},
		{null.String{}},
		{null.String{Str: "baz", Valid: false}},
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
			if c.nullable.Str != *p {
				t.Fatalf(
					"%s, case #%d: literal mismatch (expected '%s', got '%s')",
					t.Name(), n+1, c.nullable.Str, *p,
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

func TestString_Zero(t *testing.T) {
	cases := []struct {
		nullable null.String
	}{
		{null.String{Str: "foo", Valid: true}},
		{null.String{Str: "bar", Valid: true}},
		{null.String{}},
		{null.String{Str: "baz", Valid: false}},
	}

	for n, c := range cases {
		s := c.nullable.Zero()
		if c.nullable.Valid {
			if c.nullable.Str != s {
				t.Fatalf(
					"%s, case #%d: literal mismatch (expected '%s', got '%s')",
					t.Name(), n+1, c.nullable.Str, s,
				)
			}
		} else {
			if s != "" {
				t.Fatalf(
					"%s, case #%d: non-empty string returned",
					t.Name(), n+1,
				)
			}
		}
	}
}

func TestString_From(t *testing.T) {
	var s null.String
	cases := []struct {
		literal string
		valid   bool
	}{
		{"foo", true},
		{"bar", true},
	}

	for n, c := range cases {
		s.From(c.literal)
		if c.valid != s.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, s.Valid,
			)
		}
		if !s.Valid {
			continue
		}

		if c.literal != s.Str {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected '%s', got '%s')",
				t.Name(), n+1, c.literal, s.Str,
			)
		}
	}
}

func TestString_FromPtr(t *testing.T) {
	var s null.String
	cases := []struct {
		ptr   *string
		valid bool
	}{
		{stringp("foo"), true},
		{stringp("bar"), true},
		{nil, false},
	}

	for n, c := range cases {
		s.FromPtr(c.ptr)
		if c.valid != s.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, s.Valid,
			)
		}
		if !s.Valid {
			continue
		}

		if *c.ptr != s.Str {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected '%s', got '%s')",
				t.Name(), n+1, *c.ptr, s.Str,
			)
		}
	}
}

func TestString_FromZero(t *testing.T) {
	var s null.String
	cases := []struct {
		literal string
		valid   bool
	}{
		{"foo", true},
		{"bar", true},
		{"", false},
	}

	for n, c := range cases {
		s.FromZero(c.literal)
		if c.valid != s.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, s.Valid,
			)
		}
		if !s.Valid {
			continue
		}

		if c.literal != s.Str {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected '%s', got '%s')",
				t.Name(), n+1, c.literal, s.Str,
			)
		}
	}
}

func TestString_String(t *testing.T) {
	cases := []struct {
	nullable null.String
	string   string
}{
	{null.String{Str: "foo", Valid: true}, "foo"},
	{null.String{Str: "bar", Valid: true}, "bar"},
	{null.String{}, "<invalid>"},
	{null.String{Str: "baz", Valid: false}, "<invalid>"},
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

func TestString_MarshalText(t *testing.T) {
	nilType := reflect.TypeOf(nil)

	cases := []struct {
		nullable null.String
		bytes    []byte
		errType  reflect.Type
	}{
		{null.String{Str: "foo", Valid: true}, []byte("foo"), nilType},
		{null.String{Str: "bar", Valid: true}, []byte("bar"), nilType},
		{null.String{}, nil, nilType},
		{null.String{Str: "baz", Valid: false}, nil, nilType},
	}

	for n, c := range cases {
		s, err := c.nullable.MarshalText()
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if !reflect.DeepEqual(c.bytes, s) {
			t.Fatalf(
				"%s, case #%d: bytes mismatch (expected '%s', got '%s')",
				t.Name(), n+1, string(c.bytes), string(s),
			)
		}
	}
}

func TestString_MarshalJSON(t *testing.T) {
	nilType := reflect.TypeOf(nil)

	cases := []struct {
		nullable null.String
		bytes    []byte
		errType  reflect.Type
	}{
		{null.String{Str: "foo", Valid: true}, []byte(`"foo"`), nilType},
		{null.String{Str: "bar", Valid: true}, []byte(`"bar"`), nilType},
		{null.String{}, []byte("null"), nilType},
		{null.String{Str: "baz", Valid: false}, []byte("null"), nilType},
	}

	for n, c := range cases {
		s, err := c.nullable.MarshalJSON()
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if !reflect.DeepEqual(c.bytes, s) {
			t.Fatalf(
				"%s, case #%d: json mismatch (expected '%s', got '%s')",
				t.Name(), n+1, string(c.bytes), string(s),
			)
		}
	}
}

func TestString_Value(t *testing.T) {
	nilType := reflect.TypeOf(nil)
	strType := reflect.TypeOf("")

	cases := []struct {
		nullable null.String
		expType  reflect.Type
		errType  reflect.Type
	}{
		{null.String{Str: "foo", Valid: true}, strType, nilType},
		{null.String{Str: "bar", Valid: true}, strType, nilType},
		{null.String{}, nilType, nilType},
		{null.String{Str: "baz", Valid: false}, nilType, nilType},
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

		if c.nullable.Str != reflect.ValueOf(v).String() {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected '%s', got %v)",
				t.Name(), n+1, c.nullable.Str, v,
			)
		}
	}
}

func TestString_Set(t *testing.T) {
	var s null.String
	nilType := reflect.TypeOf(nil)

	cases := []struct {
		literal string
		valid   bool
		errType reflect.Type
	}{
		{"foo", true, nilType},
		{"bar", true, nilType},
		{"", false, nilType},
	}

	for n, c := range cases {
		err := s.Set(c.literal)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != s.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, s.Valid,
			)
		}
		if !s.Valid {
			continue
		}

		if c.literal != s.Str {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected '%s', got '%s')",
				t.Name(), n+1, c.literal, s.Str,
			)
		}
	}
}

func TestString_UnmarshalText(t *testing.T) {
	var s null.String
	nilType := reflect.TypeOf(nil)

	cases := []struct {
		bytes []byte
		literal string
		valid   bool
		errType reflect.Type
	}{
		{[]byte("foo"), "foo", true, nilType},
		{[]byte("bar"), "bar", true, nilType},
		{nil, "", false, nilType},
		{[]byte(""), "", true, nilType},
	}

	for n, c := range cases {
		err := s.UnmarshalText(c.bytes)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != s.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, s.Valid,
			)
		}
		if !s.Valid {
			continue
		}

		if c.literal != s.Str {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected '%s', got '%s')",
				t.Name(), n+1, c.literal, s.Str,
			)
		}
	}
}

func TestString_UnmarshalJSON(t *testing.T) {
	var s null.String
	nilType := reflect.TypeOf(nil)
	typeErrType := reflect.TypeOf(null.TypeError{})
	unmarshalErrType := reflect.TypeOf(null.UnmarshalError{})

	cases := []struct {
		json   []byte
		literal string
		valid   bool
		errType reflect.Type
	}{
		{[]byte(`"foo"`), "foo", true, nilType},
		{[]byte(`"bar"`), "bar", true, nilType},
		{[]byte(`""`), "", true, nilType},
		{nil, "", false, unmarshalErrType},
		{[]byte(""), "", true, unmarshalErrType},
		{[]byte("1"), "", true, typeErrType},
		{[]byte("x"), "", false, unmarshalErrType},
	}

	for n, c := range cases {
		err := s.UnmarshalJSON(c.json)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != s.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, s.Valid,
			)
		}
		if !s.Valid {
			continue
		}

		if c.literal != s.Str {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected '%s', got '%s')",
				t.Name(), n+1, c.literal, s.Str,
			)
		}
	}
}

func TestString_Scan(t *testing.T) {
	var s null.String
	nilType := reflect.TypeOf(nil)
	typeErrType := reflect.TypeOf(null.TypeError{})

	cases := []struct {
		source  interface{}
		valid   bool
		errType reflect.Type
	}{
		{"foo", true, nilType},
		{"bar", true, nilType},
		{nil, false, nilType},
		{float64(0.1), false, typeErrType},
	}

	for n, c := range cases {
		err := s.Scan(c.source)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != s.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, s.Valid,
			)
		}
		if !s.Valid {
			continue
		}

		if c.source.(string) != s.Str {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected '%s', got '%s')",
				t.Name(), n+1, c.source, s.Str,
			)
		}
	}
}
