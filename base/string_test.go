package base

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// 内部使用 []byte 实现，不像直接运算符这种会产生很多临时的字符串，
// 但是内部的逻辑比较复杂，有很多额外的判断，还用到了 interface，所以性能一般。
func TestFmtSprintf(t *testing.T) {
	str := fmt.Sprintf("%d:%s", 2018, "年")
	fmt.Println(str)
}

// Join会先根据字符串数组的内容，计算出一个拼接之后的长度，
// 然后申请对应大小的内存，一个一个字符串填入，
// 在已有一个数组的情况下，这种效率会很高，但是本来没有，去构造这个数据的代价也不小。
func TestStringJoin(t *testing.T) {
	str1 := strings.Join([]string{"hello", "world"}, ", ")
	fmt.Println(str1)
}

// 这个比较理想，可以当成可变字符使用，对内存的增长也有优化，
// 如果能预估字符串的长度，还可以用 buffer.Grow() 接口来设置 capacity。
func TestBuffer(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("hello")
	buffer.WriteString(", ")
	buffer.WriteString("world")
	fmt.Println(buffer.String())
}

// strings.Builder 内部通过 slice 来保存和管理内容。
// slice 内部则是通过一个指针指向实际保存内容的数组。
// strings.Builder 同样也提供了 Grow() 来支持预定义容量。
// 当我们可以预定义我们需要使用的容量时，strings.Builder 就能避免扩容而创建新的 slice 了。
// strings.Builder是非线程安全，性能上和 bytes.Buffer 相差无几。
func TestStringBuilder(t *testing.T) {
	var builder strings.Builder
	builder.Grow(6)
	builder.WriteString("ABC")
	builder.WriteString("DEF")
	fmt.Println(builder.String())
}

/*
标准库中有四个包对字符串处理尤为重要：bytes、strings、strconv和unicode包。
strings包提供了许多如字符串的查询、替换、比较、截断、拆分和合并等功能。
bytes包也提供了很多类似功能的函数，但是针对和字符串有着相同结构的[]byte类型。
因为字符串是只读的，因此逐步构建字符串会导致很多分配和复制。在这种情况下，使用bytes.Buffer类型将会更有效。
strconv包提供了布尔型、整型数、浮点数和对应字符串的相互转换，还提供了双引号转义相关的转换。
unicode包提供了IsDigit、IsLetter、IsUpper和IsLower等类似功能，它们用于给字符分类。
strings 包提供了很多操作字符串的简单函数，通常一般的字符串操作需求都可以在这个包中找到。下面简单举几个例子：
判断是否以某字符串打头/结尾 strings.HasPrefix(s, prefix string) bool strings.HasSuffix(s, suffix string) bool
字符串分割 strings.Split(s, sep string) []string
返回子串索引 strings.Index(s, substr string) int strings.LastIndex 最后一个匹配索
字符串连接 strings.Join(a []string, sep string) string 另外可以直接使用“+”来连接两个字符串
字符串替换 strings.Replace(s, old, new string, n int) string
字符串转化为大小写 strings.ToUpper(s string) string strings.ToLower(s string) string
统计某个字符在字符串出现的次数 strings.Count(s, substr string) int
判断字符串的包含关系 strings.Contains(s, substr string) bool
*/
