package testDemo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect(test *testing.T) {
	var a int = 50
	v := reflect.ValueOf(a) // 返回Value类型对象，值为50
	t := reflect.TypeOf(a)  // 返回Type类型对象，值为int
	fmt.Println(v, t, v.Type(), t.Kind())

	var b [5]int = [5]int{5, 6, 7, 8}
	fmt.Println(reflect.TypeOf(b), reflect.TypeOf(b).Kind(), reflect.TypeOf(b).Elem()) // [5]int array int

	type Student struct {
		name string
	}

	var Pupil Student
	p := reflect.ValueOf(Pupil) // 使用ValueOf()获取到结构体的Value对象

	fmt.Println(p.Type()) // 输出:Student
	fmt.Println(p.Kind()) // 输出:struct
}

func TestSet(test *testing.T) {
	type Student struct {
		name string
		Age  int
	}

	var a int = 50
	v := reflect.ValueOf(a) // 返回Value类型对象，值为50
	t := reflect.TypeOf(a)  // 返回Type类型对象，值为int
	fmt.Println(v, t, v.Type(), t.Kind(), reflect.ValueOf(&a).Elem())
	seta := reflect.ValueOf(&a).Elem() // 这样才能让seta保存a的值
	fmt.Println(seta, seta.CanSet())
	seta.SetInt(1000)
	fmt.Println(seta)

	var b [5]int = [5]int{5, 6, 7, 8}
	fmt.Println(reflect.TypeOf(b), reflect.TypeOf(b).Kind(), reflect.TypeOf(b).Elem())

	var Pupil Student = Student{"joke", 18}
	p := reflect.ValueOf(Pupil) // 使用ValueOf()获取到结构体的Value对象

	fmt.Println(p.Type()) // 输出:Student
	fmt.Println(p.Kind()) // 输出:struct

	setStudent := reflect.ValueOf(&Pupil).Elem()
	//setStudent.Field(0).SetString("Mike") // 未导出字段，不能修改，panic会发生
	setStudent.Field(1).SetInt(19)
	fmt.Println(setStudent)
}

func TestTags(t *testing.T) {
	type Student struct {
		name string
		Age  int `json:"years"`
	}

	var Pupil Student = Student{"joke", 18}
	setStudent := reflect.ValueOf(&Pupil).Elem()

	sSAge, _ := setStudent.Type().FieldByName("Age")
	fmt.Println(sSAge.Tag.Get("json")) // years
}
