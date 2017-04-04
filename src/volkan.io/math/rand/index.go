package rand

import (
    "time"
    "math/rand"
)

var seed int64

func Seed() {
    seed = time.Now().UTC().UnixNano()

    rand.Seed(seed)
}

func GenerateInt(max int) int {
    if seed == 0 {
        panic("Seed() me first!")
    }

    return rand.Intn(max)
}