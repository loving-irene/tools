package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"
)

var langs = []string{
	"Golang",
	"PHP",
	"JavaScript",
	"Python",
	"Java",
}

type langPrinter struct {
	lang string
}

func (m *langPrinter) Task() {
	log.Println(m.lang)
	time.Sleep(time.Second)
}

type User struct {
	Name    string
	Website string
	Age     int
	Male    bool
	Skills  []string
}

func main() {
	user := User{
		"lovan",
		"lovan.com",
		29,
		true,
		[]string{"python", "php"},
	}

	e, err := json.Marshal(user)
	if err != nil {
		log.Println("msg:", err)
		return
	}
	fmt.Println("解码数据:", e)
	fmt.Printf("解码数据:%s\n", e)

	var a User
	err = json.Unmarshal(e, &a)
	if err != nil {
		log.Println("msg:", err)
		return
	}
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))

}
