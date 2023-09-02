package v3utils

import (
	"fmt"
	"reflect"
)

// var v3Convert = new(V3Convert)

type V3Convert struct{}

func (uv V3Convert) StructToSliceOfSlice(data interface{}) [][]string {
	value := reflect.ValueOf(data)

	if value.Kind() != reflect.Slice {
		return nil
	}

	numRows := value.Len()
	if numRows == 0 {
		return nil
	}

	structType := value.Index(0).Type()
	numFields := structType.NumField()

	result := make([][]string, numRows)

	for i := 0; i < numRows; i++ {
		result[i] = make([]string, numFields)
		structValue := value.Index(i)
		for j := 0; j < numFields; j++ {
			field := structValue.Field(j)
			result[i][j] = fmt.Sprintf("%v", field.Interface())
		}
	}

	return result
}

func (uv V3Convert) StructToSliceOfSliceWithHeader(data interface{}) [][]string {
	value := reflect.ValueOf(data)

	if value.Kind() != reflect.Slice {
		return nil
	}

	numRows := value.Len()
	if numRows == 0 {
		return nil
	}

	structType := value.Index(0).Type()
	numFields := structType.NumField()

	header := make([]string, numFields)
	result := make([][]string, numRows)

	for i := 0; i < numFields; i++ {
		header[i] = structType.Field(i).Name
	}

	for i := 0; i < numRows; i++ {
		result[i] = make([]string, numFields)
		structValue := value.Index(i)
		for j := 0; j < numFields; j++ {
			field := structValue.Field(j)
			result[i][j] = fmt.Sprintf("%v", field.Interface())
		}
	}
	result = append([][]string{header}, result...)

	return result
}
