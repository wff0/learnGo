package testDemo

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestTimestamp(T *testing.T) {
	now := time.Now()
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
}

// 测试变量作用域
func Test1(t *testing.T) {
	if a := 1; false {
	} else if b := 2; false {
	} else if c := 3; false {
	} else {
		println(a, b, c)
	}
}

// 生成真正的随机数不是伪随机
func TestRand(t *testing.T) {
	//不加随机数种子的话每次生成一样
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(10))
	}
}
