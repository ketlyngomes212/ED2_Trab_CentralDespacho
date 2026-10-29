package main

// Estrutura do nó
type no struct {
	prox int // aponta pro prox no
	chamado Chamado //
}

// Estrutura da tabela
type tabelaChamados struct {
	chaves []int
	nos []no
	livre int
	tamanho int
}

const tamanhoTabela = 1024
const capacidadeNos = 2000   // pequeno só para testar

// func para descobrir posicao que o id cai (a primeira caixa)
func posicao(id int) int {
	return id % tamanhoTabela
}

//cria a tabela vazia e liga todos os nos livres
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

	t.nos[capacidadeNos - 1].prox = -1
	t.livre = 0
	t.tamanho = 0
	return t
}

// ponteiro para a struct de tabela / funçao inserir
func (t *tabelaChamados) Inserir(c Chamado) {
	if t.livre == -1 {
		return
	}
	n := t.livre 
	t.livre = t.nos[n].prox 
	caixa := posicao(c.ID)
	t.nos[n].chamado = c
	t.nos[n].prox =	t.chaves[caixa]
	t.chaves[caixa] = n
	t.tamanho++
}

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