package null_test

import (
	"null"
	"reflect"
	"testing"
)

func float64p(v float64) *float64 {
	return &v
}

func TestFloat64From(t *testing.T) {
	cases := []struct {
		literal float64
		valid   bool
	}{
		{0.0, true},
		{3.0, true},
		{0.23, true},
		{-1E+10, true},
	}

	for n, c := range cases {
		f := null.Float64From(c.literal)
		if c.valid != f.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, f.Valid,
			)
		}
		if !f.Valid {
			continue
		}

		if c.literal != f.Float64 {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %g, got %g)",
				t.Name(), n+1, c.literal, f.Float64,
			)
		}
	}
}

func TestFloat64FromPtr(t *testing.T) {
	cases := []struct {
		ptr   *float64
		valid bool
	}{
		{float64p(0.0), true},
		{float64p(3.0), true},
		{float64p(0.23), true},
		{float64p(-1E+10), true},
		{nil, false},
	}

	for n, c := range cases {
		f := null.Float64FromPtr(c.ptr)
		if c.valid != f.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, f.Valid,
			)
		}
		if !f.Valid {
			continue
		}

		if *c.ptr != f.Float64 {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %g, got %g)",
				t.Name(), n+1, *c.ptr, f.Float64,
			)
		}
	}
}

func TestFloat64FromZero(t *testing.T) {
	cases := []struct {
		literal float64
		valid   bool
	}{
		{0.0, false},
		{3.0, true},
		{0.23, true},
		{-1E+10, true},
	}

	for n, c := range cases {
		f := null.Float64FromZero(c.literal)
		if c.valid != f.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, f.Valid,
			)
		}
		if !f.Valid {
			continue
		}

		if c.literal != f.Float64 {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %g, got %g)",
				t.Name(), n+1, c.literal, f.Float64,
			)
		}
	}
}

func TestFloat64_Ptr(t *testing.T) {
	cases := []struct {
		nullable null.Float64
	}{
		{null.Float64{Float64: 0.0, Valid: true}},
		{null.Float64{Float64: 3.0, Valid: true}},
		{null.Float64{Float64: 0.23, Valid: true}},
		{null.Float64{Float64: -1E+10, Valid: true}},
		{null.Float64{}},
		{null.Float64{Float64: 12.34, Valid: false}},
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
			if c.nullable.Float64 != *p {
				t.Fatalf(
					"%s, case #%d: literal mismatch (expected %g, got %g)",
					t.Name(), n+1, c.nullable.Float64, *p,
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

func TestFloat64_Zero(t *testing.T) {
	cases := []struct {
		nullable null.Float64
	}{
		{null.Float64{Float64: 0.0, Valid: true}},
		{null.Float64{Float64: 3.0, Valid: true}},
		{null.Float64{Float64: 0.23, Valid: true}},
		{null.Float64{Float64: -1E+10, Valid: true}},
		{null.Float64{}},
		{null.Float64{Float64: 12.34, Valid: false}},
	}

	for n, c := range cases {
		f := c.nullable.Zero()
		if c.nullable.Valid {
			if c.nullable.Float64 != f {
				t.Fatalf(
					"%s, case #%d: literal mismatch (expected %g, got %g)",
					t.Name(), n+1, c.nullable.Float64, f,
				)
			}
		} else {
			if f != 0.0 {
				t.Fatalf(
					"%s, case #%d: non-zero uint returned",
					t.Name(), n+1,
				)
			}
		}
	}
}

func TestFloat64_From(t *testing.T) {
	var f null.Float64
	cases := []struct {
		literal float64
		valid   bool
	}{
		{0.0, true},
		{3.0, true},
		{0.23, true},
		{-1E+10, true},
	}

	for n, c := range cases {
		f.From(c.literal)
		if c.valid != f.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, f.Valid,
			)
		}
		if !f.Valid {
			continue
		}

		if c.literal != f.Float64 {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %g, got %g)",
				t.Name(), n+1, c.literal, f.Float64,
			)
		}
	}
}

