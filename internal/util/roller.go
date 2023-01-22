package util

import "math/rand"

func Roll(numDiceSides int) (result int) {
	return rand.Intn(numDiceSides) + 1
}
