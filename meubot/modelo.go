//modelo.go contém todas as representações dos elementos do jogo

package main

type Chamado struct {
	ID    int
	X     int
	Y     int
	Prazo int
	Valor int
}

//ehGrave verifica se um chamado é grave
const MinGrave = 40

func ehGrave(c Chamado) bool {
	return c.Valor >= MinGrave
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
