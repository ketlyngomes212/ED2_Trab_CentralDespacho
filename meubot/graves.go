package main

//graves.go trata da lista de chamados graves que estão em aberto

type ListaGraves struct {
	nos []int
}

func novaLista() *ListaGraves {
	return &ListaGraves{nos: make([]int, 0, 65536)}
}

//registra o nó como um chamado grave
func (l *ListaGraves) Inserir(n int) {
	tab.nos[n].posGrave = len(l.nos)
	l.nos = append(l.nos, n)
}

//tira o nó da lista e troca ele pelo último
func (l *ListaGraves) Remover(n int) {
	pos := tab.nos[n].posGrave
	ultimo := len(l.nos) - 1
	mudou := l.nos[ultimo]
	l.nos[pos] = mudou
	tab.nos[mudou].posGrave = pos
	l.nos = l.nos[:ultimo]
}
