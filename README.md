# null
```
import "github.com/farsil/null"
```
`null` is a [Go](https://golang.org/) package which aims to provide a 
reasonable way to deal with undefined values. It defines wrapper types that 
hold a boolean flag to specify whether the wrapped type is valid or not. The 
validity flag is used during marshaling (unmarshaling) to decide whether an 
undefined value should be marshaled (unmarshaled) in lieu of the underlying 
value. The undefined value varies according to the specific serialization 
interface:

- With JSON, the undefined value is the JSON string `null`
- With SQL, the undefined value is the SQL value `NULL`
- With `flag`, the undefined value is `""`
- With the default text encoding, the undefined value is an empty `[]byte`

Types in the package implement all the necessary interfaces to be used with the 
`flag` package, to handle JSON and SQL `NULL` values, so it can be used in 
place of the `sql.Nullxxx` types. The interfaces implemented by each type are:

- `TextMarshaler` from `encoding`
- `TextUnmarshaler` from `encoding`
- `Marshaler` from `encoding/json`
- `Unmarshaler` from `encoding/json`
- `Value` from `flag`
- `Scanner` from `database/sql`
- `Valuer` from `database/sql/driver`
- `Stringer` from `fmt` (as part of the `Value` interface from `flag`)

## Installation
Installation is as simple as invoking:
```
$ go get github.com/farsil/null
```

## Data Types
The package defines the following data types:

- `null.String` which wraps a `string`
- `null.Bool` which wraps a `bool`
- `null.Int` which wraps an `int`
- `null.Uint` which wraps an `uint`
- `null.Float64` which wraps a `float64`
- `null.Time` which wraps a `time.Time`

Note that JSON does not define a standard datetime representation. In this 
package, a `null.Time` object is represented as an 
[RFC3339](https://tools.ietf.org/html/rfc3339) string when a `null.Time` object 
is marshaled, and an RFC3339 string is parsed when unmarshaling from JSON.

## Example
```
package main

import "github.com/farsil/null"

var data = []byte(`{
    "title": "The foo does a bar"
}`)

type BookTitle struct {
	Title    null.String `json:"title"`
	Subtitle null.String `json:"subtitle"`
}

func main() {
	var t Title
	if err := json.Unmarshal(data, &t); err != nil {
		// process error
	}

	if !t.Title.Valid {
		fmt.Println("Book has no title")
		return
	}

	fmt.Println("Book title:", t.Title.Str)

	if t.Subtitle.Valid {
		fmt.Println("Book subtitle:", t.Subtitle.Str)
	} else {
		fmt.Println("Book has no subtitle")
	}
}
```
Output
```
Book title: 'The foo does a bar'
Book has no subtitle
```

## Documentation
The documentation is available at the address 
https://godoc.org/github.com/farsil/null.

## Credits
This repository is heavily inspired by https://github.com/guregu/null. I 
originally wanted to create a fork of that repository, but since there are 
critical incompatibilities of interfaces and behavior, I decided to start from 
scratch.
