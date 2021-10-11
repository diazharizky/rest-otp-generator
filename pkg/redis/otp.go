package redis

import (
	"fmt"
	"reflect"
)

type OTPValue struct {
	Passcode string `json:"passcode" redis:"passcode"`
	Attempts int8   `json:"attempts" redis:"attempts"`
}

func (o *OTPValue) GetPropNames() {
	v := reflect.ValueOf(&o)

	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	fmt.Println(values)
}
