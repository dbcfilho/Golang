package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	cenarioDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Avenida X", "Avenida"},
		{"Avenida Y", "Avenida"},
		{"Estrada Raposo Tavares", "Estrada"},
		{"Estrada Imigrantes", "Estrada"},
		{"Praça X", "Tipo Inválido!"},
		{"Casa Y", "Tipo Inválido!"},
		{"", "Tipo Inválido!"},
	}

	for _, cenario := range cenarioDeTeste {
		tipoEndrecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoEndrecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo %s é diferente do esperado %s", tipoEndrecoRecebido, cenario.retornoEsperado)
		}
	}
}
