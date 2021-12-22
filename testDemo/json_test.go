package testDemo

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
)

type Human struct {
	name   string `json:"name"` // 姓名
	Gender string `json:"s"`    // 性别，性别的tag表明在json中为s字段
	Age    int    `json:"Age"`  // 年龄
	Lesson
}

type Lesson struct {
	Lessons []string `json:"lessons"`
}

func TestJsonEnAndDe(t *testing.T) {
	// json数据的字符串
	jsonStr := `{"Age": 18,"name": "Jim" ,"s": "男",
	"lessons":["English","History"],"Room":201,"n":null,"b":false}`
	r := strings.NewReader(jsonStr)
	h := &Human{}
	err := json.NewDecoder(r).Decode(h)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(h)

	f, _ := os.Create("./a.json")
	err = json.NewEncoder(f).Encode(h)
	if err != nil {
		fmt.Println(err)
	}
}
