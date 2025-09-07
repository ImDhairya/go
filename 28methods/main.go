package main

import (
	"fmt"
	"math"
)

// custom type and we can define method on this type also

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// if you make any change in the value receiver it will not be reflected on the variable

// but if you make a change in the pointer receiver (*) then it will also change the original variable

type Verted struct {
	X, Y float64
}

func (v Verted) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Verted) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f

}

func main() {
	v := Verted{3, 4}

	f := MyFloat(-math.Sqrt2)
	fmt.Println(v.Abs())
	fmt.Println("Hello ")
	v.Scale(32)

	fmt.Println(v, f)

	// fmt.Println(f.Abs())
}
