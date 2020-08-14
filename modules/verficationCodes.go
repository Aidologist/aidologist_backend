package modules

import (
	"math/rand"
	"time"
)

func NumberOnlyVerficationCode(digit int) string {
	var base int32 = int32(10 ^ (digit + 1))
	var code int32 = rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(base)
	return string(code)
}