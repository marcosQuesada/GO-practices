package main

import "fmt"
import "reflect"
import "strconv"

type T struct{}

func (t *T) Foo(c, i int) int {

	res := c + i
	fmt.Printf("Called: ", c, i, res)
	return res
}
func (t *T) Bar() string {
	fmt.Printf("On Bar")
	return "fromBarrrr"
}

func main() {
	var t T
	bar := reflect.ValueOf(&t).MethodByName("Bar")
	result := bar.Call(nil)
	//fmt.Printf("Result is ", result[0].String())

	init := reflect.ValueOf(&t).MethodByName("Foo")
	in := make([]reflect.Value, 2)
	in[0] = reflect.ValueOf(1)
	in[1] = reflect.ValueOf(2)
	result = init.Call(in)
	//fmt.Printf("Result is ", result[0].Int())

	//handle type conversion
	switch result[0].Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Printf("fooo", strconv.FormatInt(result[0].Int(), 10))
	case reflect.String:
		fmt.Printf("fooo", result[0].String())
	}

}
