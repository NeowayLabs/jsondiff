package jsonDiff

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
	actual := getJsonDiff(value, value2)

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
	actual := getJsonDiff(value, value)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldReturnEmptyWhenValuesAreEmpty(t *testing.T) {
	value := map[string]interface{}{}
	value2 := map[string]interface{}{}

	expected := []string{}
	actual := getJsonDiff(value, value2)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldReturnFieldWhenTheNumberOfItensDiffer(t *testing.T) {
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
	actual := getJsonDiff(value, value2)

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
	actual := getJsonDiff(value, value2)

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
	actual := getJsonDiff(value, value2)

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
	actual := getJsonDiff(value, value2)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldReturnEmptyWhenTheElementsOfAnArrayAreJsonAndMatch(t *testing.T) {
	byt := []byte(`{
		"protocolo": "14921002",
		"cnpj": "11002088000105",
		"nomeEmpresarial": "BELLA INCORPORADORA E EMPREENDIMENTOS EPP LTDA",
		"naturezaJuridica": "SOCIEDADE ANONIMA FECHADA ",
		"nire": "29200820065",
		"arquivamento": "25/07/2014",
		"inicioAtividade": "1988-02-18",
		"endereco": "R PATRICIO FARIAS",
		"numero": "",
		"bairro": "ALPHAVILLE I",
		"complemento": "",
		"municipio": "SALVADOR",
		"cep": "41701015",
		"uf": "SC",
		"objetoSocial": "ALUGUEL DE MAQUINAS E EQUIPAMENTOS PARA ESCRITORIO, AGENCIAMENTO DE ESPACOS PARA PUBLICIDADE, EXCETO EM VEICULOS DE COMUNICACAO, ALUGUEL DE EQUIPAMENTOS CIENTIFICOS, MEDICOS E HOSPITALARES , SEM OPERADOR, COMERCIO VAREJISTA ESPECIALIZADO EM EQUIPAMENTOS E SUPRIMENTOS DE INFORMATICA, EDICAO DE CADASTROS, LISTAS E OUTROS PRODUTOS GRAFICOS, FOTOCOPIAS, IMPRESSAO DE MATERIAL PARA USO PUBLICITARIO, LOCACAO DE MAO-DE-OBRA TEMPORARIA, REPARACAO E MANUTENCAO DE COMPUTADORES E DE EQUIPAMENTOS PERIFERICOS, SERVICOS DE ACABAMENTOS GRAFICOS, EXCETO ENCADERNACAO E PLASTIFICACAO E SERVICOS DE MICROFILMAGEM.",
		"capital": "20.000.000,00",
		"capitalIntegralizado": "20.000.000,00",
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
				"nome": "MARIA ELKE GROSSENBACHER",
				"documentoSocio":		"61382388934",
				"valorParticipacao": "10.000.000,00",
				"tipo": "Sócio Administrador",
				"descricaoAdministrador": ""
		},
		{
				"nome": "TATJANA GROSSENBACHER",
				"documentoSocio":		"89118677991",
				"valorParticipacao": "10.000.000,00",
				"tipo": "Sócio Administrador",
				"descricaoAdministrador": ""
		}
		]
}`)

	var f interface{}
	json.Unmarshal(byt, &f)

	expected := []string{}
	actual := getJsonDiff(f.(map[string]interface{}), f.(map[string]interface{}))

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestShouldReturnFieldWhenTheElementsOfAnArrayAreJsonAndDoesNotMatch(t *testing.T) {
	byt := []byte(`{
		"protocolo": "14921002",
		"cnpj": "11002088000105",
		"nomeEmpresarial": "BELLA INCORPORADORA E EMPREENDIMENTOS EPP LTDA",
		"naturezaJuridica": "SOCIEDADE ANONIMA FECHADA ",
		"nire": "29200820065",
		"arquivamento": "25/07/2014",
		"inicioAtividade": "1988-02-18",
		"endereco": "R PATRICIO FARIAS",
		"numero": "",
		"bairro": "ALPHAVILLE I",
		"complemento": "",
		"municipio": "SALVADOR",
		"cep": "41701015",
		"uf": "SC",
		"objetoSocial": "ALUGUEL DE MAQUINAS E EQUIPAMENTOS PARA ESCRITORIO, AGENCIAMENTO DE ESPACOS PARA PUBLICIDADE, EXCETO EM VEICULOS DE COMUNICACAO, ALUGUEL DE EQUIPAMENTOS CIENTIFICOS, MEDICOS E HOSPITALARES , SEM OPERADOR, COMERCIO VAREJISTA ESPECIALIZADO EM EQUIPAMENTOS E SUPRIMENTOS DE INFORMATICA, EDICAO DE CADASTROS, LISTAS E OUTROS PRODUTOS GRAFICOS, FOTOCOPIAS, IMPRESSAO DE MATERIAL PARA USO PUBLICITARIO, LOCACAO DE MAO-DE-OBRA TEMPORARIA, REPARACAO E MANUTENCAO DE COMPUTADORES E DE EQUIPAMENTOS PERIFERICOS, SERVICOS DE ACABAMENTOS GRAFICOS, EXCETO ENCADERNACAO E PLASTIFICACAO E SERVICOS DE MICROFILMAGEM.",
		"capital": "20.000.000,00",
		"capitalIntegralizado": "20.000.000,00",
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
				"nome": "MARIA ANTONIETA",   
				"documentoSocio":		"61382388934",
				"valorParticipacao": "10.000.000,00",
				"tipo": "Sócio Administrador",
				"descricaoAdministrador": ""
		},
		{
				"nome": "TATJANA GROSSENBACHER",
				"documentoSocio":		"89118677991",
				"valorParticipacao": "10.000.000,00",
				"tipo": "Sócio Administrador",
				"descricaoAdministrador": ""
		}
		]
}`)

	byt2 := []byte(`{
		"protocolo": "14921002",
		"cnpj": "11002088000105",
		"nomeEmpresarial": "BELLA INCORPORADORA E EMPREENDIMENTOS EPP LTDA",
		"naturezaJuridica": "SOCIEDADE ANONIMA FECHADA ",
		"nire": "29200820065",
		"arquivamento": "25/07/2014",
		"inicioAtividade": "1988-02-18",
		"endereco": "R PATRICIO FARIAS",
		"numero": "",
		"bairro": "ALPHAVILLE I",
		"complemento": "",
		"municipio": "SALVADOR",
		"cep": "41701015",
		"uf": "SC",
		"objetoSocial": "ALUGUEL DE MAQUINAS E EQUIPAMENTOS PARA ESCRITORIO, AGENCIAMENTO DE ESPACOS PARA PUBLICIDADE, EXCETO EM VEICULOS DE COMUNICACAO, ALUGUEL DE EQUIPAMENTOS CIENTIFICOS, MEDICOS E HOSPITALARES , SEM OPERADOR, COMERCIO VAREJISTA ESPECIALIZADO EM EQUIPAMENTOS E SUPRIMENTOS DE INFORMATICA, EDICAO DE CADASTROS, LISTAS E OUTROS PRODUTOS GRAFICOS, FOTOCOPIAS, IMPRESSAO DE MATERIAL PARA USO PUBLICITARIO, LOCACAO DE MAO-DE-OBRA TEMPORARIA, REPARACAO E MANUTENCAO DE COMPUTADORES E DE EQUIPAMENTOS PERIFERICOS, SERVICOS DE ACABAMENTOS GRAFICOS, EXCETO ENCADERNACAO E PLASTIFICACAO E SERVICOS DE MICROFILMAGEM.",
		"capital": "20.000.000,00",
		"capitalIntegralizado": "20.000.000,00",
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
				"nome": "MARIA ELKE GROSSENBACHER",
				"documentoSocio":		"61382388934",
				"valorParticipacao": "10.000.000,00",
				"tipo": "Sócio Administrador",
				"descricaoAdministrador": ""
		},
		{
				"nome": "TATJANA GROSSENBACHER",
				"documentoSocio":		"89118677991",
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
	actual := getJsonDiff(f.(map[string]interface{}), g.(map[string]interface{}))

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}
