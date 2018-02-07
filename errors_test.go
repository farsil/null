package null

import (
	"fmt"
	"strings"
	"testing"
)

func TestMarshalError_Error(t *testing.T) {
	cases := []struct {
		prefix string
		src    interface{}
	}{
		{"test", true},
		{"test", 1},
		{"test", 0.23},
		{"test", "foo"},
	}

	for n, c := range cases {
		err := makeMarshalError(c.prefix, c.src).(MarshalError)

		str := fmt.Sprintf(
			"%s: cannot marshal %v of type %s",
			err.prefix, err.SrcValue, err.SrcType,
		)

		if str != err.Error() {
			t.Fatalf(
				"%s, case #%d: error message mismatch "+
					"(expected '%s', got '%s')",
				t.Name(), n+1, err.Error(), str,
			)
		}
	}
}

func TestUnmarshalError_Error(t *testing.T) {
	cases := []struct {
		prefix string
		src    []byte
		dest   interface{}
	}{
		{"test", []byte("12"), false},
		{"test", []byte("1.2"), 0},
		{"test", []byte("true"), 0.0},
		{"test", []byte("x"), ""},
	}

	for n, c := range cases {
		err := makeUnmarshalError(c.prefix, c.src, c.dest).(UnmarshalError)

		str := fmt.Sprintf(
			"%s: cannot unmarshal '%s' into Go value of type %s",
			err.prefix, err.SrcValue, err.DestType,
		)

		if str != err.Error() {
			t.Fatalf(
				"%s, case #%d: error message mismatch "+
					"(expected '%s', got '%s')",
				t.Name(), n+1, err.Error(), str,
			)
		}
	}
}

func TestConversionError_Error(t *testing.T) {
	cases := []struct {
		prefix string
		src    interface{}
		dest   interface{}
	}{
		{"test", 12, false},
		{"test", 1.2, 0},
		{"test", true, 0.0},
		{"test", 0, ""},
	}

	for n, c := range cases {
		err := makeConversionError(c.prefix, c.src, c.dest).(ConversionError)

		str := fmt.Sprintf(
			"%s: cannot convert %v of type %s into Go value of type %s",
			err.prefix, err.SrcValue, err.SrcType, err.DestType,
		)

		if str != err.Error() {
			t.Fatalf(
				"%s, case #%d: error message mismatch "+
					"(expected '%s', got '%s')",
				t.Name(), n+1, err.Error(), str,
			)
		}
	}
}

func TestParseError_Error(t *testing.T) {
	cases := []struct {
		prefix string
		src    string
		dest   interface{}
	}{
		{"test", "12", false},
		{"test", "1.2", 0},
		{"test", "true", 0.0},
		{"test", "x", ""},
	}

	for n, c := range cases {
		err := makeParseError(c.prefix, c.src, c.dest).(ParseError)

		str := fmt.Sprintf(
			"%s: cannot parse '%s' into Go value of type %s",
			err.prefix, err.SrcString, err.DestType,
		)

		if str != err.Error() {
			t.Fatalf(
				"%s, case #%d: error message mismatch "+
					"(expected '%s', got '%s')",
				t.Name(), n+1, err.Error(), str,
			)
		}
	}
}

func TestTypeError_Error(t *testing.T) {
	cases := []struct {
		prefix   string
		invalid  interface{}
		expected []string
	}{
		{"test", true, []string{"int", "uint"}},
		{"test", 12, []string{"bool"}},
		{"test", 1.2, []string{"string"}},
		{"test", "x", []string{"float64"}},
	}

	for n, c := range cases {
		err := makeTypeError(c.prefix, c.invalid, c.expected...).(TypeError)

		str := fmt.Sprintf(
			"%s: invalid type (expected %s, got %s)", err.prefix,
			strings.Join(err.ExpectedTypes, " or "), err.InvalidType,
		)

		if str != err.Error() {
			t.Fatalf(
				"%s, case #%d: error message mismatch "+
					"(expected '%s', got '%s')",
				t.Name(), n+1, err.Error(), str,
			)
		}
	}
}
