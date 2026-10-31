//O estrategia.go determina de que maneira uma ambulância deve agir ao longo da partida.

package main

// chamadoOcupado verifica se determinado chamado já foi escolhido por outra ambulância no turno
func chamadoOcupado(ids []int, id int) bool {
	for _, atual := range ids {
		if atual == id {
			return true
		}
	}
	return false
}

// escolherChamado seleciona o chamado mais próximo
// o vetor de atribuidos garante que uma ambulância não siga para um chamado já escolhido por outra
func escolherChamado(ambulancia Ambulancia, chamados []Chamado, atribuidos []int,
	turno int) (Chamado, bool) {

	var escolhido Chamado
	encontrou := false
	menorDistancia := 0

	for _, chamado := range chamados {
		if chamadoOcupado(atribuidos, chamado.ID) { //se o chamado já foi escolhido, continua
			continue
		}

		dist := calculaDistancia(ambulancia.X, ambulancia.Y, chamado.X, chamado.Y)
		if turno+dist > chamado.Prazo { //se não vai dar tempo de chegar, continua
			continue
		}

		if !encontrou || dist < menorDistancia ||
			(dist == menorDistancia && chamado.Valor > escolhido.Valor) {
			escolhido = chamado
			encontrou = true
			menorDistancia = dist
		}
	}

	return escolhido, encontrou
}

// escolherHospital seleciona o hospital mais próximo
func escolherHospital(ambulancia Ambulancia, hospitais []Hospital) (Hospital, bool) {
	var escolhido Hospital
	encontrou := false
	menorDistancia := 0

	for _, hospital := range hospitais {
		dist := calculaDistancia(ambulancia.X, ambulancia.Y, hospital.X, hospital.Y)

		if !encontrou || dist < menorDistancia {
			escolhido = hospital
			encontrou = true
			menorDistancia = dist
		}
	}

	return escolhido, encontrou
}

// proximoMovimento define para onde a ambulância vai
func proximoMovimento(ambulancia Ambulancia, hospitais []Hospital,
	atribuidos []int, turno int) (string, int) {

	//Ambulância livre vai para o chamado mais próximo
	if ambulancia.Livre {
		chamado, encontrou := indice.buscarMelhor(ambulancia.X, ambulancia.Y, turno, atribuidos, tab)

		if !encontrou {
			return "P", -1
		}

		movimento := direcao(ambulancia.X, ambulancia.Y, chamado.X, chamado.Y)

		return movimento, chamado.ID
	}

	//Ambulância ocupada vai para o hospital mais próximo
	hospital, encontrou := escolherHospital(ambulancia, hospitais)

	if !encontrou {
		return "P", -1
	}

	movimento := direcao(ambulancia.X, ambulancia.Y, hospital.X, hospital.Y)

	return movimento, -1
}
