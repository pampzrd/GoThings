/* Você e seus colegas querem criar em formato de código uma antiga brincadeira: Ao falar os números sempre que aparecer um múltiplo de 3, o participante deve falar "Pin" e nos múltiplos de 5 o participante deve falar "Pan". Então, seu programa deve imprimir números de 1 a 100 e nos casos citados, não devem aparecer os números, mas sim o que o participante deve falar. Caso seja múltiplo de 3 e 5 ao mesmo tempo, deve mostrar "Pin""Pan".
 */

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Pin-Pan!")
			continue
		} else if i%5 == 0 {
			fmt.Println("Pan!")
			continue
		} else if i%3 == 0 {
			fmt.Println("Pin!")
			continue
		} else {
			fmt.Println(i)
		}

	}
}
