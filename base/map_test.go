package base

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	// 声明但未初始化map，此时是map的零值状态
	//map1 := make(map[string]string, 5)

	//map2 := make(map[string]string)

	// 创建了初始化了一个空的的map，这个时候没有任何元素
	//map3 := map[string]string{}

	// map中有三个值
	map4 := map[string]string{"a": "1", "b": "2", "c": "3"}

	// 删除key
	delete(map4, "a")

	// 一般用来判断key是否存在
	if _, ok := map4["a"]; !ok {
		fmt.Println("no entry")
	}
}

// 在"range"语句中生成的数据的值是真实集合元素的拷贝，它们不是原有元素的引用。
// 这意味着更新这些值将不会修改原来的数据。同时也意味着使用这些值的地址将不会得到原有数据的指针。
func TestRange(t *testing.T) {
	data := []int{1, 2, 3}
	for _, v := range data {
		v *= 10 // 通常数据项不会改变
	}
	fmt.Println("data:", data) // 程序输出: [1 2 3]
	// 如果你需要更新原有集合中的数据，使用索引操作符来获得数据。
	for i, _ := range data {
		data[i] *= 10
	}
	fmt.Println("data:", data) // 程序输出 data: [10 20 30]
}
