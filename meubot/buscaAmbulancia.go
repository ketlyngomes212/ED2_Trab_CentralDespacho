package main

// escolhe o chamado mais próximo que ainda vale a pena para uma ambulância livre, usando o índice espacial.
// 1- buscarMelhor olha a linha da ambulância, depois uma acima e uma abaixo, depois duas acima e duas abaixo, e assim por diante;
// 2- melhorNaLinha, em cada linha, compara o melhor chamado à direita e o melhor à esquerda da ambulância;
// 3- vizinhoDireita e vizinhoEsquerda andam pela árvore da linha, a partir da coluna da ambulância, pulando os chamados que não servem.

// diz se o chamado c pode ser escolhido pela ambulância em (ax, ay) no turno
func chamadoServe(c Chamado, ax, ay, turno int, atribuidos []int) bool {
	return !chamadoOcupado(atribuidos, c.ID) && turno+calculaDistancia(c.X, c.Y, ax, ay) <= c.Prazo
}

// devolve, na linha y, o chamado mais próximo que serve e que esteja na coluna ax ou à direita dela. Só aceita chamados a uma distância menor que limite (a melhor distância já encontrada). Devolve false se não houver nenhum
func (e *indiceEspacial) vizinhoDireita(y, ax, ay, turno, limite int, atribuidos []int, tab *tabelaChamados) (Chamado, bool) {
	raiz := e.linhas[y]
	// chave(ax, 0) fica antes de qualquer chamado com x == ax
	k, ok := sucessorOuIgual(raiz, chave(ax, 0))
	for ok {
		if calculaDistancia(xDaChave(k), y, ax, ay) >= limite {
			return Chamado{}, false // este e todos os seguintes estão longe demais
		}
		// A árvore só guarda a chave; os demais dados vêm da tabela hash
		c, achou := tab.Buscar(idDaChave(k))
		if achou && chamadoServe(c, ax, ay, turno, atribuidos) {
			return c, true
		}
		//  Próxima chave à direita; pergunta-se por k+1, porque sucessorOuIgual(k) devolveria a própria chave k
		k, ok = sucessorOuIgual(raiz, k+1)
	}
	return Chamado{}, false
}

// é o espelho de vizinhoDireita: o chamado mais próximo que serve, na coluna ax ou à esquerda dele
func (e *indiceEspacial) vizinhoEsquerda(y, ax, ay, turno, limite int, atribuidos []int, tab *tabelaChamados) (Chamado, bool) {
	raiz := e.linhas[y]
	k, ok := antecessorOuIgual(raiz, chave(ax, deslocamento-1))
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

// devolve o chamado mais próximo que serve na linha y, olhando os dois lados da ambulância. Em caso de empate de distância, fica com o da direita
func (e *indiceEspacial) melhorNaLinha(y, ax, ay, turno, limite int, atribuidos []int, tab *tabelaChamados) (Chamado, bool) {
	dir, okDir := e.vizinhoDireita(y, ax, ay, turno, limite, atribuidos, tab)
	esq, okEsq := e.vizinhoEsquerda(y, ax, ay, turno, limite, atribuidos, tab)

	if !okDir && !okEsq {
		return Chamado{}, false
	}
	if !okEsq {
		return dir, true
	}
	if !okDir {
		return esq, true
	}

	if calculaDistancia(dir.X, dir.Y, ax, ay) <= calculaDistancia(esq.X, esq.Y, ax, ay) {
		return dir, true
	}
	return esq, true
}

// devolve o chamado mais próximo que serve para a ambulância em (ax, ay), ou false se nenhum servir
func (e *indiceEspacial) buscarMelhor(ax, ay, turno int, atribuidos []int, tab *tabelaChamados) (Chamado, bool) {
	altura := len(e.linhas)
	limite := 1 << 30 // enorme: ainda não existe candidato
	var melhor Chamado
	achou := false

	for d := 0; d < limite; d++ {
		acima := ay - d
		abaixo := ay + d
		if acima < 0 && abaixo >= altura {
			break // as duas linhas saíram do mapa
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
