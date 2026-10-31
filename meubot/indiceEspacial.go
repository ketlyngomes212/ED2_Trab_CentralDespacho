package main

//índice para achar chamados perto de uma posição
// O mapa é uma grade, e cada linha y tem a sua própria AVL, com chave x*2^32 + id. Para saber o que existe perto de uma ambulância, olha-se a árvore da linha dela e as das linhas vizinhas, e em cada uma busca-se o chamado de x mais próximo.

// reúne as árvores de todas as linhas.
type indiceEspacial struct {
	linhas  []*noAvl // linhas[y] é a raiz da AVL da linha y
	tamanho int
}

// cria o índice para um mapa com o número de linhas dado
func novoIndice(altura int) *indiceEspacial {
	return &indiceEspacial{linhas: make([]*noAvl, altura)}
}

// guarda o chamado na árvore da sua linha. A atribuição é necessária pq a inserção pode trocar a raiz da árvore (por causa de uma rotação).
func (e *indiceEspacial) Inserir(c Chamado) {
	e.linhas[c.Y] = inserir(e.linhas[c.Y], chave(c.X, c.ID))
	e.tamanho++
}

// tira o chamado da árvore da sua linha
func (e *indiceEspacial) Remover(c Chamado) {
	e.linhas[c.Y] = remover(e.linhas[c.Y], chave(c.X, c.ID))
	e.tamanho--
}
