package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var (
	leitor    *bufio.Scanner
	escritor  *bufio.Writer
	hospitais []Hospital
	chamados  []Chamado
)

// Lê uma linha do servidor e separa seu conteúdo em campos.
// Retorna nil quando não há mais entrada disponível.
func lerCampos() []string {
	if !leitor.Scan() {
		return nil
	}

	return strings.Fields(leitor.Text())
}

// Converte um campo de texto em inteiro.
// Se a conversão falhar, interrompe a execução com uma mensagem de erro.
func inteiro(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic("numero invalido na entrada: " + s)
	}

	return v
}

// Lê as informações iniciais da partida.
// Guarda os hospitais e consome os dados das ambulâncias e dos bloqueios.
// Na Etapa 1, os bloqueios não são usados na estratégia.
func lerInicio() {
	lerCampos() // INICIO

	campos := lerCampos() // HOSPITAIS quantidade
	quantidadeHospitais := inteiro(campos[1])

	for i := 0; i < quantidadeHospitais; i++ {
		campos = lerCampos()

		hospital := Hospital{
			X: inteiro(campos[0]),
			Y: inteiro(campos[1]),
		}

		hospitais = append(hospitais, hospital)
	}

	campos = lerCampos() // AMBULANCIAS quantidade
	quantidadeAmbulancias := inteiro(campos[1])

	for i := 0; i < quantidadeAmbulancias; i++ {
		lerCampos() // ID e coordenadas iniciais da ambulância
	}

	campos = lerCampos() // BLOQUEIOS quantidade
	quantidadeBloqueios := inteiro(campos[1])

	for i := 0; i < quantidadeBloqueios; i++ {
		lerCampos() // Coordenadas do bloqueio
	}

	lerCampos() // FIM da mensagem inicial
}

// Remove da lista chamados já escolhidos
func removerPorID(id int) {
	for i := range chamados {
		if chamados[i].ID == id {
			ultimo := len(chamados) - 1
			chamados[i] = chamados[ultimo]
			chamados = chamados[:ultimo]
			return
		}
	}
}

// Remove da lista chamados já expirados.
func removerExpirados(turno int) {
	vivos := chamados[:0]

	for _, chamado := range chamados {
		if chamado.Prazo >= turno {
			vivos = append(vivos, chamado)
		}
	}

	chamados = vivos
}

// Lê todas as informações de um turno.
// Retorna o número do turno, a lista atual de ambulâncias e um indicador
// que informa se foi possível ler um novo turno.
func lerTurno() (int, []Ambulancia, bool) {
	campos := lerCampos() // TURNO número
	if len(campos) < 2 {
		return 0, nil, false
	}

	turno := inteiro(campos[1])

	// Lê os novos chamados disponibilizados neste turno.
	campos = lerCampos() // CHAMADOS quantidade
	quantidadeChamados := inteiro(campos[1])

	for i := 0; i < quantidadeChamados; i++ {
		campos = lerCampos()

		chamado := Chamado{
			ID:    inteiro(campos[0]),
			X:     inteiro(campos[1]),
			Y:     inteiro(campos[2]),
			Valor: inteiro(campos[3]),
			Prazo: inteiro(campos[4]),
		}

		chamados = append(chamados, chamado)
	}

	// Remove da lista os chamados que já foram recolhidos.
	campos = lerCampos() // RECOLHIDOS quantidade
	quantidadeRecolhidos := inteiro(campos[1])

	for i := 0; i < quantidadeRecolhidos; i++ {
		campos = lerCampos()
		id := inteiro(campos[0])
		removerPorID(id)
	}

	// Lê a posição e o estado atual de cada ambulância.
	campos = lerCampos() // AMBULANCIAS quantidade
	quantidadeAmbulancias := inteiro(campos[1])

	ambulancias := make([]Ambulancia, 0, quantidadeAmbulancias)

	for i := 0; i < quantidadeAmbulancias; i++ {
		campos = lerCampos()

		ambulancia := Ambulancia{
			ID:    inteiro(campos[0]),
			X:     inteiro(campos[1]),
			Y:     inteiro(campos[2]),
			Livre: campos[3] == "LIVRE",
		}

		ambulancias = append(ambulancias, ambulancia)
	}

	// Consome os dados do placar. Eles não são necessários para
	// escolher os movimentos nesta versão do bot.
	campos = lerCampos() // PLACAR quantidade
	quantidadeEquipes := inteiro(campos[1])

	for i := 0; i < quantidadeEquipes; i++ {
		lerCampos()
	}

	lerCampos() // FIM da mensagem do turno

	// Descarta chamados expirados antes de calcular os movimentos.
	removerExpirados(turno)

	return turno, ambulancias, true
}

// Envia uma decisão para cada ambulância e termina a resposta com FIM.
// A lista atribuidos impede que duas ambulâncias escolham o mesmo
// chamado durante este turno.
func jogarTurno(turno int, ambulancias []Ambulancia) {
	atribuidos := make([]int, 0, len(ambulancias))

	for _, ambulancia := range ambulancias {
		movimento, chamadoID := proximoMovimento(
			ambulancia,
			chamados,
			hospitais,
			atribuidos,
			turno,
		)

		if chamadoID != -1 {
			atribuidos = append(atribuidos, chamadoID)
		}

		fmt.Fprintln(escritor, ambulancia.ID, movimento)
	}

	fmt.Fprintln(escritor, "FIM")
	escritor.Flush()
}

// Inicializa o protocolo, informa que o bot está pronto e processa
// os turnos até que o servidor encerre a comunicação.
func iniciar() {
	lerInicio()

	fmt.Fprintln(escritor, "PRONTO")
	escritor.Flush()

	for {
		turno, ambulancias, ok := lerTurno()
		if !ok {
			return
		}

		jogarTurno(turno, ambulancias)
	}
}
