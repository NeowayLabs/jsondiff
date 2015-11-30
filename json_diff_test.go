package jsondiff

import (
	"encoding/json"
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

func TestShouldReturnEmptyWhenTheElementsOfAnArrayAreJsonAndMatch(t *testing.T) {
	byt := []byte(`{
		"protocolo": "149123123",
		"cnpj": "110020312312313",
		"nomeEmpresarial": "INCORPORADORA E EMPREENDIMENTOS LTDA",
		"naturezaJuridica": "SOCIEDADE ANONIMA FECHADA ",
		"nire": "2923123123",
		"arquivamento": "25/07/2014",
		"inicioAtividade": "1988-02-18",
		"endereco": "RUA JOSE",
		"numero": "",
		"bairro": "ALPHAVILLE",
		"complemento": "",
		"municipio": "FAKE",
		"cep": "123123123",
		"uf": "SC",
		"objetoSocial": "ABCD iABCDABCDABCDABCDABCDABCDABCDABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD.",
		"capital": "10.000.000,00",
		"capitalIntegralizado": "10.000.000,00",
		"microEmpresa": "Nao",
		"prazoDuracao": "Indeterminado",
		"situacao": "REGISTRO ATIVO",
		"status": "",
		"dia": 4,
		"mes": "Setembro",
		"ano": 2015,
		"moedaCapital": "R$",
		"moedaCapitalIntegralizado": "R$",
		"socios": [
		{
				"nome": "MARIA EDUARDA",
				"documentoSocio":		"123123123",
				"valorParticipacao": "10.000.000,00",
				"tipo": "Sócio Administrador",
				"descricaoAdministrador": ""
		},
		{
				"nome": "TATIANA",
				"documentoSocio":		"123123123",
				"valorParticipacao": "10.000.000,00",
				"tipo": "Sócio Administrador",
				"descricaoAdministrador": ""
		}
		]
    }`)

	var f interface{}
	json.Unmarshal(byt, &f)

	expected := []string{}
	actual := Diff(f.(map[string]interface{}), f.(map[string]interface{}))

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldReturnFieldWhenTheElementsOfAnArrayAreJsonAndDoesNotMatch(t *testing.T) {
	byt := []byte(`{
		"protocolo": "149123123",
		"cnpj": "110020312312313",
		"nomeEmpresarial": "INCORPORADORA E EMPREENDIMENTOS LTDA",
		"naturezaJuridica": "SOCIEDADE ANONIMA FECHADA ",
		"nire": "2923123123",
		"arquivamento": "25/07/2014",
		"inicioAtividade": "1988-02-18",
		"endereco": "RUA JOSE",
		"numero": "",
		"bairro": "ALPHAVILLE",
		"complemento": "",
		"municipio": "FAKE",
		"cep": "123123123",
		"uf": "SC",
		"objetoSocial": "ABCD iABCDABCDABCDABCDABCDABCDABCDABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD.",
		"capital": "10.000.000,00",
		"capitalIntegralizado": "10.000.000,00",
		"microEmpresa": "Nao",
		"prazoDuracao": "Indeterminado",
		"situacao": "REGISTRO ATIVO",
		"status": "",
		"dia": 4,
		"mes": "Setembro",
		"ano": 2015,
		"moedaCapital": "R$",
		"moedaCapitalIntegralizado": "R$",
		"socios": [
		{
				"nome": "MARIA EDUARDA",
				"documentoSocio":		"123123123",
				"valorParticipacao": "10.000.000,00",
				"tipo": "Sócio Administrador",
				"descricaoAdministrador": ""
		},
		{
				"nome": "TATIANA",
				"documentoSocio":		"123123123",
				"valorParticipacao": "10.000.000,00",
				"tipo": "Sócio Administrador",
				"descricaoAdministrador": ""
		}
		]
    }`)

	byt2 := []byte(`{
		"protocolo": "149123123",
		"cnpj": "110020312312313",
		"nomeEmpresarial": "INCORPORADORA E EMPREENDIMENTOS LTDA",
		"naturezaJuridica": "SOCIEDADE ANONIMA FECHADA ",
		"nire": "2923123123",
		"arquivamento": "25/07/2014",
		"inicioAtividade": "1988-02-18",
		"endereco": "RUA JOSE",
		"numero": "",
		"bairro": "ALPHAVILLE",
		"complemento": "",
		"municipio": "FAKE",
		"cep": "123123123",
		"uf": "SC",
		"objetoSocial": "ABCD iABCDABCDABCDABCDABCDABCDABCDABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD ABCD.",
		"capital": "10.000.000,00",
		"capitalIntegralizado": "10.000.000,00",
		"microEmpresa": "Nao",
		"prazoDuracao": "Indeterminado",
		"situacao": "REGISTRO ATIVO",
		"status": "",
		"dia": 4,
		"mes": "Setembro",
		"ano": 2015,
		"moedaCapital": "R$",
		"moedaCapitalIntegralizado": "R$",
		"socios": [
		{
				"nome": "MARIA",
				"documentoSocio":		"123123123",
				"valorParticipacao": "10.000.000,00",
				"tipo": "Sócio Administrador",
				"descricaoAdministrador": ""
		},
		{
				"nome": "TATIANA",
				"documentoSocio":		"123123123",
				"valorParticipacao": "10.000.000,00",
				"tipo": "Sócio Administrador",
				"descricaoAdministrador": ""
		}
		]
    }`)

	var f interface{}
	var g interface{}
	json.Unmarshal(byt, &f)
	json.Unmarshal(byt2, &g)

	expected := []string{"socios"}
	actual := Diff(f.(map[string]interface{}), g.(map[string]interface{}))

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}
