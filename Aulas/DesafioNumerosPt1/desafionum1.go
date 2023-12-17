/* Você e seus colegas querem criar um código que exiba todos os números de 1 até 100 que sejam divisíveis por 3. % e for.
 */

package main

import "fmt"

func main() {
	fmt.Println("Divisíveis por 3 até 100")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

}
