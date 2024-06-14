package tasks

import "fmt"

func SumOfSlice(slc []int) (int,error) {
	if len(slc) <= 0 {
		return 0,fmt.Errorf("error: empty slice")
	}
	sum := 0
	for _, v := range slc {
		sum += v
	}
	if sum < 0 {
		return 0,fmt.Errorf("error: negative sum")
	} else {
		return sum,nil
	}
}