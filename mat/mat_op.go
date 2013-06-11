package mat

func sameSize(a Matrix, b Matrix) bool {
	return a.Rows() == b.Rows() && a.Cols() == b.Cols()
}

func sameSizeOrDie(a Matrix, b Matrix) {
	if !sameSize(a, b) {
		panic("Invalid size")
	}
}

//
// +
//
func add(a Matrix, b Matrix) Matrix {
	sameSizeOrDie(a, b)
	c := New(a.Size())
	for j := 0; j < a.Cols(); j++ {
		for i := 0; i < a.Rows(); i++ {
			c.Set(i, j, a.Get(i, j)+b.Get(i, j))
		}
	}
	return c
}

func addScalar(a Matrix, b float64) Matrix {
	c := New(a.Size())
	for j := 0; j < a.Cols(); j++ {
		for i := 0; i < a.Rows(); i++ {
			c.Set(i, j, a.Get(i, j)+b)
		}
	}
	return c
}

func (a *FullMat) Add(b Matrix) Matrix {
	return add(a, b)
}

func (a *FullMat) AddScalar(b float64) Matrix {
	return addScalar(a, b)
}

func (a *SubMat) Add(b Matrix) Matrix {
	return add(a, b)
}

func (a *SubMat) AddScalar(b float64) Matrix {
	return addScalar(a, b)
}

//
// +=
//
func addAssign(a, b Matrix) Matrix {
	sameSizeOrDie(a, b)
	for j := 0; j < a.Cols(); j++ {
		for i := 0; i < a.Rows(); i++ {
			*a.Pointer(i, j) += b.Get(i, j)
		}
	}
	return a
}

func addAssignScalar(a Matrix, b float64) Matrix {
	for j := 0; j < a.Cols(); j++ {
		for i := 0; i < a.Rows(); i++ {
			*a.Pointer(i, j) += b
		}
	}
	return a
}

func (a *FullMat) AddAssign(b Matrix) Matrix {
	return addAssign(a, b)
}

func (a *FullMat) AddAssignScalar(b float64) Matrix {
	return addAssignScalar(a, b)
}

func (a *SubMat) AddAssign(b Matrix) Matrix {
	return addAssign(a, b)
}

func (a *SubMat) AddAssignScalar(b float64) Matrix {
	return addAssignScalar(a, b)
}
