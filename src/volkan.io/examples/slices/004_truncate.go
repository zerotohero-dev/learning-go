package main

import (
	"bytes"
	"fmt"
)

type path []byte

func (p *path) TruncateAtFinalSlash() *path {
	i := bytes.LastIndex(*p, []byte("/"))

	if i > 0 {
		*p = (*p)[0:i]
	}

	return p
}

func (p path) TruncateAtFinalSlashByVal() path {
	i := bytes.LastIndex(p, []byte("/"))

	if i > 0 {
		p = (p)[0:i]
	}

	return p
}

func main() {
	pathName := path("/usr/bin/sudo")

	res1 := pathName.TruncateAtFinalSlashByVal()

	fmt.Printf("%s\n", pathName)
	fmt.Printf("%s\n", res1)

	// same as (&pathName).TruncateAtFinalSlah()
	res2 := pathName.TruncateAtFinalSlash()

	fmt.Printf("%s\n", pathName)
	fmt.Printf("%s\n", res2)
}
