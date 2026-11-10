//buscaAmbulancia.go realiza a busca pelo chamado mais próximo

package main

func chamadoOcupado(ids []int, id int) bool {
	for _, atual := range ids {
		if atual == id {
			return true
		}
	}
	return false
}

//chamadoServe diz se o chamado c pode ser atribuído pela ambulância
func chamadoServe(c Chamado, ax, ay, turno int, atribuidos []int) bool {
	return !chamadoOcupado(atribuidos, c.ID) &&
		turno+calculaDistancia(c.X, c.Y, ax, ay)-1 <= c.Prazo
}

//vizinhoDireita devolve, na linha y, o chamado mais próximo que serve e que esteja na coluna ax ou à direita
func (e *indiceEspacial) vizinhoDireita(y, ax, ay, turno, limite int, atribuidos []int, tab *tabelaChamados) (Chamado, bool) {
	raiz := e.linhas[y]
	k, ok := sucessorOuIgual(raiz, chave(ax, 0))
	for ok {
		if calculaDistancia(xDaChave(k), y, ax, ay) >= limite {
			return Chamado{}, false
		}
		c, achou := tab.Buscar(idDaChave(k))
		if achou && chamadoServe(c, ax, ay, turno, atribuidos) {
			return c, true
		}
		k, ok = sucessorOuIgual(raiz, k+1)
	}
	return Chamado{}, false
}

//vizinhoEsquerda é o espelho: o chamado mais próximo que serve na coluna ax-1 ou à esquerda dela.
func (e *indiceEspacial) vizinhoEsquerda(y, ax, ay, turno, limite int, atribuidos []int, tab *tabelaChamados) (Chamado, bool) {
	raiz := e.linhas[y]
	k, ok := antecessorOuIgual(raiz, chave(ax, 0)-1)
	for ok {
		if calculaDistancia(xDaChave(k), y, ax, ay) >= limite {
			return Chamado{}, false
		}
		c, achou := tab.Buscar(idDaChave(k))
		if achou && chamadoServe(c, ax, ay, turno, atribuidos) {
			return c, true
		}
		k, ok = antecessorOuIgual(raiz, k-1)
	}
	return Chamado{}, false
}

//melhorNaLinha compara o melhor candidato de cada lado e define o melhor entre os dois
func (e *indiceEspacial) melhorNaLinha(y, ax, ay, turno, limite int, atribuidos []int, tab *tabelaChamados) (Chamado, bool) {
	cd, okD := e.vizinhoDireita(y, ax, ay, turno, limite, atribuidos, tab)
	ce, okE := e.vizinhoEsquerda(y, ax, ay, turno, limite, atribuidos, tab)

	if !okD && !okE {
		return Chamado{}, false
	}
	if !okE {
		return cd, true
	}
	if !okD {
		return ce, true
	}
	if calculaDistancia(cd.X, cd.Y, ax, ay) <= calculaDistancia(ce.X, ce.Y, ax, ay) {
		return cd, true
	}
	return ce, true
}

//buscarMelhor realiza a expansão linha a linha, quando encontra um candidato, define ele como limite
func (e *indiceEspacial) buscarMelhor(ax, ay, turno, limiteInicial int, atribuidos []int, tab *tabelaChamados) (Chamado, bool) {
	altura := len(e.linhas)
	limite := limiteInicial
	var melhor Chamado
	achou := false

	for d := 0; d < limite; d++ {
		acima := ay - d
		abaixo := ay + d
		if acima < 0 && abaixo >= altura {
			break
		}

		if acima >= 0 {
			c, ok := e.melhorNaLinha(acima, ax, ay, turno, limite, atribuidos, tab)
			if ok {
				melhor, achou = c, true
				limite = calculaDistancia(melhor.X, melhor.Y, ax, ay)
			}
		}

		if d > 0 && abaixo < altura {
			c, ok := e.melhorNaLinha(abaixo, ax, ay, turno, limite, atribuidos, tab)
			if ok {
				melhor, achou = c, true
				limite = calculaDistancia(melhor.X, melhor.Y, ax, ay)
			}
		}
	}
	return melhor, achou
}
