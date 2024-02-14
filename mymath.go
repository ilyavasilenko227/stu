package mymath

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}
func Abs(x float64) float64 {
	return math.Abs(x)
}
func Max(a, b float64) float64 {
	return math.Max(a, b)
}
func Yn(a int, b float64) float64 {
	return math.Yn(a, b)
}
func Ceil(a float64) float64 {
	return math.Ceil(a)
}
func Min(a, b float64) float64 {
	return math.Min(a, b)
}
func Floor(a float64) float64 {
	return math.Floor(a)
}
func jsdhfjshf() {

}
func main() {
	var a float64
	var b float64
	fmt.Println(Sqrt(a))
	fmt.Println(Abs(a))
	fmt.Println(Max(a, b))
	fmt.Println(Yn(7, b))
	fmt.Println(Ceil(a))
	fmt.Println(Min(a, b))
	fmt.Println(Floor(a))
}
