package field

type Field struct {
	s    [][]bool
	w, h int
}

func New(w, h int) *Field {
	s := make([][]bool, h)

	for i := range s {
		s[i] = make([]bool, w)
	}

	return &Field{s: s, w: w, h: h}
}

func (f *Field) Set(x, y int, b bool) {
	f.s[y][x] = b
}

func (f *Field) Alive(x, y int) bool {
	return f.s[(y+f.h)%f.h][(x+f.w)%f.w]
}

func (f *Field) Next(x, y int) bool {
	alive := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if f.Alive(x+i, y+j) {
				alive++
			}
		}
	}

	return alive == 3 || alive == 2 && f.Alive(x, y)
}
