package main

// no guarda um chamado e o índice do próximo nó da msm caixa
type no struct {
	prox     int // aponta pro prox no
	posGrave int // caso o nó seja grave, guarda a posição dele na lista dos graves
	chamado  Chamado
}

// Estrutura da tabela hash dos chamados abertos indexada pelo id
type tabelaChamados struct {
	chaves  []int
	nos     []no
	livre   int
	tamanho int
}

const tamanhoTabela = 100003  // tamanhoTabela é o número de caixas
const capacidadeNos = 1000000 // é o tamanho do pool de nós

// posicao devolve a caixa em que o chamado de identificador id deve ficar
func posicao(id int) int {
	return id % tamanhoTabela
}

// cria a tabela vazia, todas as caixas vazias e todos os nós
func novaTabela() *tabelaChamados {
	t := &tabelaChamados{}
	t.chaves = make([]int, tamanhoTabela)
	t.nos = make([]no, capacidadeNos)

	for i := 0; i < len(t.chaves); i++ {
		t.chaves[i] = -1
	}

	for i := 0; i < len(t.nos); i++ {
		t.nos[i].prox = i + 1
	}

	t.nos[capacidadeNos-1].prox = -1
	t.livre = 0
	t.tamanho = 0
	return t
}

// devolve o índice do nó e reserva
func (t *tabelaChamados) Inserir(c Chamado) int {
	if t.livre == -1 {
		return -1
	}

	n := t.livre
	t.livre = t.nos[n].prox
	caixa := posicao(c.ID)
	c.Dona = -1
	t.nos[n].chamado = c
	t.nos[n].prox = t.chaves[caixa]
	t.chaves[caixa] = n
	t.tamanho++

	return n
}

// devolve o índice do nó do chamado de identificador id, ou -1 se não existir
func (t *tabelaChamados) AcharNo(id int) int {
	n := t.chaves[posicao(id)]

	for n != -1 {
		if t.nos[n].chamado.ID == id {
			return n
		}
		n = t.nos[n].prox
	}

	return -1
}

// devolve um ponteiro para o chamado ou nil
func (t *tabelaChamados) Achar(id int) *Chamado {
	n := t.AcharNo(id)

	if n == -1 {
		return nil
	}

	return &t.nos[n].chamado
}

// procura o chamado de identificador id
func (t *tabelaChamados) Buscar(id int) (Chamado, bool) {
	caixa := posicao(id)
	n := t.chaves[caixa]
	for n != -1 {
		if t.nos[n].chamado.ID == id {
			return t.nos[n].chamado, true
		}
		n = t.nos[n].prox
	}
	return Chamado{}, false
}

// tira da tabela o chamado de identificador id e o devolve, para que quem chamou possa retira-lo também de outras estruturas
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
