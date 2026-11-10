//main.go é o ponto inicial do bot

package main

import (
	"bufio"
	"os"
)

func main() {
	//configurações do buffer
	leitor = bufio.NewScanner(os.Stdin)
	leitor.Buffer(make([]byte, 64*1024), 16*1024*1024)
	escritor = bufio.NewWriter(os.Stdout)

	iniciar()
}
