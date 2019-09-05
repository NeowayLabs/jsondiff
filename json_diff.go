package jsondiff

import (
	"sort"
	"strconv"
)

func Diff(firstJson map[string]interface{}, secondJson map[string]interface{}) []string {

	changesSet := make(map[string]struct{})
	compare(firstJson, secondJson, "", changesSet)
	compare(secondJson, firstJson, "", changesSet)

	changes := make([]string, 0)
	for key := range changesSet {
		changes = append(changes, key)
	}
	sort.Strings(changes)

	return changes
}

func compare(firstValue, secondValue interface{}, parent string, changes map[string]struct{}) {

	switch v := firstValue.(type) {
	case map[string]interface{}:
		if parent != "" {
			parent = parent + "."
		}
		for k, value := range v {
			switch v2 := secondValue.(type) {
			case map[string]interface{}:
				compare(value, v2[k], parent+k, changes)
			default:
				changes[parent+k] = struct{}{}
			}
		}
	case []string:
		switch v2 := secondValue.(type) {
		case []string:
			equal := compareArray(v, v2)
			if !equal {
				changes[parent] = struct{}{}
			}
		default:
			changes[parent] = struct{}{}
		}
	case []interface{}:
		for k, value := range v {
			switch v2 := secondValue.(type) {
			case []interface{}:
				compare(value, v2[k], parent+"["+strconv.Itoa(k)+"]", changes)
			default:
				changes[parent] = struct{}{}
			}
		}
	default:
		if firstValue != secondValue {
			changes[parent] = struct{}{}
		}
	}
}

func compareArray(firstArray []string, secondArray []string) bool {
	if len(firstArray) != len(secondArray) {
		return false
	}
	for index := range firstArray {
		if firstArray[index] != secondArray[index] {
			return false
		}
	}
	return true
}
