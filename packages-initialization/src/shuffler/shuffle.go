package shuffler

import (
	"math/rand"
)

var (
	MEANING int
)

// Go will make sure that the main program will run only
// when the init() functions of all the imported packages are run.
func init() {
	MEANING = 42
}

type Shuffleable interface {
	Len() int
	Swap(i, j int)
}

type WeightedShuffleable interface {
	Shuffleable
	Weight(i int) int
}

func Shuffle(s Shuffleable) {
	for i := 0; i < s.Len(); i++ {
		j := rand.Intn(s.Len() - i)
		s.Swap(i, j)
	}
}