func TestFloat64_FromPtr(t *testing.T) {
	var f null.Float64
	cases := []struct {
		ptr   *float64
		valid bool
	}{
		{float64p(0.0), true},
		{float64p(3.0), true},
		{float64p(0.23), true},
		{float64p(-1E+10), true},
		{nil, false},
	}

	for n, c := range cases {
		f.FromPtr(c.ptr)
		if c.valid != f.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, f.Valid,
			)
		}
		if !f.Valid {
			continue
		}

		if *c.ptr != f.Float64 {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %g, got %g)",
				t.Name(), n+1, *c.ptr, f.Float64,
			)
		}
	}
}

func TestFloat64_FromZero(t *testing.T) {
	var f null.Float64
	cases := []struct {
		literal float64
		valid   bool
	}{
		{0.0, false},
		{3.0, true},
		{0.23, true},
		{-1E+10, true},
	}

	for n, c := range cases {
		f.FromZero(c.literal)
		if c.valid != f.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, f.Valid,
			)
		}
		if !f.Valid {
			continue
		}

		if c.literal != f.Float64 {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %g, got %g)",
				t.Name(), n+1, c.literal, f.Float64,
			)
		}
	}
}

func TestFloat64_String(t *testing.T) {
	cases := []struct {
		nullable null.Float64
		string   string
	}{
		{null.Float64{Float64: 0.0, Valid: true}, "0"},
		{null.Float64{Float64: 3.0, Valid: true}, "3"},
		{null.Float64{Float64: 0.23, Valid: true}, "0.23"},
		{null.Float64{Float64: -1E+10, Valid: true}, "-1e+10"},
		{null.Float64{}, "<invalid>"},
		{null.Float64{Float64: 12.34, Valid: false}, "<invalid>"},
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

func TestFloat64_MarshalText(t *testing.T) {
	nilType := reflect.TypeOf(nil)

	cases := []struct {
		nullable null.Float64
		bytes    []byte
		errType  reflect.Type
	}{
		{
			null.Float64{Float64: 0.0, Valid: true},
			[]byte("0"), nilType,
		},
		{
			null.Float64{Float64: 3.0, Valid: true},
			[]byte("3"), nilType,
		},
		{
			null.Float64{Float64: 0.23, Valid: true},
			[]byte("0.23"), nilType,
		},
		{
			null.Float64{Float64: -1E+10, Valid: true},
			[]byte("-1e+10"), nilType,
		},
		{
			null.Float64{},
			nil, nilType,
		},
		{
			null.Float64{Float64: 12.34, Valid: false},
			nil, nilType,
		},
	}

	for n, c := range cases {
		f, err := c.nullable.MarshalText()
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if !reflect.DeepEqual(c.bytes, f) {
			t.Fatalf(
				"%s, case #%d: bytes mismatch (expected '%s', got '%s')",
				t.Name(), n+1, string(c.bytes), string(f),
			)
		}
	}
}

func TestFloat64_MarshalJSON(t *testing.T) {
	nilType := reflect.TypeOf(nil)

	cases := []struct {
		nullable null.Float64
		json     []byte
		errType  reflect.Type
	}{
		{
			null.Float64{Float64: 0.0, Valid: true},
			[]byte("0"), nilType,
		},
		{
			null.Float64{Float64: 3.0, Valid: true},
			[]byte("3"), nilType,
		},
		{
			null.Float64{Float64: 0.23, Valid: true},
			[]byte("0.23"), nilType,
		},
		{
			null.Float64{Float64: -1E+10, Valid: true},
			[]byte("-10000000000"), nilType,
		},
		{
			null.Float64{},
			[]byte("null"), nilType,
		},
		{
			null.Float64{Float64: 12.34, Valid: false},
			[]byte("null"), nilType,
		},
	}

	for n, c := range cases {
		f, err := c.nullable.MarshalJSON()
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if !reflect.DeepEqual(c.json, f) {
			t.Fatalf(
				"%s, case #%d: json mismatch (expected '%s', got '%s')",
				t.Name(), n+1, string(c.json), string(f),
			)
		}
	}
}

func TestFloat64_Value(t *testing.T) {
	nilType := reflect.TypeOf(nil)
	f64Type := reflect.TypeOf(float64(0.0))

	cases := []struct {
		nullable null.Float64
		expType  reflect.Type
		errType  reflect.Type
	}{
		{null.Float64{Float64: 0.0, Valid: true}, f64Type, nilType},
		{null.Float64{Float64: 3.0, Valid: true}, f64Type, nilType},
		{null.Float64{Float64: 0.23, Valid: true}, f64Type, nilType},
		{null.Float64{Float64: -1E+10, Valid: true}, f64Type, nilType},
		{null.Float64{}, nilType, nilType},
		{null.Float64{Float64: 12.34, Valid: false}, nilType, nilType},
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

		if c.nullable.Float64 != reflect.ValueOf(v).Float() {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %g, got %v)",
				t.Name(), n+1, c.nullable.Float64, v,
			)
		}
	}
}

