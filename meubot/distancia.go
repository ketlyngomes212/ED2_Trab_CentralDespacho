//O distancia.go implementa a distância de Manhattan, fazendo a locomoção dos elementos
//durante a partida.

package main

func calculaDistancia(x1, y1, x2, y2 int) int {
	dx := x1 - x2
	if dx < 0 {
		dx = -dx
	}

	dy := y1 - y2
	if dy < 0 {
		dy = -dy
	}

	return dx + dy
}

func direcao(x, y, destinoX, destinoY int) string {
	if x < destinoX {
		return "L"
	}

	if x > destinoX {
		return "O"
	}

	if y < destinoY {
		return "S"
	}

	if y > destinoY {
		return "N"
	}

	return "P"
}
