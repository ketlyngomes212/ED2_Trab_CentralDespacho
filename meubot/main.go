//O main.go inicializa o programa, avisa o servidor do estado PRONTO,
//lê as informações a cada rodada (chamados, posições e estado das ambulâncias) e
//envia um movimento para cada ambulância.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// encontrarToken retorna a posição de uma palavra na mensagem.
// Retorna -1 caso ela não seja encontrada.
func encontrarToken(campos []string, palavra string) int {
	for i, campo := range campos {
		if campo == palavra {
			return i
		}
	}
	return -1
}

// responderTurno lê os IDs das ambulâncias e envia P (permanecer parado).
func responderTurno(campos []string, saida *bufio.Writer) {
	pos := encontrarToken(campos, "AMBULANCIAS")
	if pos == -1 || pos+1 >= len(campos) {
		return
	}

	quantidade, err := strconv.Atoi(campos[pos+1])
	if err != nil || quantidade < 0 {
		return
	}

	// Cada ambulância possui: ID, X, Y e estado.
	indice := pos + 2

	for i := 0; i < quantidade; i++ {
		if indice+3 >= len(campos) {
			return
		}

		id := campos[indice]

		// P significa permanecer na posição atual.
		fmt.Fprintf(saida, "%s P\n", id)

		indice += 4
	}

	fmt.Fprintln(saida, "FIM")
	saida.Flush()
}

func main() {
	entrada := bufio.NewScanner(os.Stdin)

	// As mensagens podem conter milhares de chamados.
	// Aumentamos o limite padrão de leitura do Scanner.
	entrada.Buffer(make([]byte, 64*1024), 16*1024*1024)

	saida := bufio.NewWriter(os.Stdout)
	defer saida.Flush()

	for entrada.Scan() {
		linha := strings.TrimSpace(entrada.Text())
		if linha == "" {
			continue
		}

		campos := strings.Fields(linha)

		switch campos[0] {
		case "INICIO":
			// A inicialização foi recebida.
			fmt.Fprintln(saida, "PRONTO")
			saida.Flush()

		case "TURNO":
			// Por enquanto, todas as ambulâncias ficam paradas.
			responderTurno(campos, saida)
		}
	}
}
