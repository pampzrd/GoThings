package main

import "testing"

//A convenção é Should Sum Correct or Incorrect mas ao alterar os nomes, ele não identifica mais como testes. Precisa de "TEST" no nome?
//Precisa ter o Test no nome.
//Precisa ter o nome da função
func TestSoma_ShouldSumCorrect(t *testing.T) {
	teste := soma(1, 2, 3)
	resultado := 6

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, " Valor retornado: ", teste)
	}
}

func TestSoma_ShouldSumIncorrect(t *testing.T) {
	teste := soma(1, 2, 3)
	resultado := 7

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, " Valor retornado: ", teste)
	}
}

func TestMultiplica_ShouldMultCorrect(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 100
	if teste != resultado {
		t.Error("Valor esperado: ", resultado, " Valor retornado: ", teste)
	}
}

func TestMultiplica_ShouldMultIncorrect(t *testing.T) {
	teste := multiplica(5, 5)
	resultado := 30
	if teste != resultado {
		t.Error("Valor esperado: ", resultado, " Valor retornado: ", teste)
	}
}
