package mat

import (
	"fmt"
)

//
// Debugging utilities
//
func toString(a Matrix) string {
	s := "\n"
	for i := 0; i < a.Rows(); i++ {
		for j := 0; j < a.Cols(); j++ {
			s += fmt.Sprintf("%8.3f", a.Get(i, j))
		}
		s += fmt.Sprintf("\n")
	}
	return s
}

func (a *FullMat) String() string {
	return toString(a)
}

func (a *FullMat) Print() {
	fmt.Printf("%s", a.String())
}

func (a *SubMat) String() string {
	return toString(a)
}

func (a *SubMat) Print() {
	fmt.Printf("%s", a.String())
}
