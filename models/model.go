package models

import (
	"reflect"
)

func UpdateFields(input interface{}, dest interface{}) {
	inputValues := reflect.ValueOf(input)
	destValues := reflect.ValueOf(dest)

	for i := 0; i < inputValues.NumField(); i++ {
		inputVal := inputValues.Field(i).Elem()
		inputTyp := inputValues.Type().Field(i).Name

		if inputVal.IsValid() {
			reflect.Indirect(destValues).FieldByName(inputTyp).Set(inputVal)
		}

	}
}
