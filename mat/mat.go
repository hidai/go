package mat

type Matrix interface {
	Get(i, j int) float64
	Set(i, j int, v float64) Matrix
	Pointer(i, j int) *float64

	Size() (int, int)
	Rows() int
	Cols() int
	NumEl() int

	Sub(i, j, m, n int) Matrix
	SubRows(j, n int) Matrix
	SubCols(i, m int) Matrix

	String() string
	Print()

	Add(a Matrix) Matrix
	AddScalar(a float64) Matrix
	AddAssign(a Matrix) Matrix
	AddAssignScalar(a float64) Matrix
}

type FullMat struct {
	m    int
	n    int
	data []float64
}

type SubMat struct {
	FullMat
	si int
	sj int
	sm int
	sn int
}

var _1 Matrix = &FullMat{} // Type check
var _2 Matrix = &SubMat{}  // Type check

//
// Constructor
//
func New(m, n int) *FullMat {
	return NewWithData(m, n, make([]float64, m*n))
}

func NewWithData(m, n int, data []float64) *FullMat {
	if len(data) != m*n {
		panic("Invalid size data.")
	}
	return &FullMat{m, n, data}
}

//
// Sub matrix
//
func (a *FullMat) Sub(i, j, m, n int) Matrix {
	if i < 0 || j < 0 || a.m <= i || a.n <= j ||
		m < 1 || n < 1 || a.m < i+m || a.n < j+n {
		panic("Invalid sub matrix size")
	}
	return &SubMat{*a, i, j, m, n}
}

func (a *SubMat) Sub(i, j, m, n int) Matrix {
	i += a.si
	j += a.sj
	if i < 0 || j < 0 || a.m <= i || a.n <= j ||
		m < 1 || n < 1 || a.m < i+m || a.n < j+n {
		panic("Invalid sub matrix size")
	}
	return &SubMat{a.FullMat, i, j, m, n}
}

func (a *FullMat) SubRows(j, n int) Matrix {
	return a.Sub(0, j, a.Rows(), n)
}

func (a *FullMat) SubCols(i, m int) Matrix {
	return a.Sub(i, 0, m, a.Cols())
}

func (a *SubMat) SubRows(j, n int) Matrix {
	return a.Sub(0, j, a.Rows(), n)
}

func (a *SubMat) SubCols(i, m int) Matrix {
	return a.Sub(i, 0, m, a.Cols())
}

//
// Element accessers
//
func (a *FullMat) Get(i, j int) float64 {
	return a.data[i+j*a.m]
}

func (a *FullMat) Set(i, j int, v float64) Matrix {
	a.data[i+j*a.m] = v
	return a
}

func (a *FullMat) Pointer(i, j int) *float64 {
	return &a.data[i+j*a.m]
}

func (a *SubMat) Get(i, j int) float64 {
	i += a.si
	j += a.sj
	return a.data[i+j*a.m]
}

func (a *SubMat) Set(i, j int, v float64) Matrix {
	i += a.si
	j += a.sj
	a.data[i+j*a.m] = v
	return a
}

func (a *SubMat) Pointer(i, j int) *float64 {
	i += a.si
	j += a.sj
	return &a.data[i+j*a.m]
}

//
// Size
//
func (a *FullMat) Size() (int, int) {
	return a.m, a.n
}

func (a *FullMat) Rows() int {
	return a.m
}

func (a *FullMat) Cols() int {
	return a.n
}

func (a *FullMat) NumEl() int {
	return a.m * a.n
}

func (a *SubMat) Size() (int, int) {
	return a.sm, a.sn
}

func (a *SubMat) Rows() int {
	return a.sm
}

func (a *SubMat) Cols() int {
	return a.sn
}

func (a *SubMat) NumEl() int {
	return a.sm * a.sn
}
