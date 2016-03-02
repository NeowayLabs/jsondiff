package jsondiff

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestShouldReturnFieldWhenValuesDiffer(t *testing.T) {

	value := map[string]interface{}{
		"num": 6.13,
		"str": "a",
	}

	value2 := map[string]interface{}{
		"num": 6.13,
		"str": "b",
	}

	expected := []string{"str"}
	actual := Diff(value, value2)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldReturnEmptyWhenValuesMatch(t *testing.T) {
	value := map[string]interface{}{
		"num": 6.13,
		"str": "a",
	}

	expected := []string{}
	actual := Diff(value, value)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldReturnEmptyWhenValuesAreEmpty(t *testing.T) {
	value := map[string]interface{}{}
	value2 := map[string]interface{}{}

	expected := []string{}
	actual := Diff(value, value2)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldReturnFieldWhenTheSecondJsonHasMoreItens(t *testing.T) {
	value := map[string]interface{}{
		"num": 6.13,
		"str": "a",
	}

	value2 := map[string]interface{}{
		"num":  6.13,
		"str":  "a",
		"bool": true,
	}

	expected := []string{"bool"}
	actual := Diff(value, value2)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldReturnFieldWhenTheFirstJsonHasMoreItens(t *testing.T) {
	value := map[string]interface{}{
		"num":  6.13,
		"str":  "a",
		"bool": true,
	}

	value2 := map[string]interface{}{
		"num": 6.13,
		"str": "a",
	}

	expected := []string{"bool"}
	actual := Diff(value, value2)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShoudReturnEmptyWhenValuesMatchAndAreInDifferentOrder(t *testing.T) {
	value := map[string]interface{}{
		"num": 6.13,
		"str": "a",
	}

	value2 := map[string]interface{}{
		"str": "a",
		"num": 6.13,
	}

	expected := []string{}
	actual := Diff(value, value2)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldReturnFieldWhenThereIsAnArrayThatIsDifferent(t *testing.T) {
	value := map[string]interface{}{
		"str": []string{"a", "b"},
		"num": 6.13,
	}

	value2 := map[string]interface{}{
		"str": []string{"a", "b", "c"},
		"num": 6.13,
	}

	expected := []string{"str"}
	actual := Diff(value, value2)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldReturnFieldWhenTheSecondArrayHasDifferentElement(t *testing.T) {
	value := map[string]interface{}{
		"str": []string{"a", "a"},
		"num": 6.13,
	}

	value2 := map[string]interface{}{
		"str": []string{"a", "b"},
		"num": 6.13,
	}

	expected := []string{"str"}
	actual := Diff(value, value2)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldReturnFieldWhenTheElementsOfAnArrayDoesNotMatch(t *testing.T) {
	value := map[string]interface{}{
		"str": []string{"a", "b"},
		"num": 6.13,
	}

	value2 := map[string]interface{}{
		"str": []string{"a", "c"},
		"num": 6.13,
	}

	expected := []string{"str"}
	actual := Diff(value, value2)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldReturnEmptyWhenTheElementsOfAnArrayMatch(t *testing.T) {
	value := map[string]interface{}{
		"str": []string{"a", "b"},
		"num": 6.13,
	}

	value2 := map[string]interface{}{
		"str": []string{"a", "b"},
		"num": 6.13,
	}

	expected := []string{}
	actual := Diff(value, value2)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldReturnEmptyWhenTheFieldIsAJsonAndMatch(t *testing.T) {
	value := map[string]interface{}{
		"json": map[string]interface{}{
			"str": "a",
			"num": 1,
		},
	}

	expected := []string{}
	actual := Diff(value, value)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldReturnFieldWhenTheFieldIsAJsonAndDoesNotMatch(t *testing.T) {
	value := map[string]interface{}{
		"json": map[string]interface{}{
			"str": "a",
			"num": 1,
		},
	}
	value2 := map[string]interface{}{
		"json": map[string]interface{}{
			"str": "a",
			"num": 2,
		},
	}

	expected := []string{"json"}
	actual := Diff(value, value2)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldReturnFielsdWhenMoreThanOneFieldAreDifferent(t *testing.T) {
	value := map[string]interface{}{
		"json": map[string]interface{}{
			"str": "a",
			"num": 1,
		},
		"array": []string{"a", "b"},
		"str":   "a",
		"num":   1,
	}
	value2 := map[string]interface{}{
		"json": map[string]interface{}{
			"str": "a",
			"num": 2,
		},
		"array": []string{"a", "c"},
		"str":   "b",
		"num":   2,
	}

	actual := Diff(value, value2)

	assert.Contains(t, actual, "json", "array", "str", "num")
}

func TestShouldReturnEmptyWhenTheElementsOfAnArrayAreJsonAndMatch(t *testing.T) {

	value := map[string]interface{}{
		"array": []interface{}{
			map[string]interface{}{
				"json": map[string]interface{}{
					"str": "a",
					"num": 1,
				},
			},
			map[string]interface{}{
				"json": map[string]interface{}{
					"str": "b",
					"num": 2,
				},
			},
		},
		"str": "a",
		"num": 1,
	}
	expected := []string{}
	actual := Diff(value, value)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldReturnFieldWhenTheElementsOfAnArrayAreJsonAndDoesNotMatch(t *testing.T) {
	value := map[string]interface{}{
		"data": []interface{}{
			map[string]interface{}{
				"socios": map[string]interface{}{
					"nome":           "Maria",
					"documentoSocio": 1,
				},
			},
		},
		"str": "a",
		"num": 1,
	}
	value2 := map[string]interface{}{
		"data": []interface{}{
			map[string]interface{}{
				"socios": map[string]interface{}{
					"nome":           "Jose",
					"documentoSocio": 2,
				},
			},
		},
		"str": "a",
		"num": 1,
	}
	expected := []string{"data"}
	actual := Diff(value, value2)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldNotExplodeWhenAFieldIsNil(t *testing.T) {
	value := map[string]interface{}{
		"socios": nil,
	}

	expected := []string{}
	actual := Diff(value, value)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}
