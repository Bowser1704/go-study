package main

import (
	"fmt"
	"reflect"
	"github.com/gin-gonic/gin"
)

func main() {
	type values map[string][]string
	var  x values
	test(x)
}

func test(i interface{}) {
	switch x := i.(type) {
	default:
		fmt.Println(x)
	}
	v := reflect.ValueOf(i)
	//v.Elem()
	//reflect.Indirect()
	fmt.Println(v.String())
	fmt.Println('\u263a')

	//fmt.Println(v.Kind(), v.Type())
	//v.NumField()
}