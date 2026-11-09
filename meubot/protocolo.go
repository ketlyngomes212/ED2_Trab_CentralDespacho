package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var (
	leitor      *bufio.Scanner
	escritor    *bufio.Writer
	hospitais   []Hospital
	tab         *tabelaChamados // chamados abertos por id
	filaPrazos  *heapPrazos     // prazos dos chamados abertos (o menor no topo)
	alturaMapa  int             // número de linhas do mapa
	indice      *indiceEspacial // chamados comuns abertos por posição
	buffAtribui []int           // reaproveitado a cada turno
)

// lerCampos lê uma linha e separa em campos; nil quando acaba a entrada.
func lerCampos() []string {
	if !leitor.Scan() {
		return nil
	}
	return strings.Fields(leitor.Text())
}

func inteiro(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic("numero invalido na entrada: " + s)
	}
	return v
}

// lerInicio lê a mensagem inicial: guarda a altura e os hospitais e consome
// ambulâncias e bloqueios (etapa 1 não usa bloqueios).
func lerInicio() {
	campos := lerCampos() // INICIO largura altura turnos limite_turno limite_preparo equipe num_equipes
	alturaMapa = inteiro(campos[2])

	campos = lerCampos() // HOSPITAIS quantidade
	qtdHospitais := inteiro(campos[1])
	for i := 0; i < qtdHospitais; i++ {
		campos = lerCampos()
		hospitais = append(hospitais, Hospital{X: inteiro(campos[0]), Y: inteiro(campos[1])})
	}

	campos = lerCampos() // AMBULANCIAS quantidade
	qtdAmb := inteiro(campos[1])
	for i := 0; i < qtdAmb; i++ {
		lerCampos()
	}

	campos = lerCampos() // BLOQUEIOS quantidade
	qtdBloq := inteiro(campos[1])
	for i := 0; i < qtdBloq; i++ {
		lerCampos()
	}

	lerCampos() // FIM
}

// lerTurno lê a mensagem de um turno e atualiza as estruturas.
// Devolve o turno, as ambulâncias da equipe e se foi possível ler.
func lerTurno() (int, []Ambulancia, bool) {
	campos := lerCampos() // TURNO número
	if len(campos) < 2 {
		return 0, nil, false
	}
	turno := inteiro(campos[1])

	campos = lerCampos() // CHAMADOS quantidade
	qtdChamados := inteiro(campos[1])
	for i := 0; i < qtdChamados; i++ {
		campos = lerCampos()
		c := Chamado{
			ID:    inteiro(campos[0]),
			X:     inteiro(campos[1]),
			Y:     inteiro(campos[2]),
			Valor: inteiro(campos[3]),
			Prazo: inteiro(campos[4]),
		}
		if tab.Inserir(c) {
			inserirNoIndice(c)
			filaPrazos.Inserir(c.Prazo, c.ID)
		}
	}

	campos = lerCampos() // RECOLHIDOS quantidade
	qtdRecolhidos := inteiro(campos[1])
	for i := 0; i < qtdRecolhidos; i++ {
		campos = lerCampos()
		c, ok := tab.Remover(inteiro(campos[0]))
		if ok {
			removerDoIndice(c)
		}
	}

	campos = lerCampos() // AMBULANCIAS quantidade
	qtdAmb := inteiro(campos[1])
	ambulancias := make([]Ambulancia, 0, qtdAmb)
	for i := 0; i < qtdAmb; i++ {
		campos = lerCampos()
		ambulancias = append(ambulancias, Ambulancia{
			ID:    inteiro(campos[0]),
			X:     inteiro(campos[1]),
			Y:     inteiro(campos[2]),
			Livre: campos[3] == "LIVRE",
		})
	}

	campos = lerCampos() // PLACAR quantidade
	qtdEquipes := inteiro(campos[1])
	for i := 0; i < qtdEquipes; i++ {
		lerCampos()
	}

	lerCampos() // FIM

	expirar(turno, filaPrazos, tab)
	return turno, ambulancias, true
}

// jogarTurno envia um movimento por ambulância e termina com FIM.
// atribuidos impede duas ambulâncias da equipe de irem ao mesmo chamado.
func jogarTurno(turno int, ambulancias []Ambulancia) {
	atribuidos := buffAtribui[:0]

	for _, a := range ambulancias {
		mov, id := proximoMovimento(a, hospitais, atribuidos, turno)
		if id != -1 {
			atribuidos = append(atribuidos, id)
		}
		fmt.Fprintln(escritor, a.ID, mov)
	}
	buffAtribui = atribuidos

	fmt.Fprintln(escritor, "FIM")
	escritor.Flush()
}

// iniciar lê a mensagem inicial, cria as estruturas (memória já reservada),
// avisa PRONTO e processa os turnos até o servidor fechar a entrada.
func iniciar() {
	lerInicio()

	tab = novaTabela()
	indice = novoIndice(alturaMapa)
	indiceGrave = novoIndice(alturaMapa)
	filaPrazos = &heapPrazos{itens: make([]item, 0, 1000000)}
	buffAtribui = make([]int, 0, 256)

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