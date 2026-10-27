//O modelo.go contém todas as representações dos elementos do jogo.

package main

type Chamado struct {
	ID    int
	X     int
	Y     int
	Valor int
	Prazo int
}

type Ambulancia struct {
	ID    int
	X     int
	Y     int
	Livre bool
}

type Hospital struct {
	X int
	Y int
}
