//graves.go controla os índices de graves e possui a mecânica de escolha do alvo

package main

const alcanceGrave = 200 //raio máximo dos graves (acima disso não é interessante mais)
const pesoGrave = 1      //fator de favorecimento dos graves

var indiceGrave *indiceEspacial

func inserirNoIndice(c Chamado) {
	if ehGrave(c) {
		indiceGrave.Inserir(c)
	} else {
		indice.Inserir(c)
	}
}

//removerDoIndice decide o índice pelo valor do chamado devolvido pela tabela
func removerDoIndice(c Chamado) {
	if ehGrave(c) {
		indiceGrave.Remover(c)
	} else {
		indice.Remover(c)
	}
}

//distanciaHospital analisa as distâncias entre os hospitais de acordo com uma posição
func distanciaHospital(x, y int, hospitais []Hospital) int {
	melhor := 1 << 30
	for _, h := range hospitais {
		if d := calculaDistancia(x, y, h.X, h.Y); d < melhor {
			melhor = d
		}
	}
	return melhor
}

//escolherAlvo compara o melhor grave com o melhor comum pelo valor por turno gasto
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

	if pesoGrave*g.Valor*custoC >= c.Valor*custoG {
		return g, true
	}
	return c, true
}
