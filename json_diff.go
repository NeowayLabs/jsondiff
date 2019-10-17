// Package jsondiff provides comparing for json maps.
package jsondiff

import (
	"sort"
	"strconv"
	"strings"
)

// Diff compare two jsons map and returns the fields names that has been changed.
// If input maps are equal, all return values will be empty.
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

// DiffWithValues compare two jsons map and returns the fields names and values that has been changed.
// If input maps are equal, all return values will be empty.
// Otherwise diff slice will contain the field names list,
// those field names will also be present in FirstValues
// and SecondValues combined with their respective
// values from firstJson and secondJson.
func DiffWithValues(firstJson map[string]interface{}, secondJson map[string]interface{}) (diff []string, firstValues map[string]interface{}, secondValues map[string]interface{}) {
	diff = Diff(firstJson, secondJson)
	firstValues = keepDiffFields(firstJson, diff)
	secondValues = keepDiffFields(secondJson, diff)
	return diff, firstValues, secondValues
}

func keepDiffFields(jsonMap map[string]interface{}, diffFields []string) map[string]interface{} {
	const rootPath string = ""
	allowedFieldsSet := make(map[string]struct{})
	for _, f := range diffFields {
		allowedFieldsSet[f] = struct{}{}
	}
	diff := copyMap(jsonMap)
	removeNotAllowedFieldsFromPath(diff, rootPath, allowedFieldsSet)
	return diff
}

func copyMap(m map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range m {
		if vm, ok := v.(map[string]interface{}); ok {
			cp[k] = copyMap(vm)
			continue
		}
		cp[k] = v
	}

	return cp
}

func removeNotAllowedFieldsFromPath(data map[string]interface{}, path string, allowedFields map[string]struct{}) {
	for k, v := range data {
		field := path + k
		if _, exists := allowedFields[field]; exists {
			continue
		}
		if !hasSubFieldAllowed(allowedFields, field) {
			delete(data, k)
			continue
		}
		switch value := v.(type) {
		case map[string]interface{}:
			currentPath := field + "."
			removeNotAllowedFieldsFromPath(value, currentPath, allowedFields)
			if len(value) == 0 {
				delete(data, k)
			}
		}
	}
}

func hasSubFieldAllowed(allowedField map[string]struct{}, parentField string) bool {
	for field := range allowedField {
		if strings.HasPrefix(field, parentField+".") {
			return true
		}
	}
	return false
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
