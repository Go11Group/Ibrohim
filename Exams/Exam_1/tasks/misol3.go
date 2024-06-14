package tasks

import "math"

func Abs(n chan int) {
	n <- int(math.Abs(float64(<-n)))
}