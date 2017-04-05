package reflect

import (
	"fmt"
)

func Describe(i interface{}) {
	fmt.Printf("(%v, %T)", i, i)
}
