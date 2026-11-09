package main

// graves.go: lista dos chamados graves abertos. Além do índice do nó na tabela,
// guarda uma cópia dos campos que a busca usa para podar, para varrer a lista
// sem ir à tabela hash a cada grave.

type grave struct {
	no    int // índice do nó do chamado na tabela
	x, y  int
	valor int
	prazo int
}

type ListaGraves struct {
	itens []grave
}

func novaLista() *ListaGraves {
	return &ListaGraves{itens: make([]grave, 0, 65536)}
}

// Inserir registra o nó n como um chamado grave.
func (l *ListaGraves) Inserir(n int) {
	c := tab.nos[n].chamado
	tab.nos[n].posGrave = len(l.itens)
	l.itens = append(l.itens, grave{no: n, x: c.X, y: c.Y, valor: c.Valor, prazo: c.Prazo})
}

// Remover tira o nó da lista, trocando-o pelo último.
func (l *ListaGraves) Remover(n int) {
	pos := tab.nos[n].posGrave
	ultimo := len(l.itens) - 1
	mudou := l.itens[ultimo]
	l.itens[pos] = mudou
	tab.nos[mudou.no].posGrave = pos
	l.itens = l.itens[:ultimo]
}