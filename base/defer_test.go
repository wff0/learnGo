package base

import (
	"fmt"
	"testing"
	"time"
)

/*
规则一 当defer被声明时，其参数就会被实时解析
规则二 defer执行顺序为先进后出
规则三 defer可以读取有名返回值，也就是可以改变有名返回参数的值。
*/

func TestDefer1(t *testing.T) {
	var i int = 1
	defer fmt.Println("result1=>", func() int { return i * 2 }())
	i++
	defer fmt.Println("result2=>", func() int { return i * 2 }())
}

func TestDefer2(t *testing.T) {
	defer fmt.Print(" !!!\n")
	defer fmt.Print(" world ")
	fmt.Print(" hello ")
}

func TestDefer3(t *testing.T) {
	fmt.Println("=========================")
	fmt.Println("return:", fun1())

	fmt.Println("=========================")
	fmt.Println("return:", fun2())
	fmt.Println("=========================")

	fmt.Println("return:", fun3())
	fmt.Println("=========================")

	fmt.Println("return:", fun4())
}

func fun1() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer2: 2
	}()

	// 规则二 defer执行顺序为先进后出

	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer1: 1
	}()

	// 规则三 defer可以读取有名返回值（函数指定了返回参数名）

	return 100 //这里实际结果为2。如果是return 100呢
}

func fun2() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer2: 2
	}()

	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer1: 1
	}()
	return i
}

func fun3() (r int) {
	t := 5
	defer func() {
		t = t + 5
		fmt.Println(t)
	}()
	return t
}

func fun4() int {
	i := 8
	// 规则一 当defer被声明时，其参数就会被实时解析
	defer func(i int) {
		i = 99
		fmt.Println(i)
	}(i)
	i = 19
	return i
}

/*
为了弄清上述两种情况的区别，我们首先要理解return 返回值的运行机制:
return 并非原子操作，分为赋值，和返回值两步操作
eg1 : 实际上return 执行了两步操作，因为返回值没有命名，所以
return 默认指定了一个返回值（假设为s），首先将i赋值给s,后续
的操作因为是针对i,进行的，所以不会影响s, 此后因为s不会更新，所以
return s 不会改变
相当于：
var i int
s := i
return s
eg2 : 同上，s 就相当于 命名的变量i, 因为所有的操作都是基于
命名变量i(s),返回值也是i, 所以每一次defer操作，都会更新
返回值i
具体defer实现可以查看 https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-defer/
*/

func TestFuncCostTime(t *testing.T) {
	startTime := time.Now()
	defer func() {
		fmt.Println(time.Since(startTime))
	}()
	fmt.Println("start program")
	time.Sleep(5 * time.Second)
	fmt.Println("finish program")
}
