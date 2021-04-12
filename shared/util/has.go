package util

import (
	"fmt"
	"reflect"
)

func Has(v interface{}, requiredFields ... string) {

}

// Inplace modification
func CopyFieldsByName(s interface{}, res interface{}) {
	va := reflect.ValueOf(s)
	ta := reflect.TypeOf(s)
	tt := reflect.ValueOf(res).Elem().Type()
	tv := reflect.ValueOf(res).Elem()

	fmt.Println(va.NumField())

	for i:=0;i<va.NumField();i++ {
		field := va.Field(i)
		fname := ta.Field(i).Name
		fmt.Println(fname)
		tf,ok  := tt.FieldByName(fname)
		if !ok {
			panic(fmt.Sprintf("No such field: %v", fname))
		}
		if field.Type() == tf.Type {
			tv.FieldByName(fname).Set(field)
		} else {
			panic(fmt.Sprintf("mismatch type for field %v, %v vs %v", fname, field.Type().String(), tf.Type))
		}
	}
}