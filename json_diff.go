package jsondiff

import (
	"reflect"
)

func Diff(firstJson map[string]interface{}, secondJson map[string]interface{}) []string {

	jsonDiff := make([]string, 0)

	for firstKey, firstValue := range firstJson {
		for secondKey, secondValue := range secondJson {
			if firstKey != secondKey {
				continue
			} else if firstValue == nil || secondValue == nil {
				if firstValue == secondValue {
					continue
				} else {
					jsonDiff = append(jsonDiff, firstKey)
				}
			} else if reflect.TypeOf(firstValue).String() == "[]string" {
				if !compareArray(firstValue.([]string), secondValue.([]string)) {
					jsonDiff = append(jsonDiff, firstKey)
				}
			} else if reflect.TypeOf(firstValue).String() == "[]interface {}" {
				if !reflect.DeepEqual(firstValue, secondValue) {
					jsonDiff = append(jsonDiff, firstKey)
				}
			} else if reflect.TypeOf(firstValue).String() == "map[string]interface {}" {
				isDifferent := Diff(firstValue.(map[string]interface{}), secondValue.(map[string]interface{}))
				if len(isDifferent) > 0 {
					jsonDiff = append(jsonDiff, firstKey)
				}
			} else if firstValue != secondValue {
				jsonDiff = append(jsonDiff, firstKey)
			}
		}
	}

	for firstKey := range firstJson {
		if secondJson[firstKey] == nil && firstJson[firstKey] != nil {
			jsonDiff = append(jsonDiff, firstKey)
		}
	}

	for firstKey := range secondJson {
		if firstJson[firstKey] == nil && secondJson[firstKey] != nil {
			jsonDiff = append(jsonDiff, firstKey)
		}
	}

	return jsonDiff
}

func compareArray(firstArray []string, secondArray []string) bool {

	if len(firstArray) != len(secondArray) {
		return false
	}

	for index, _ := range firstArray {
		if firstArray[index] != secondArray[index] {
			return false
		}
	}

	return true
}
