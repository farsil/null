package null_test

import (
	"null"
	"reflect"
	"testing"
	"time"
)

func TestTimeFrom(t *testing.T) {
	zero := time.Time{}
	now := time.Now()
	future := now.AddDate(10000, 0, 0)

	cases := []struct {
		literal time.Time
		valid   bool
	}{
		{zero, true},
		{now, true},
		{future, true},
	}

	for n, c := range cases {
		tm := null.TimeFrom(c.literal)
		if c.valid != tm.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, tm.Valid,
			)
		}
		if !tm.Valid {
			continue
		}

		if !c.literal.Equal(tm.Time) {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %s, got %s)",
				t.Name(), n+1, c.literal, tm.Time,
			)
		}
	}
}

func TestTimeFromPtr(t *testing.T) {
	zero := time.Time{}
	now := time.Now()
	future := now.AddDate(10000, 0, 0)

	cases := []struct {
		ptr   *time.Time
		valid bool
	}{
		{&zero, true},
		{&now, true},
		{&future, true},
		{nil, false},
	}

	for n, c := range cases {
		tm := null.TimeFromPtr(c.ptr)
		if c.valid != tm.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, tm.Valid,
			)
		}
		if !tm.Valid {
			continue
		}

		if !c.ptr.Equal(tm.Time) {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %s, got %s)",
				t.Name(), n+1, *c.ptr, tm.Time,
			)
		}
	}
}

func TestTimeFromZero(t *testing.T) {
	zero := time.Time{}
	now := time.Now()
	future := now.AddDate(10000, 0, 0)

	cases := []struct {
		literal time.Time
		valid   bool
	}{
		{zero, false},
		{now, true},
		{future, true},
	}

	for n, c := range cases {
		tm := null.TimeFromZero(c.literal)
		if c.valid != tm.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, tm.Valid,
			)
		}
		if !tm.Valid {
			continue
		}

		if !c.literal.Equal(tm.Time) {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %s, got %s)",
				t.Name(), n+1, c.literal, tm.Time,
			)
		}
	}
}

func TestTime_Ptr(t *testing.T) {
	zero := time.Time{}
	now := time.Now()
	future := now.AddDate(10000, 0, 0)

	cases := []struct {
		nullable null.Time
	}{
		{null.Time{Time: zero, Valid: true}},
		{null.Time{Time: now, Valid: true}},
		{null.Time{Time: future, Valid: true}},
		{null.Time{}},
		{null.Time{Time: future, Valid: false}},
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
			if !c.nullable.Time.Equal(*p) {
				t.Fatalf(
					"%s, case #%d: literal mismatch (expected %s, got %s)",
					t.Name(), n+1, c.nullable.Time, *p,
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

func TestTime_Zero(t *testing.T) {
	zero := time.Time{}
	now := time.Now()
	future := now.AddDate(10000, 0, 0)

	cases := []struct {
		nullable null.Time
	}{
		{null.Time{Time: zero, Valid: true}},
		{null.Time{Time: now, Valid: true}},
		{null.Time{Time: future, Valid: true}},
		{null.Time{}},
		{null.Time{Time: future, Valid: false}},
	}

	for n, c := range cases {
		tm := c.nullable.Zero()
		if c.nullable.Valid {
			if !c.nullable.Time.Equal(tm) {
				t.Fatalf(
					"%s, case #%d: literal mismatch (expected %s, got %s)",
					t.Name(), n+1, c.nullable.Time, tm,
				)
			}
		} else {
			if !tm.IsZero() {
				t.Fatalf(
					"%s, case #%d: non-zero uint returned",
					t.Name(), n+1,
				)
			}
		}
	}
}

func TestTime_From(t *testing.T) {
	var tm null.Time
	zero := time.Time{}
	now := time.Now()
	future := now.AddDate(10000, 0, 0)

	cases := []struct {
		literal time.Time
		valid   bool
	}{
		{zero, true},
		{now, true},
		{future, true},
	}

	for n, c := range cases {
		tm.From(c.literal)
		if c.valid != tm.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, tm.Valid,
			)
		}
		if !tm.Valid {
			continue
		}

		if !c.literal.Equal(tm.Time) {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %s, got %s)",
				t.Name(), n+1, c.literal, tm.Time,
			)
		}
	}
}

func TestTime_FromPtr(t *testing.T) {
	var tm null.Time
	zero := time.Time{}
	now := time.Now()
	future := now.AddDate(10000, 0, 0)

	cases := []struct {
		ptr   *time.Time
		valid bool
	}{
		{&zero, true},
		{&now, true},
		{&future, true},
		{nil, false},
	}

	for n, c := range cases {
		tm.FromPtr(c.ptr)
		if c.valid != tm.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, tm.Valid,
			)
		}
		if !tm.Valid {
			continue
		}

		if !c.ptr.Equal(tm.Time) {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %s, got %s)",
				t.Name(), n+1, *c.ptr, tm.Time,
			)
		}
	}
}

