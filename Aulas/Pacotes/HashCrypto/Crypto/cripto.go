package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	c := sha1.New()
	c.Write([]byte("código criptografado com go"))
	bs := c.Sum([]byte{})
	fmt.Printf("%x", bs)

	/* Versão da Doc>
		c := sha1.New()
	f := "código criptografado com go"
	c.Write([]byte(f))
	bs := c.Sum(nil)
	fmt.Println(f)
	fmt.Printf("%x\n", bs) //Mostra como hexadecimal
	*/
}
