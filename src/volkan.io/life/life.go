package life

import (
	"bytes"
	"math/rand"
)

type Life struct {
	a, b *Field
	w, h int
}

func NewLife(w, h int) *Life {
	a := NewField(w, h)

	for i := 0; i < (w * h / 4); i++ {
		a.Set(rand.Intn(w), rand.Intn(h), true)
	}

	return &Life{
		a: a, b: NewField(w, h),
		w: w, h: h,
	}
}

func (l *Life) Step() {
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			l.b.Set(x, y, l.a.Next(x, y))
		}
	}

	l.a, l.b = l.b, l.a
}

func (l *Life) String() string {
	var buf bytes.Buffer

	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			b := byte(' ')

			if l.a.Alive(x, y) {
				b = '*'
			}

			buf.WriteByte(b)
		}

		buf.WriteByte('\n')
	}

	return buf.String()
}
