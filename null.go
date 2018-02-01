/*
Package null provides nullable types that are conscious of undefined values
when marshaling or unmarshaling.

omitempty

Nullable types in this package offer a Ptr() method that is useful to deal with
JSON's omitempty struct tag:

	var json := struct {
		Mandatory   string  `json:"mandatory"`
		Optional    *string `json:"optional,omitempty"`
	}{
		Mandatory:  "foo",
		Optional:   bar.Ptr(),
	}

In the example, bar is a nullable string. If bar is valid, Ptr() returns a
pointer to the underlying value, otherwise returns nil. json.
Marshal will recognize nil pointers as empty values,
omitting the relative name from the JSON output.
*/
package null

const (
	// Returned by String methods when the nullable is not valid.
	InvalidNullableStr = "<invalid>"

	// 32 or 64 bit integers
	intSize = 32 << (^uint(0) >> 63)
)

var (
	jTrue  = []byte("true")
	jFalse = []byte("false")
	jNull  = []byte("null")
)
