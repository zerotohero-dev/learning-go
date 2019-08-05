package poetry

import (
	"bufio"
	"os"
)

type Line string
type Stanza []Line
type Poem []Stanza

func NewPoem() Poem {
	return Poem{}
}

func LoadPoem(name string) (Poem, error) {
	f, err := os.Open(name)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	scan := bufio.NewScanner(f)

	var s Stanza
	var p Poem

	for scan.Scan() {
		l := scan.Text()

		if l == "" {
			p = append(p, s)
			s = Stanza{}
		} else {
			s = append(s, Line(l))
		}
	}

	p = append(p, s)

	if scan.Err() != nil {
		return nil, scan.Err()
	}

	return p, nil
}

func (p Poem) NumStanzas() int {
	return len(p)
}

func (s Stanza) NumLines() int {
	return len(s)
}