func TestTime_FromZero(t *testing.T) {
	var tm null.Time
	zero := time.Time{}
	now := time.Now()
	future := now.AddDate(10000, 0, 0)

	cases := []struct {
		literal time.Time
		valid   bool
	}{
		{zero, false},
		{now, true},
		{future, true},
	}

	for n, c := range cases {
		tm.FromZero(c.literal)
		if c.valid != tm.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, tm.Valid,
			)
		}
		if !tm.Valid {
			continue
		}

		if !c.literal.Equal(tm.Time) {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %s, got %s)",
				t.Name(), n+1, c.literal, tm.Time,
			)
		}
	}
}

func TestTime_String(t *testing.T) {
	zero := time.Time{}
	now := time.Now()
	future := now.AddDate(10000, 0, 0)

	zeroStr := zero.Format(time.RFC3339Nano)
	nowStr := now.Format(time.RFC3339Nano)
	futureStr := future.Format(time.RFC3339Nano)

	cases := []struct {
		nullable null.Time
		string   string
	}{
		{null.Time{Time: zero, Valid: true}, zeroStr},
		{null.Time{Time: now, Valid: true}, nowStr},
		{null.Time{Time: future, Valid: true}, futureStr},
		{null.Time{}, "<invalid>"},
		{null.Time{Time: future, Valid: false}, "<invalid>"},
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

func TestTime_MarshalText(t *testing.T) {
	zero := time.Time{}
	now := time.Now()
	future := now.AddDate(10000, 0, 0)

	zeroBytes, _ := zero.MarshalText()
	nowBytes, _ := now.MarshalText()
	// future cannot be marshaled, RFC3339 does not allow year > 10000

	nilType := reflect.TypeOf(nil)
	marshalErrType := reflect.TypeOf(null.MarshalError{})

	cases := []struct {
		nullable null.Time
		bytes    []byte
		errType  reflect.Type
	}{
		{null.Time{Time: zero, Valid: true}, zeroBytes, nilType},
		{null.Time{Time: now, Valid: true}, nowBytes, nilType},
		{null.Time{Time: future, Valid: true}, nil, marshalErrType},
		{null.Time{}, nil, nilType},
		{null.Time{Time: future, Valid: false}, nil, nilType},
	}

	for n, c := range cases {
		tm, err := c.nullable.MarshalText()
		if c.errType != reflect.TypeOf(err) {
			t.Log(err)
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if !reflect.DeepEqual(c.bytes, tm) {
			t.Fatalf(
				"%s, case #%d: bytes mismatch (expected '%s', got '%s')",
				t.Name(), n+1, string(c.bytes), string(tm),
			)
		}
	}
}

func TestTime_MarshalJSON(t *testing.T) {
	zero := time.Time{}
	now := time.Now()
	future := now.AddDate(10000, 0, 0)

	zeroBytes, _ := zero.MarshalJSON()
	nowBytes, _ := now.MarshalJSON()
	// future cannot be marshaled, RFC3339 does not allow year > 10000

	nilType := reflect.TypeOf(nil)
	marshalErrType := reflect.TypeOf(null.MarshalError{})

	cases := []struct {
		nullable null.Time
		json     []byte
		errType  reflect.Type
	}{
		{null.Time{Time: zero, Valid: true}, zeroBytes, nilType},
		{null.Time{Time: now, Valid: true}, nowBytes, nilType},
		{null.Time{Time: future, Valid: true}, nil, marshalErrType},
		{null.Time{}, []byte("null"), nilType},
		{null.Time{Time: future, Valid: false}, []byte("null"), nilType},
	}

	for n, c := range cases {
		tm, err := c.nullable.MarshalJSON()
		if c.errType != reflect.TypeOf(err) {
			t.Log(err)
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if !reflect.DeepEqual(c.json, tm) {
			t.Fatalf(
				"%s, case #%d: json mismatch (expected '%s', got '%s')",
				t.Name(), n+1, string(c.json), string(tm),
			)
		}
	}
}

func TestTime_Value(t *testing.T) {
	zero := time.Time{}
	now := time.Now()
	future := now.AddDate(10000, 0, 0)

	nilType := reflect.TypeOf(nil)
	timeType := reflect.TypeOf(zero)

	cases := []struct {
		nullable null.Time
		expType  reflect.Type
		errType  reflect.Type
	}{
		{null.Time{Time: zero, Valid: true}, timeType, nilType},
		{null.Time{Time: now, Valid: true}, timeType, nilType},
		{null.Time{Time: future, Valid: true}, timeType, nilType},
		{null.Time{}, nilType, nilType},
		{null.Time{Time: future, Valid: false}, nilType, nilType},
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

		if !c.nullable.Time.Equal(reflect.ValueOf(v).Interface().(time.Time)) {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %s, got %v)",
				t.Name(), n+1, c.nullable.Time, v,
			)
		}
	}
}

func TestTime_Set(t *testing.T) {
	var tm null.Time
	zero := time.Time{}
	now := time.Now()
	future := now.AddDate(10000, 0, 0)

	zeroStr := zero.Format(time.RFC3339Nano)
	nowStr := now.Format(time.RFC3339Nano)
	futureStr := future.Format(time.RFC3339Nano)

	nilType := reflect.TypeOf(nil)
	parseErrType := reflect.TypeOf(null.ParseError{})

	cases := []struct {
		string  string
		literal time.Time
		valid   bool
		errType reflect.Type
	}{
		{zeroStr, zero, true, nilType},
		{nowStr, now, true, nilType},
		{futureStr, future, false, parseErrType},
		{"", zero, false, nilType},
		{"-1", zero, false, parseErrType},
		{"0.1", zero, false, parseErrType},
		{"x", zero, false, parseErrType},
	}

	for n, c := range cases {
		err := tm.Set(c.string)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != tm.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, tm.Valid,
			)
		}
		if !tm.Valid {
			continue
		}

		if !c.literal.Equal(tm.Time) {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %s, got %s)",
				t.Name(), n+1, c.literal, tm.Time,
			)
		}
	}
}

