/*
Package null provides nullable types that are conscious of undefined values
when marshaling or unmarshaling.

JSON

Nullable types in this package can be marshaled to (unmarshaled from) JSON,
as they all implement the json.Marshaler and json.Unmarshaler interfaces.

 var json := struct {
     Mandatory   string      `json:"mandatory"`
     Optional    null.String `json:"optional"`
 }{
     Mandatory: "foo",
     Optional:  bar,
 }

In the example, bar is a String. If bar is valid,
the resulting JSON is going to be:

 {
     "mandatory": "foo",
     "optional": <bar.Str>
 }

otherwise:

 {
     "mandatory": "foo",
     "optional": null
 }

Unmarshaling from JSON works the other way around: JSON value null is
converted to an invalid nullable,
otherwise a compatible JSON type is unmarshaled into the appropriate nullable.

JSON and the omitempty struct tag

Nullable types in this package offer a Ptr() method that is useful to deal with
the omitempty struct tag:

 var json := struct {
     Mandatory  string  `json:"mandatory"`
     Optional   *string `json:"optional,omitempty"`
 }{
     Mandatory: "foo",
     Optional:  bar.Ptr(),
 }

In the example, bar is a String. If bar is valid, Ptr() returns a
pointer to the underlying value, otherwise returns nil. json.Marshal
will recognize nil pointers as empty values,
omitting the associated name from the JSON output.

SQL

Nullable types in this package recognize SQL NULL values and implement the
driver.Valuer and sql.Scanner interfaces.
Suppose we have the following table in our database:

 CREATE TABLE example (
     mandatory varchar(50) primary key not null,
     optional  varchar(50) default null
 );

We may use the following struct that matches the table struct:

 var sql := struct {
     Mandatory   string      `db:"mandatory"`
     Optional    null.String `db:"optional"`
 }{
     Mandatory:  "foo",
     Optional:   bar,
}

In the example, bar is a String. If sql is inserted into the database,
and bar is not valid, a SELECT query will return:

 +-----------+-----------+
 | mandatory | optional  |
 +-----------+-----------+
 | foo       | <bar.Str> |
 +-----------+-----------+

otherwise:

 +-----------+----------+
 | mandatory | optional |
 +-----------+----------+
 | foo       | NULL     |
 +-----------+----------+

It is also possible to scan values from the database. In that case,
if the scanned value corresponds to an SQL NULL,
an invalid nullable is initialized, otherwise,
a compatible SQL value is scanned into the appropriate nullable.

Package flag

Nullable types may also receive values from the command line via the flag
package, as they implement the flag.Value interface:

 var bar null.String

 func init() {
     flag.Var(&bar, "bar", "holds a bar")
 }

 func main() {
     flag.Parse()
     // now bar can be used
 }

bar will be invalid if the command line option "-bar" is not passed or is
empty, otherwise it will be valid and it will hold the content of "-bar".
*/
package null

const (
	// InvalidNullableString is returned by String methods when the nullable is
	// not valid.
	InvalidNullableString = "<invalid>"

	// 32 or 64 bit integers
	intSize = 32 << (^uint(0) >> 63)
)

var (
	jTrue  = []byte("true")
	jFalse = []byte("false")
	jNull  = []byte("null")
)
