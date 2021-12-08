package base

import (
	"fmt"
	"testing"
)

// 如果切片取值时索引值大于长度会导致panic错误发生，即使容量远远大于长度也没有用
func TestSlice(test *testing.T) {
	sli := make([]int, 5, 10)
	fmt.Printf("切片sli长度和容量：%d, %d\n", len(sli), cap(sli))
	fmt.Println(sli)
	newsli := sli[:cap(sli)]
	fmt.Println(newsli)

	var x = []int{2, 3, 5, 7, 11}
	fmt.Printf("切片x长度和容量：%d, %d\n", len(x), cap(x))

	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:3:5] // a[low : high : max]  max-low的结果表示容量  high-low为长度
	fmt.Printf("切片t长度和容量：%d, %d\n", len(t), cap(t))

	// fmt.Println(t[2]) // panic ，索引不能超过切片的长度
}

func get() []byte {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0]) // 显示: 10000 10000 数组首字节地址
	res := make([]byte, 3)
	// copy(new, old)
	copy(res, raw[:3]) // 利用copy 函数复制，raw 可被GC释放
	return res
}

// 当我们在一个切片基础上重新划分一个切片时，新的切片会继续引用原有切片的数组。
// 如果你忘了这个行为的话，在你的应用分配大量临时的切片用于创建新的切片来
// 引用原有数据的一小部分时，会导致难以预期的内存使用。
func TestCopy(t *testing.T) {
	data := get()
	fmt.Println(len(data), cap(data), &data[0]) // 显示: 3 3 数组首字节地址
}

// append()函数将 0 个或多个具有相同类型S的元素追加到切片s后面并且返回新的切片；
// 追加的元素必须和原切片的元素同类型。如果s的容量不足以存储新增元素，
// append()会分配新的切片来保证已有切片元素和新增元素的存储。
// 因此，append()函数返回的切片可能已经指向一个不同的相关数组了。append()函数总是返回成功，除非系统内存耗尽了。
// append()函数操作如果导致分配新的切片来保证已有切片元素和新增元素的存储，
// 也就是返回的切片可能已经指向一个不同的相关数组了，那么新的切片
// 已经和原来切片没有任何关系，即使修改了数据也不会同步。
func TestAppend(t *testing.T) {
	s0 := []int{0, 0}
	s1 := append(s0, 2)              // append 单个元素     s1 == []int{0, 0, 2}
	s2 := append(s1, 3, 5, 7)        // append 多个元素    s2 == []int{0, 0, 2, 3, 5, 7}
	s3 := append(s2, s0...)          // append 一个切片     s3 == []int{0, 0, 2, 3, 5, 7, 0, 0}
	s4 := append(s3[3:6], s3[2:]...) // append 切片片段    s4 == []int{3, 5, 7, 2, 3, 5, 7, 0, 0}
	fmt.Println(s4)
}

// 多个切片可以引用同一个底层数组。在某些情况下，在一个切片中添加新的数据，
// 在原有数组无法保持更多新的数据时，将导致分配一个新的数组。而现在其他的切片还指向老的数组（和老的数据）。
func TestStaleSlices(t *testing.T) {
	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1) // 输出 3 3 [1 2 3]
	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2) // 输出 2 2 [2 3]
	for i := range s2 {
		s2[i] += 20
	}
	// s2的修改会影响到数组数据，s1输出新数据
	fmt.Println(s1) // 输出 [1 22 23]
	fmt.Println(s2) // 输出 [22 23]

	s2 = append(s2, 4) // append  s2容量为2，这个操作导致了切片 s2扩容，会生成新的底层数组。

	for i := range s2 {
		s2[i] += 10
	}
	// s1 的数据现在是老数据，而s2扩容了，复制数据到了新数组，他们的底层数组已经不是同一个了。
	fmt.Println(len(s1), cap(s1), s1) // 输出3 3 [1 22 23]
	fmt.Println(len(s2), cap(s2), s2) // 输出3 4 [32 33 14]
}
