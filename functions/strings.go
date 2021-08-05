package functions

import (
	"math/rand"
	"strconv"
)

func RandomString() string {
	return strconv.Itoa(rand.Int())
}
