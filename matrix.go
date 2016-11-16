package gmatrix

import (
	"bytes"
	"strconv"
)

type Matrix struct {
	rows, columns int       // the number of rows and columns.
	data          []float32 // the contents of the matrix as one long slice.
}

// Set lets you define the value of a matrix at the given row and
// column.

func (A *Matrix) Set(r int, c int, val float32) {
	A.data[findIndex(r, c, A)] = val
}

// Get retrieves the contents of the matrix at the row and column.

func (A *Matrix) Get(r, c int) float32 {
	return A.data[findIndex(r, c, A)]
}

// Column returns a slice that represents a column from the matrix.
// This works by examining each row, and adding the nth element of
// each to the column slice.

func (A *Matrix) Column(n int) []float32 {
	col := make([]float32, A.rows)
	for i := 1; i <= A.rows; i++ {
		col[i - 1] = A.Row(i)[n - 1]
	}
	return col
}

// Row returns a slice that represents a row from the matrix.

func (A *Matrix) Row(n int) []float32 {
	return A.data[findIndex(n, 1, A):findIndex(n, A.columns, A) + 1]
}

// Multiply multiplies two matrices together and return the resulting matrix.
// For each element of the result matrix, we get the dot product of the
// corresponding row from matrix A and column from matrix B.

func Multiply(A, B *Matrix) *Matrix {
	C := Zeros(A.rows, B.columns)
	for r := 1; r <= C.rows; r++ {
		A_row := A.Row(r)
		for c := 1; c <= C.columns; c++ {
			B_col := B.Column(c)
			C.Set(r, c, dotProduct(A_row, B_col))
		}
	}
	return C
}

// Add adds two matrices together and returns the resulting matrix.  To do
// this, we just add together the corresponding elements from each matrix.

func Add(A, B *Matrix) *Matrix {
	C := Zeros(A.rows, A.columns)
	for r := 1; r <= A.rows; r++ {
		for c := 1; c <= A.columns; c++ {
			C.Set(r, c, A.Get(r, c) + B.Get(r, c))
		}
	}
	return C
}

// Identity creates an identity matrix with n rows and n columns.  When you
// multiply any matrix by its corresponding identity matrix, you get the
// original matrix.  The identity matrix looks like a zero-filled matrix with
// a diagonal line of one's starting at the upper left.

func Identity(n int) *Matrix {
	A := Zeros(n, n)
	for i := 0; i < len(A.data); i += (n + 1) {
		A.data[i] = 1
	}
	return A
}

// Zeros creates an r x c sized matrix that's filled with zeros.  The initial
// state of an int is 0, so we don't have to do any initialization.

func Zeros(r, c int) *Matrix {
	return &Matrix{r, c, make([]float32, r * c)}
}

// New creates an r x c sized matrix that is filled with the provided data.
// The matrix data is represented as one long slice.

func New(r, c int, data []float32) *Matrix {
	if len(data) != r * c {
		panic("[]int data provided to matrix.New is great than the provided capacity of the matrix!'")
	}
	A := Zeros(r, c)
	A.data = data
	return A
}

// findIndex takes a row and column and returns the corresponding index
// from the underlying data slice.

func findIndex(r, c int, A *Matrix) int {
	return (r - 1) * A.columns + c - 1
}

// dotProduct calculates the algebraic dot product of two slices.  This is just
// the sum  of the products of corresponding elements in the slices.  We use
// this when we multiply matrices together.

func dotProduct(a, b []float32) float32 {
	var total float32
	for i := 0; i < len(a); i++ {
		total += a[i] * b[i]
	}
	return total
}

func (m *Matrix)Scalar(a float32) {
	for r := 1; r <= m.rows; r++ {
		for c := 1; c <= m.columns; c++ {
			m.Set(r, c, m.Get(r, c) * a)
		}
	}
}
func (m *Matrix)Trans() {

}

func (m *Matrix)String() string {
	var bf bytes.Buffer
	bf.WriteByte('[')
	for r := 1; r <= m.rows; r++ {
		for c := 1; c <= m.columns; c++ {
			bf.WriteString(strconv.FormatFloat(float64(m.data[findIndex(r, c, m)]), 'f', -1, 32))
			if c == m.columns {
				if r == m.rows {
					bf.WriteByte(']')
				} else {
					bf.WriteString(",\n ")
				}
			} else {
				bf.WriteString(", ")
			}
		}
	}
	return bf.String()
}
