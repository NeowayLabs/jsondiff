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
		compareMap(v, secondValue, parent, changes)
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
		compareSlice(v, secondValue, parent, changes)
	default:
		if firstValue != secondValue {
			changes[parent] = struct{}{}
		}
	}
}

func compareMap(v map[string]interface{}, secondValue interface{}, parent string, changes map[string]struct{}) {

	v2 := make(map[string]interface{})
	switch v := secondValue.(type) {
	case map[string]interface{}:
		if parent != "" {
			parent = parent + "."
		}
		v2 = v
	default:
		changes[parent] = struct{}{}
		return
	}

	for k, value := range v {
		compare(value, v2[k], parent+k, changes)
	}
}

func compareSlice(v []interface{}, secondValue interface{}, parent string, changes map[string]struct{}) {
	for k, value := range v {
		switch v2 := secondValue.(type) {
		case []interface{}:
			if len(v) != len(v2) {
				changes[parent] = struct{}{}
			} else {
				compare(value, v2[k], parent+"["+strconv.Itoa(k)+"]", changes)
			}
		default:
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
