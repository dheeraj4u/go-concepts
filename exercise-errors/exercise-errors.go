package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number:", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for {
		y := z
		z -= (z*z - x) / (2 * z)
		sizey := float64(len(fmt.Sprintf("%f", y)))
		sizez := float64(len(fmt.Sprintf("%f", z)))
		//fmt.Println(y, z, sizey, sizez, y*math.Pow(10, sizey-1), z*math.Pow(10, sizez-1), int(y*math.Pow(10, sizey-1)), int(z*math.Pow(10, sizez-1)))
		if int(y*math.Pow(10, sizey-1)) == int(z*math.Pow(10, sizez-1)) {
			break
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