func TestTime_UnmarshalText(t *testing.T) {
	var tm null.Time
	zero := time.Time{}
	now := time.Now()
	future := now.AddDate(10000, 0, 0)

	zeroBytes, _ := zero.MarshalText()
	nowBytes, _ := now.MarshalText()
	futureBytes := []byte(future.Format(time.RFC3339Nano))

	nilType := reflect.TypeOf(nil)
	unmarshalErrType := reflect.TypeOf(null.UnmarshalError{})

	cases := []struct {
		bytes   []byte
		literal time.Time
		valid   bool
		errType reflect.Type
	}{
		{zeroBytes, zero, true, nilType},
		{nowBytes, now, true, nilType},
		{futureBytes, future, false, unmarshalErrType},
		{nil, zero, false, nilType},
		{[]byte(""), zero, false, nilType},
		{[]byte("-1"), zero, false, unmarshalErrType},
		{[]byte("0.1"), zero, false, unmarshalErrType},
		{[]byte("x"), zero, false, unmarshalErrType},
	}

	for n, c := range cases {
		err := tm.UnmarshalText(c.bytes)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != tm.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, tm.Valid,
			)
		}
		if !tm.Valid {
			continue
		}

		if !c.literal.Equal(tm.Time) {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %s, got %s)",
				t.Name(), n+1, c.literal, tm.Time,
			)
		}
	}
}

func TestTime_UnmarshalJSON(t *testing.T) {
	var tm null.Time
	zero := time.Time{}
	now := time.Now()
	future := now.AddDate(10000, 0, 0)

	zeroJSON, _ := zero.MarshalJSON()
	nowJSON, _ := now.MarshalJSON()
	futureJSON := []byte(`"` + future.Format(time.RFC3339Nano) + `"`)

	nilType := reflect.TypeOf(nil)
	parseErrType := reflect.TypeOf(null.ParseError{})
	typeErrType := reflect.TypeOf(null.TypeError{})
	unmarshalErrType := reflect.TypeOf(null.UnmarshalError{})

	cases := []struct {
		json    []byte
		literal time.Time
		valid   bool
		errType reflect.Type
	}{
		{zeroJSON, zero, true, nilType},
		{nowJSON, now, true, nilType},
		{futureJSON, future, true, parseErrType},
		{[]byte("null"), zero, false, nilType},
		{nil, zero, false, unmarshalErrType},
		{[]byte("-1"), zero, false, typeErrType},
		{[]byte("0.1"), zero, false, typeErrType},
		{[]byte("x"), zero, false, unmarshalErrType},
	}

	for n, c := range cases {
		err := tm.UnmarshalJSON(c.json)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != tm.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, tm.Valid,
			)
		}
		if !tm.Valid {
			continue
		}

		if !c.literal.Equal(tm.Time) {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %s, got %s)",
				t.Name(), n+1, c.literal, tm.Time,
			)
		}
	}
}

func TestTime_Scan(t *testing.T) {
	var tm null.Time
	zero := time.Time{}
	now := time.Now()
	future := now.AddDate(10000, 0, 0)

	nilType := reflect.TypeOf(nil)
	typeErrType := reflect.TypeOf(null.TypeError{})

	cases := []struct {
		source  interface{}
		valid   bool
		errType reflect.Type
	}{
		{zero, true, nilType},
		{now, true, nilType},
		{future, true, nilType},
		{nil, false, nilType},
		{float64(0.1), false, typeErrType},
		{"x", false, typeErrType},
	}

	for n, c := range cases {
		err := tm.Scan(c.source)
		if c.errType != reflect.TypeOf(err) {
			t.Fatalf(
				"%s, case #%d: wrong error type (expected %v, got %v)",
				t.Name(), n+1, c.errType, reflect.TypeOf(err),
			)
		}
		if err != nil {
			continue
		}

		if c.valid != tm.Valid {
			t.Fatalf(
				"%s, case #%d: validity mismatch (expected %t, got %t)",
				t.Name(), n+1, c.valid, tm.Valid,
			)
		}
		if !tm.Valid {
			continue
		}

		if c.source.(time.Time) != tm.Time {
			t.Fatalf(
				"%s, case #%d: literal mismatch (expected %s, got %s)",
				t.Name(), n+1, c.source, tm.Time,
			)
		}
	}
}
