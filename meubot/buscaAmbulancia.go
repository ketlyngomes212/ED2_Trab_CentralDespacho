package main

// Escolhe o chamado mais próximo que ainda vale a pena para uma ambulância
// livre, usando um índice espacial.
// 1- buscarMelhor olha a linha da ambulância, depois uma acima e uma abaixo,
//    depois duas acima e duas abaixo, e assim por diante;
// 2- melhorNaLinha, em cada linha, compara o melhor à direita e o melhor à esquerda;
// 3- vizinhoDireita e vizinhoEsquerda andam pela árvore da linha, a partir da
//    coluna da ambulância, pulando os chamados que não servem.

func chamadoOcupado(ids []int, id int) bool {
	for _, atual := range ids {
		if atual == id {
			return true
		}
	}
	return false
}

// chamadoServe diz se o chamado c pode ser escolhido pela ambulância em
// (ax, ay) no turno. A ambulância chega no turno turno+dist-1.
func chamadoServe(c Chamado, ax, ay, turno int, atribuidos []int) bool {
	return !chamadoOcupado(atribuidos, c.ID) &&
		turno+calculaDistancia(c.X, c.Y, ax, ay)-1 <= c.Prazo
}

// vizinhoDireita devolve, na linha y, o chamado mais próximo que serve e que
// esteja na coluna ax ou à direita. Só aceita distância menor que limite.
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

// vizinhoEsquerda é o espelho: o chamado mais próximo que serve na coluna
// ax-1 ou à esquerda dela.
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

// melhorNaLinha olha os dois lados da ambulância. Empate: fica o da direita.
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

// buscarMelhor devolve o chamado mais próximo que serve para a ambulância em
// (ax, ay), a menos de limiteInicial células, ou false se nenhum servir.
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
