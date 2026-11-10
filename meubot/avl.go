//avl.go implementa a AVL utilizada nas buscas por chamados

package main

type noAvl struct {
	chave int
	alt   int
	esq   *noAvl
	dir   *noAvl
}

func alturaNo(n *noAvl) int {
	if n == nil {
		return 0
	}
	return n.alt
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func atualizar(n *noAvl) {
	n.alt = 1 + maxInt(alturaNo(n.esq), alturaNo(n.dir))
}

//fator de balanceamento da AVL
func fatorBalanceamento(n *noAvl) int {
	return alturaNo(n.esq) - alturaNo(n.dir)
}

//rotações da AVL
func rotacaoDireita(y *noAvl) *noAvl {
	x := y.esq
	y.esq = x.dir
	x.dir = y
	atualizar(y)
	atualizar(x)
	return x
}

func rotacaoEsquerda(x *noAvl) *noAvl {
	y := x.dir
	x.dir = y.esq
	y.esq = x
	atualizar(x)
	atualizar(y)
	return y
}

//faz o balanceamento da AVL
func balancear(n *noAvl) *noAvl {
	atualizar(n)
	fb := fatorBalanceamento(n)
	if fb > 1 {
		if fatorBalanceamento(n.esq) < 0 {
			n.esq = rotacaoEsquerda(n.esq)
		}
		return rotacaoDireita(n)
	}
	if fb < -1 {
		if fatorBalanceamento(n.dir) > 0 {
			n.dir = rotacaoDireita(n.dir)
		}
		return rotacaoEsquerda(n)
	}
	return n
}

//inserir devolve a nova raiz da árvore depois da inserção de uma chave k
func inserir(n *noAvl, k int) *noAvl {
	if n == nil {
		return &noAvl{chave: k, alt: 1}
	}
	if k < n.chave {
		n.esq = inserir(n.esq, k)
	} else if k > n.chave {
		n.dir = inserir(n.dir, k)
	} else {
		return n
	}
	return balancear(n)
}

func menorNo(n *noAvl) *noAvl {
	for n.esq != nil {
		n = n.esq
	}
	return n
}

//remover devolve a nova raiz da árvore depois de remover uma chave k
func remover(n *noAvl, k int) *noAvl {
	if n == nil {
		return nil
	}
	if k < n.chave {
		n.esq = remover(n.esq, k)
	} else if k > n.chave {
		n.dir = remover(n.dir, k)
	} else {
		if n.esq == nil {
			return n.dir
		}
		if n.dir == nil {
			return n.esq
		}
		s := menorNo(n.dir)
		n.chave = s.chave
		n.dir = remover(n.dir, s.chave)
	}
	return balancear(n)
}

//sucessorOuIgual devolve a menor chave >= k.
func sucessorOuIgual(n *noAvl, k int) (int, bool) {
	res := 0
	achou := false
	for n != nil {
		if n.chave == k {
			return k, true
		}
		if n.chave > k {
			res = n.chave
			achou = true
			n = n.esq
		} else {
			n = n.dir
		}
	}
	return res, achou
}

//antecessorOuIgual devolve a maior chave <= k.
func antecessorOuIgual(n *noAvl, k int) (int, bool) {
	res := 0
	achou := false
	for n != nil {
		if n.chave == k {
			return k, true
		}
		if n.chave < k {
			res = n.chave
			achou = true
			n = n.dir
		} else {
			n = n.esq
		}
	}
	return res, achou
}
