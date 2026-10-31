package main

// no guarda um chamado e o índice do próximo nó da msm caixa
type no struct {
	prox    int     // aponta pro prox no
	chamado Chamado //
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

// guarda o chamado c na tabela, no início da fila da sua caixa
func (t *tabelaChamados) Inserir(c Chamado) {
	if t.livre == -1 {
		return
	}
	n := t.livre            // pega o primeiro nó livre
	t.livre = t.nos[n].prox // a fila de livres avança
	caixa := posicao(c.ID)
	t.nos[n].chamado = c
	t.nos[n].prox = t.chaves[caixa] // a fila antiga fica pendurada no nó novo
	t.chaves[caixa] = n             // e a caixa passa a apontar para ele
	t.tamanho++
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
