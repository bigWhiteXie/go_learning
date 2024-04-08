package main

import (
	"fmt"
	"reflect"
	"testing"
)

type student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestStruct(t *testing.T) {
	m := make(map[string]*student)
	stus := []student{
		{Name: "小王子", Age: 18},
		{Name: "娜扎", Age: 23},
		{Name: "大王八", Age: 9000},
	}

	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	//for k, v := range m {
	//	fmt.Println(k, "=>", v.Name)
	//}

	for _, v := range stus {
		//jsonStr, _ := json.Marshal(v)
		//fmt.Printf("%s", jsonStr)
		t := reflect.TypeOf(v)
		for i := 0; i < t.NumField(); i++ {
			fmt.Println(t.Field(i).Tag)
		}
	}
}

func TestBuffer(t *testing.T) {
	// 初始零值的缓冲区为空
	//var buf bytes.Buffer
	//println(buffer.ReadByte())

}
