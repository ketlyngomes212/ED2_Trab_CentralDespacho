package main

// Tabela hash (encadeamento com pool de nós) dos chamados abertos, indexada
// pelo id. Sem map e sem alocação durante os turnos.

// no guarda um chamado e o índice do próximo nó da mesma caixa
type no struct {
	prox    int
	chamado Chamado
}

type tabelaChamados struct {
	chaves  []int // chaves[caixa] = índice do primeiro nó da caixa (-1 se vazia)
	nos     []no
	livre   int // primeiro nó livre do pool
	tamanho int
}

const tamanhoTabela = 1000003 // número de caixas (primo)
const capacidadeNos = 1000000 // tamanho do pool de nós

func posicao(id int) int {
	return id % tamanhoTabela
}

func novaTabela() *tabelaChamados {
	t := &tabelaChamados{}
	t.chaves = make([]int, tamanhoTabela)
	t.nos = make([]no, capacidadeNos)
	for i := range t.chaves {
		t.chaves[i] = -1
	}
	for i := range t.nos {
		t.nos[i].prox = i + 1
	}
	t.nos[capacidadeNos-1].prox = -1
	t.livre = 0
	t.tamanho = 0
	return t
}

// Inserir guarda o chamado. Devolve false se o pool acabou.
func (t *tabelaChamados) Inserir(c Chamado) bool {
	if t.livre == -1 {
		return false
	}
	n := t.livre
	t.livre = t.nos[n].prox
	caixa := posicao(c.ID)
	t.nos[n].chamado = c
	t.nos[n].prox = t.chaves[caixa]
	t.chaves[caixa] = n
	t.tamanho++
	return true
}

// Buscar procura o chamado de identificador id.
func (t *tabelaChamados) Buscar(id int) (Chamado, bool) {
	n := t.chaves[posicao(id)]
	for n != -1 {
		if t.nos[n].chamado.ID == id {
			return t.nos[n].chamado, true
		}
		n = t.nos[n].prox
	}
	return Chamado{}, false
}

// Remover tira o chamado da tabela e o devolve, para que quem chamou possa
// retirá-lo também dos outros índices.
func (t *tabelaChamados) Remover(id int) (Chamado, bool) {
	caixa := posicao(id)
	anterior := -1
	n := t.chaves[caixa]
	for n != -1 {
		if t.nos[n].chamado.ID == id {
			c := t.nos[n].chamado
			if anterior == -1 {
				t.chaves[caixa] = t.nos[n].prox
			} else {
				t.nos[anterior].prox = t.nos[n].prox
			}
			t.nos[n].prox = t.livre
			t.livre = n
			t.tamanho--
			return c, true
		}
		anterior = n
		n = t.nos[n].prox
	}
	return Chamado{}, false
}
