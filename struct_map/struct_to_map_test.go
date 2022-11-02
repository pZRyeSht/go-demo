package struct_map

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStructSliceToMapSlice(t *testing.T) {
	type Student struct {
		Name string `json:"name"`
		Age  uint   `json:"age"`
	}
	type Teacher struct {
		Name string `json:"name"`
		Gender  uint   `json:"gender"`
	}
	var people []interface{}
	// 切片添加Student元素
	stu1 := Student{
		Name: "esc",
		Age:  18,
	}
	stu2 := Student{
		Name: "alice",
		Age:  20,
	}
	people = append(people, stu1)
	people = append(people, stu2)
	
	// 切片添加Teacher元素
	tech := Teacher{
		Name: "esca",
		Gender:  0,
	}
	people = append(people, tech)
	slice, err := StructSliceToMapSlice(people)
	assert.Nil(t, err)
	fmt.Println(slice)
}

func TestStructToMap(t *testing.T) {
	type Student struct {
		Name string `json:"name"`
		Age  uint   `json:"age"`
	}
	stu1 := Student{
		Name: "esc",
		Age:  18,
	}
	mp1, err := StructToMap(stu1)
	assert.Nil(t, err)
	fmt.Println(mp1)
	stu2 := &Student{
		Name: "alice",
		Age:  20,
	}
	mp2, err := StructToMap(stu2)
	assert.Nil(t, err)
	fmt.Println(mp2)
}
