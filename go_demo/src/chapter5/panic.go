package main

import (
	"fmt"
	"math"
)

/*
test panic.
*/

func main() {
	i, err := IntFromInt64(1000011231231231)
	if err != nil {
		fmt.Printf("#####  err %t\n", err)
		return
	}
	fmt.Printf("i = %d , is err is %d", i, err)
}

func ConvertInt64ToInt(x int64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		return int(x)
	}
	panic(fmt.Sprintf("%d is out of the int 32 range", x))
}

func IntFromInt64(x int64) (i int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	i = ConvertInt64ToInt(x)
	return i, nil
}

/*func Test_Caller_G(t *testing.T) {
	defaultOfCaller := func() {
		println("the defer of Test_Caller_G")
	}
}*/
