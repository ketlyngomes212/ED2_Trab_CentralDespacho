//O estrategia.go determina de que maneira uma ambulância deve agir ao longo da partida.

package main

// contemID verifica se um ID já está presente na lista.
func contemID(ids []int, id int) bool {
	for _, atual := range ids {
		if atual == id {
			return true
		}
	}
	return false
}

// escolherChamado seleciona o chamado viável mais próximo.
// Chamados já atribuídos a outra ambulância neste turno são ignorados.
func escolherChamado(
	ambulancia Ambulancia,
	chamados []Chamado,
	atribuidos []int,
	turno int,
) (Chamado, bool) {
	var melhor Chamado
	encontrou := false
	menorDistancia := 0

	for _, chamado := range chamados {
		if contemID(atribuidos, chamado.ID) {
			continue
		}

		dist := calculaDistancia(
			ambulancia.X, ambulancia.Y,
			chamado.X, chamado.Y,
		)

		// Só consideramos chamados que podem ser alcançados
		// até o prazo informado pelo servidor.
		if turno+dist > chamado.Prazo {
			continue
		}

		if !encontrou || dist < menorDistancia ||
			(dist == menorDistancia && chamado.Valor > melhor.Valor) {
			melhor = chamado
			menorDistancia = dist
			encontrou = true
		}
	}

	return melhor, encontrou
}

// escolherHospital encontra o hospital mais próximo.
func escolherHospital(
	ambulancia Ambulancia,
	hospitais []Hospital,
) (Hospital, bool) {
	var melhor Hospital
	encontrou := false
	menorDistancia := 0

	for _, hospital := range hospitais {
		dist := calculaDistancia(
			ambulancia.X, ambulancia.Y,
			hospital.X, hospital.Y,
		)

		if !encontrou || dist < menorDistancia {
			melhor = hospital
			menorDistancia = dist
			encontrou = true
		}
	}

	return melhor, encontrou
}

// decidirMovimento escolhe o próximo movimento da ambulância.
// Retorna também o ID do chamado escolhido, ou -1 se nenhum foi escolhido.
func decidirMovimento(
	ambulancia Ambulancia,
	chamados []Chamado,
	hospitais []Hospital,
	atribuidos []int,
	turno int,
) (string, int) {
	if ambulancia.Livre {
		chamado, encontrou := escolherChamado(
			ambulancia, chamados, atribuidos, turno,
		)

		if !encontrou {
			return "P", -1
		}

		movimento := direcao(
			ambulancia.X, ambulancia.Y,
			chamado.X, chamado.Y,
		)

		return movimento, chamado.ID
	}

	// Ambulância ocupada: seguir para o hospital mais próximo.
	hospital, encontrou := escolherHospital(ambulancia, hospitais)

	if !encontrou {
		return "P", -1
	}

	movimento := direcao(
		ambulancia.X, ambulancia.Y,
		hospital.X, hospital.Y,
	)

	return movimento, -1
}
