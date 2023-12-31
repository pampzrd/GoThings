package main

import "testing"

func TestSoma_ShouldSumCorrect(t *testing.T) {
	teste := soma(1, 2, 3)
	resultado := 6
	if teste != resultado {
		t.Error("Valor esperado ", resultado, " Valor obtido: ", teste)
	}
}

func TestSoma_ShouldSumIncorrect(t *testing.T) {
	teste := soma(1, 2, 3)
	resultado := 7
	if teste != resultado {
		t.Error("Valor esperado ", resultado, " Valor obtido: ", teste)
	}
}

func TestSubtracao_ShouldSubCorrect(t *testing.T) {
	teste := subtracao(13, 1, 2)
	resultado := 10
	if teste != resultado {
		t.Error("Valor esperado ", resultado, " Valor obtido: ", teste)
	}
}

func TestSubtracao_ShouldSubIncorrect(t *testing.T) {
	teste := subtracao(13, 1, 2)
	resultado := -16
	if teste != resultado {
		t.Error("Valor esperado ", resultado, " Valor obtido: ", teste)
	}
}
func TestMultiplicacao_ShouldMultCorrect(t *testing.T) {
	teste := multiplicacao(10, 10, 2, 3)
	resultado := 600
	if teste != resultado {
		t.Error("Valor esperado: ", resultado, " Valor obtido: ", teste)
	}
}

func TestMultiplicacao_ShouldMultIncorrect(t *testing.T) {
	teste := multiplicacao(10, 10, 2, 3)
	resultado := 200
	if teste != resultado {
		t.Error("Valor esperado: ", resultado, " Valor obtido: ", teste)
	}
}

func TestDivisao_ShouldDivCorrect(t *testing.T) {
	teste := divisao(20, 2, 5)

	resultado := 2.00
	if teste != resultado {
		t.Error("Valor esperado: ", resultado, " Valor obtido: ", teste)
	}
}

func TestDivisao_ShouldDivIncorrect(t *testing.T) {
	teste := divisao(20, 2, 5)

	resultado := 0.00
	if teste != resultado {
		t.Error("Valor esperado: ", resultado, " Valor obtido: ", teste)
	}
}
