//chaveAVL.go define a chave usada nas AVLs em cada linha

package main

//deslocamento é 2 elevado a 32
const deslocamento = 4294967296

//a chave é resultado da junção da coluna do chamado e do seu id
func chave(x, id int) int {
	return x*deslocamento + id
}

//xDaChave devolve a coluna x guardada numa chave
func xDaChave(k int) int {
	return k / deslocamento
}

//idDaChave devolve o id do chamado guardado numa chave
func idDaChave(k int) int {
	return k % deslocamento
}
