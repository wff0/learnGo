package base

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		// return 0, fmt.Errorf("square root of negative number %g", f)
		return 0, errors.New("math - square root of negative number")
	}
	return math.Sqrt(f), nil
}

func div(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到异常：%s\n", r)
		}
	}()

	if b < 0 {
		panic("除数需要大于0")
	}

	fmt.Println("余数为：", a/b)
}

func TestPanic(t *testing.T) {
	// 捕捉内部的异常
	div(10, 0)

	// 捕捉主动的异常
	div(10, -1)
}

/*
程序输出：

捕获到异常：runtime error: integer divide by zero
捕获到异常：除数需要大于0
*/
