package main

type noAvl struct {
	chave    int // número guardado no nó (x*deslocamento + id >> ver chaveAVL.go)
	alt      int // altura da subárvore
	esquerda *noAvl
	direita  *noAvl
}

// maior devolve o maior entre dois inteiros.
func maior(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// devolve a altura guardada no nó
func alturaAVL(a *noAvl) int {
	if a == nil {
		return -1
	}
	return a.alt

}

// recalcula a altura de a a partir da altura dos filhos
func atualizarAltura(a *noAvl) {
	a.alt = 1 + maior(alturaAVL(a.direita), alturaAVL(a.esquerda))
}

// é o fator de balanceamento (altura da esquerda menos altura da direita)
func fator(a *noAvl) int {
	if a == nil {
		return 0
	}

	return alturaAVL(a.esquerda) - alturaAVL(a.direita)
}

// rotações para o balanceamento
func rotacaoDireita(a *noAvl) *noAvl {
	b := a.esquerda
	a.esquerda = b.direita
	b.direita = a

	atualizarAltura(a)
	atualizarAltura(b)
	return b
}

func rotacaoEsquerda(a *noAvl) *noAvl {
	b := a.direita
	a.direita = b.esquerda
	b.esquerda = a

	atualizarAltura(a)
	atualizarAltura(b)
	return b
}

// rebalancear a arvore para corrrigir o fator
func rebalancear(a *noAvl) *noAvl {
	f := fator(a)

	if f > 1 {
		if fator(a.esquerda) < 0 {
			a.esquerda = rotacaoEsquerda(a.esquerda)
		}
		return rotacaoDireita(a)
	}

	if f < -1 {
		if fator(a.direita) > 0 {
			a.direita = rotacaoDireita(a.direita)
		}
		return rotacaoEsquerda(a)
	}

	return a
}

// coloca a chave k na árvore de raiz a e devolve a nova raiz
// É a inserção da árvore de busca comum com uma diferença: na volta da
// recursão, cada nó do caminho atualiza a altura e se rebalanceia.
func inserir(a *noAvl, k int) *noAvl {
	if a == nil {
		return &noAvl{chave: k}
	}

	if k < a.chave {
		a.esquerda = inserir(a.esquerda, k)
	} else if k > a.chave {
		a.direita = inserir(a.direita, k)
	} else {
		return a
	}

	atualizarAltura(a)
	return rebalancear(a)
}

// tira a chave k da árvore de raiz a e devolve a nova raiz
// Como na inserção, cada nó do caminho de volta atualiza a altura e rebalanceia
func remover(a *noAvl, k int) *noAvl {
	if a == nil {
		return nil
	}

	if k < a.chave {
		a.esquerda = remover(a.esquerda, k)
	} else if k > a.chave {
		a.direita = remover(a.direita, k)
	} else {
		if a.esquerda == nil {
			return a.direita
		}
		if a.direita == nil {
			return a.esquerda
		}

		aux := a.esquerda
		for aux.direita != nil {
			aux = aux.direita
		}

		a.chave = aux.chave
		a.esquerda = remover(a.esquerda, aux.chave)
	}

	atualizarAltura(a)
	return rebalancear(a)
}

// devolve a menor chave que seja maior ou igual a k, e true.
// Se todas as chaves são menores que k, devolve false. Desce por um único caminho, guardando o melhor candidato visto.
func sucessorOuIgual(a *noAvl, k int) (int, bool) {
	melhor := 0
	achou := false
	for a != nil {
		if a.chave == k {
			return a.chave, true
		}
		if a.chave > k {
			melhor = a.chave
			achou = true
			a = a.esquerda
		} else {
			a = a.direita
		}
	}
	return melhor, achou
}

// devolve a maior chave que seja menor ou igual a k, e true.
func antecessorOuIgual(a *noAvl, k int) (int, bool) {
	melhor := 0
	achou := false
	for a != nil {
		if a.chave == k {
			return a.chave, true
		}
		if a.chave < k {
			melhor = a.chave
			achou = true
			a = a.direita
		} else {
			a = a.esquerda
		}
	}
	return melhor, achou
}
