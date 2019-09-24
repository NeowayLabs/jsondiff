package jsondiff_test

import (
	"reflect"
	"testing"

	"github.com/NeowayLabs/jsondiff"
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
	actual := jsondiff.Diff(value, value2)

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
	actual := jsondiff.Diff(value, value)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldReturnEmptyWhenValuesAreEmpty(t *testing.T) {
	value := map[string]interface{}{}
	value2 := map[string]interface{}{}

	expected := []string{}
	actual := jsondiff.Diff(value, value2)

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
	actual := jsondiff.Diff(value, value2)

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
	actual := jsondiff.Diff(value, value2)

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
	actual := jsondiff.Diff(value, value2)

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
	actual := jsondiff.Diff(value, value2)

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
	actual := jsondiff.Diff(value, value2)

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
	actual := jsondiff.Diff(value, value2)

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
	actual := jsondiff.Diff(value, value2)

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
	actual := jsondiff.Diff(value, value)

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

	expected := []string{"json.num"}
	actual := jsondiff.Diff(value, value2)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldReturnFielsdWhenMoreThanOneFieldIsDifferent(t *testing.T) {
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

	expected := []string{"array", "json.num", "num", "str"}
	actual := jsondiff.Diff(value, value2)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
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
	actual := jsondiff.Diff(value, value)

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
	expected := []string{"data[0].socios.documentoSocio", "data[0].socios.nome"}
	actual := jsondiff.Diff(value, value2)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldReturnFieldWhenTheLenghtOfAnJsonArrayDoesNotMatch(t *testing.T) {
	value := map[string]interface{}{
		"data": map[string]interface{}{
			"socios": []interface{}{
				map[string]interface{}{
					"nome":           "Maria",
					"documentoSocio": 1,
				},
			},
		},
	}
	value2 := map[string]interface{}{
		"data": map[string]interface{}{
			"socios": []interface{}{
				map[string]interface{}{
					"nome":           "Maria",
					"documentoSocio": 1,
				},
				map[string]interface{}{
					"nome":           "Joao",
					"documentoSocio": 2,
				},
			},
		},
	}
	expected := []string{"data.socios"}
	actual := jsondiff.Diff(value, value2)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldReturnFieldWhenTheJsonHasThirdLevelAndNotMatch(t *testing.T) {
	value := map[string]interface{}{
		"data": map[string]interface{}{
			"socios": map[string]interface{}{
				"nome":           "Maria",
				"documentoSocio": 1,
			},
		},
		"str": "a",
		"num": 1,
	}
	value2 := map[string]interface{}{
		"data": map[string]interface{}{
			"socios": map[string]interface{}{
				"nome":           "Jose",
				"documentoSocio": 2,
			},
		},
		"str": "a",
		"num": 1,
	}
	expected := []string{"data.socios.documentoSocio", "data.socios.nome"}
	actual := jsondiff.Diff(value, value2)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldNotExplodeWhenAFieldIsNil(t *testing.T) {
	value := map[string]interface{}{
		"socios": nil,
	}

	expected := []string{}
	actual := jsondiff.Diff(value, value)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldNotExplodeWhenFieldsAreFromDifferentType(t *testing.T) {
	value := map[string]interface{}{
		"socios": map[string]interface{}{
			"a": "b",
		},
	}
	value2 := map[string]interface{}{
		"socios": "a",
	}

	expected := []string{"socios"}
	actual := jsondiff.Diff(value, value2)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldReturnChangesWhenAObjectIsMissing(t *testing.T) {
	value := map[string]interface{}{
		"socios": map[string]interface{}{
			"a": "b",
			"c": "d",
		},
		"nome": "maria",
	}
	value2 := map[string]interface{}{
		"nome": "maria",
	}

	expected := []string{"socios"}
	actual := jsondiff.Diff(value, value2)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}
