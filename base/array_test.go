package base

import (
	"testing"
)

// 一维数组的初始化方式
func TestArrayInit(t *testing.T) {
	//var arrAge  = [5]int{18, 20, 15, 22, 16}
	//var arrName = [5]string{3: "Chris", 4: "Ron"} //指定索引位置初始化
	//// {"","","","Chris","Ron"}
	//var arrCount = [4]int{500, 2: 100} //指定索引位置初始化 {500,0,100,0}
	//var arrLazy = [...]int{5, 6, 7, 8, 22} //数组长度初始化时根据元素多少确定
	//var arrPack = [...]int{10, 5: 100} //指定索引位置初始化，数组长度与此有关 {10,0,0,0,0,100}
	//var arrRoom [20]int
	//var arrBed = new([20]int)
}

// 如数组元素类型支持”==，!=”操作符，那么数组也支持此操作，
// 但如果数组类型不一样则不支持（需要长度和数据类型一致，否则编译不通过）。如：
func TestArrayJudge(t *testing.T) {
	var arrRoom [20]int
	var arrBed [20]int

	println(arrRoom == arrBed) //true
}