func TestFloat64_Set(t *testing.T) {
	var f null.Float64
	nilType := reflect.TypeOf(nil)
	parseErrType := reflect.TypeOf(null.ParseError{})

	cases := []struct {
		string  string
		literal float64
		valid   bool
		errType reflect.Type
	}{
		{"0", 0, true, nilType},
		{"3", 3, true, nilType},
		{"0.23", 0.23, true, nilType},
		{"-1e+10", -1E+10, true, nilType},
		{"", 0, false, nilType},
		{"x", 0, false, parseErrType},
	}

	for n, c := range cases {
		err := f.Set(c.string)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != f.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, f.Valid,
			)
		}
		if !f.Valid {
			continue
		}

		if c.literal != f.Float64 {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %g, got %g)",
				t.Name(), n+1, c.literal, f.Float64,
			)
		}
	}
}

func TestFloat64_UnmarshalText(t *testing.T) {
	var f null.Float64
	nilType := reflect.TypeOf(nil)
	unmarshalErrType := reflect.TypeOf(null.UnmarshalError{})

	cases := []struct {
		bytes   []byte
		literal float64
		valid   bool
		errType reflect.Type
	}{
		{[]byte("0"), 0, true, nilType},
		{[]byte("3"), 3, true, nilType},
		{[]byte("0.23"), 0.23, true, nilType},
		{[]byte("-1e+10"), -1E+10, true, nilType},
		{nil, 0, false, nilType},
		{[]byte(""), 0, false, nilType},
		{[]byte("x"), 0, false, unmarshalErrType},
	}

	for n, c := range cases {
		err := f.UnmarshalText(c.bytes)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != f.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, f.Valid,
			)
		}
		if !f.Valid {
			continue
		}

		if c.literal != f.Float64 {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %g, got %g)",
				t.Name(), n+1, c.literal, f.Float64,
			)
		}
	}
}

func TestFloat64_UnmarshalJSON(t *testing.T) {
	var f null.Float64
	nilType := reflect.TypeOf(nil)
	typeErrType := reflect.TypeOf(null.TypeError{})
	unmarshalErrType := reflect.TypeOf(null.UnmarshalError{})

	cases := []struct {
		json    []byte
		literal float64
		valid   bool
		errType reflect.Type
	}{
		{[]byte("0"), 0, true, nilType},
		{[]byte("3"), 3, true, nilType},
		{[]byte("0.23"), 0.23, true, nilType},
		{[]byte("-10000000000"), -1E+10, true, nilType},
		{[]byte("null"), 0, false, nilType},
		{[]byte(`"x"`), 0, false, typeErrType},
		{nil, 0, false, unmarshalErrType},
		{[]byte("x"), 0, false, unmarshalErrType},
	}

	for n, c := range cases {
		err := f.UnmarshalJSON(c.json)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != f.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, f.Valid,
			)
		}
		if !f.Valid {
			continue
		}

		if c.literal != f.Float64 {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %g, got %g)",
				t.Name(), n+1, c.literal, f.Float64,
			)
		}
	}
}

func TestFloat64_Scan(t *testing.T) {
	var f null.Float64
	nilType := reflect.TypeOf(nil)
	typeErrType := reflect.TypeOf(null.TypeError{})

	cases := []struct {
		source  interface{}
		valid   bool
		errType reflect.Type
	}{
		{float64(0.0), true, nilType},
		{float64(3.0), true, nilType},
		{float64(0.23), true, nilType},
		{float64(-1E+10), true, nilType},
		{nil, false, nilType},
		{"x", false, typeErrType},
	}

	for n, c := range cases {
		err := f.Scan(c.source)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != f.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, f.Valid,
			)
		}
		if !f.Valid {
			continue
		}

		if c.source.(float64) != f.Float64 {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %g, got %g)",
				t.Name(), n+1, c.source, f.Float64,
			)
		}
	}
}
