// Métodos são funções que executam em um objeto (struct).
/*https://www.youtube.com/watch?v=5b8MMXgBnp0   4:12
 */

package main

import "fmt"

// Exemplo de Objeto a ser referenciado pelo método: Uma struct de alunos.
type Student struct {
	name   string
	grades []int
	age    int
}

// Método FUNC (APELIDO OBJETO) NOME_FUNCAO() TIPO_RETORNO {}
func (s Student) getAge() int {
	return s.age
}

//receber valor e colocá-lo em alguma propriedade do objeto: //Método com argumentos
func (s *Student) setAge(age int) { //usa ponteiro para mudar um valor definitivamente. Se for apenas acessar o valor da propriedade, não precisa usar o ponteiro.
	s.age = age
}

//Método para calcular a média dos alunos.
func (s Student) getAverageGrades() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

//Para achar a nota máxima
func (s *Student) getMaxGrade() int {
	curMax := 0
	for _, v := range s.grades {
		if curMax < v {
			curMax = v
		}
	}
	return curMax
}

func main() {
	s1 := Student{"Tim", []int{70, 90, 80, 85}, 19} // Atribui valores para o objeto
	/*fmt.Println(s1.getAge())
	s1.setAge(32)
	fmt.Println(s1.getAge())*/
	average := s1.getAverageGrades()
	fmt.Println(average)
	maximo := s1.getMaxGrade()
	fmt.Println(maximo)

}
