//indiceEspacial.go cria o índice para achar chamados perto de uma posição
//cada linha y do mapa tem a sua própria AVL, com chave x*2^32 + id

package main

type indiceEspacial struct {
	linhas  []*noAvl //linhas[y] é a raiz da AVL da linha y
	tamanho int
}

func novoIndice(altura int) *indiceEspacial {
	return &indiceEspacial{linhas: make([]*noAvl, altura)}
}

//Inserir guarda o chamado na árvore da sua linha. A atribuição é necessária
//porque a inserção pode trocar a raiz (rotação).
func (e *indiceEspacial) Inserir(c Chamado) {
	e.linhas[c.Y] = inserir(e.linhas[c.Y], chave(c.X, c.ID))
	e.tamanho++
}

//Remover tira o chamado da árvore da sua linha.
func (e *indiceEspacial) Remover(c Chamado) {
	e.linhas[c.Y] = remover(e.linhas[c.Y], chave(c.X, c.ID))
	e.tamanho--
}
