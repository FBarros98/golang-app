// Go é uma linguagem compilável. Então, é necessário transformar o código em um executável. Informa que esse é o principal pacote da app.
// "Go build main.go" para compilar e criar um exe
// "Go run main.go" Já compila e executa
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Função principal
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	test := os.Getenv("Test") // := utilizado para declarar que se trata de uma variável
	fmt.Println(test)         // A função deve iniciar com letra maiúscula para indicar que veio de um pacote externo.

}
