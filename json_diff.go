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
			} else if reflect.TypeOf(firstValue).String() == "[]string" {
				if !compareArray(firstValue.([]string), secondValue.([]string)) {
					jsonDiff = append(jsonDiff, firstKey)
				}
			} else if reflect.TypeOf(firstValue).String() == "[]interface {}" {
				if !reflect.DeepEqual(firstValue, secondValue) {
					jsonDiff = append(jsonDiff, firstKey)
				}
			} else if firstValue != secondValue {
				jsonDiff = append(jsonDiff, firstKey)
			}
		}
	}

	for firstKey := range firstJson {
		if secondJson[firstKey] == nil {
			jsonDiff = append(jsonDiff, firstKey)
		}
	}

	for firstKey := range secondJson {
		if firstJson[firstKey] == nil {
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

	for index, _ := range secondArray {
		if firstArray[index] != secondArray[index] {
			return false
		}
	}

	return true
}
