package main

// Índice separado para os chamados graves (ehGrave, em modelo.go). Só ~5% dos chamados
// são graves, então buscar neles é barato. A tabela hash continua sendo a
// dona dos dados: os dois índices só guardam x e id.

const alcanceGrave = 200 // prazo dos graves é no máximo 150: nenhum serve além disso

// pesoGrave favorece levemente os graves na comparação valor/custo.
const pesoGrave = 1

var indiceGrave *indiceEspacial

func inserirNoIndice(c Chamado) {
	if ehGrave(c) {
		indiceGrave.Inserir(c)
	} else {
		indice.Inserir(c)
	}
}

// removerDoIndice decide o índice pelo Valor do chamado devolvido pela tabela,
// por isso remover por id (recolhido pela outra equipe) funciona.
func removerDoIndice(c Chamado) {
	if ehGrave(c) {
		indiceGrave.Remover(c)
	} else {
		indice.Remover(c)
	}
}

// distanciaHospital: poucos hospitais (até 20), varredura linear basta.
func distanciaHospital(x, y int, hospitais []Hospital) int {
	melhor := 1 << 30
	for _, h := range hospitais {
		if d := calculaDistancia(x, y, h.X, h.Y); d < melhor {
			melhor = d
		}
	}
	return melhor
}

// escolherAlvo compara o melhor grave com o melhor comum pelo valor por turno
// gasto (ir até o chamado + levar ao hospital).
func escolherAlvo(ax, ay, turno int, atribuidos []int, hospitais []Hospital) (Chamado, bool) {
	g, okG := indiceGrave.buscarMelhor(ax, ay, turno, alcanceGrave, atribuidos, tab)
	c, okC := indice.buscarMelhor(ax, ay, turno, 1<<30, atribuidos, tab)

	if !okG && !okC {
		return Chamado{}, false
	}
	if !okG {
		return c, true
	}
	if !okC {
		return g, true
	}

	custoG := calculaDistancia(g.X, g.Y, ax, ay) + distanciaHospital(g.X, g.Y, hospitais) + 1
	custoC := calculaDistancia(c.X, c.Y, ax, ay) + distanciaHospital(c.X, c.Y, hospitais) + 1

	// valorG/custoG >= valorC/custoC, sem divisão
	if pesoGrave*g.Valor*custoC >= c.Valor*custoG {
		return g, true
	}
	return c, true
}
