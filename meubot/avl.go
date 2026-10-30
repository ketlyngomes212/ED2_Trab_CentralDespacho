package main

type noAvl struct {
	chave int 
	alt int 
	esquerda *noAvl
	direita *noAvl
}

func maior(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func alturaAVL(a *noAvl) int {
	if a == nil {
		return -1
	}
	return a.alt
	
}

func atualizarAltura(a *noAvl) {
	a.alt = 1 + maior(alturaAVL(a.direita), alturaAVL(a.esquerda))
}

func fator(a *noAvl) int {
	if a == nil {
		return 0
	}

	return alturaAVL(a.esquerda) - alturaAVL(a.direita)
}

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