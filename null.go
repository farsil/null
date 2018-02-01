package null

import (
	"time"
)

const (
	sInvalid = "<invalid>"
	intSize  = 32 << (^uint(0) >> 63)
)

var tEmpty time.Time

var (
	jTrue  = []byte("true")
	jFalse = []byte("false")
	jNull  = []byte("null")
)
