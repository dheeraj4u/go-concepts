package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
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
	return z
}

func main() {
	fmt.Print(Sqrt(112222))
}
