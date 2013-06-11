package mat

import (
	"testing"
)

func TestMatrix(t *testing.T) {
	a := New(3, 3)
	a.Print()
	b := a.AddScalar(77)
	b.Print()

	c := NewWithData(3, 3, []float64{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	})
	c.Print()
	c.Sub(0, 0, 2, 2).Print()
	c.Sub(1, 1, 2, 2).Print()
	c.SubRows(1, 1).Print()
	c.SubCols(1, 1).Print()

	c.Print()
	c.Sub(0, 0, 2, 2).AddScalar(100).Print()
	c.Print()
	c.Sub(0, 0, 2, 2).Add(c.Sub(1, 1, 2, 2)).Print()
	c.Print()
}
