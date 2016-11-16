package gmatrix

import (
	"testing"
	"fmt"
)

func TestNew(t *testing.T) {
	m := &Matrix{
		rows:2,
		columns:3,
		data:[]float32{1, 2, 3, 4, 5, 6},
	}
	n := Zeros(3, 2)
	fmt.Println(m.String())
	m.Set(1, 2, 3.3)
	fmt.Println(m.String())
	fmt.Println(m.Get(1, 2))
	fmt.Println(Multiply(m, n))

	i := Identity(3)
	fmt.Println(Multiply(m, i))
}
