package main

// a chave que as AVLs por linha usam para guardar os chamados
// Cada nó da AVL guarda um único número inteiro, mas precisamos guardar vários
// chamados por linha, inclusive vários na mesma coluna x. A solução é juntar
// x e id num número só (x * deslocamento + id)

// deslocamento é 2 elevado a 32
const deslocamento = 4294967296

// junta a coluna x e o id do chamado num único número
func chave(x, id int) int {
	return x*deslocamento + id
}

// devolve a coluna x guardada numa chave
func xDaChave(k int) int {
	return k / deslocamento
}

// devolve o id do chamado guardado numa chave.
func idDaChave(k int) int {
	return k % deslocamento
}
