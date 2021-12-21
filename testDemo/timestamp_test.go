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

	n1, _ := time.ParseInLocation("2006-01-02", "2019-01-09", time.Local)
	n2, _ := time.ParseInLocation("2006-01-02", "2019-01-01", time.Local)
	fmt.Println((n1.Unix() - n2.Unix()) / (60 * 60 * 24))
	fmt.Println(n1.Sub(n2).Hours() / 24)
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
