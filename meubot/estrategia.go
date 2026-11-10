//estrategia.go define como cada ambulância vai jogar no turno

package main

//escolherHospital seleciona o hospital mais próximo
func escolherHospital(ambulancia Ambulancia, hospitais []Hospital) (Hospital, bool) {
	var escolhido Hospital
	encontrou := false
	melhor := 1 << 30

	for _, h := range hospitais {
		d := calculaDistancia(ambulancia.X, ambulancia.Y, h.X, h.Y)
		if d < melhor {
			melhor = d
			escolhido = h
			encontrou = true
		}
	}
	return escolhido, encontrou
}

//proximoMovimento define para onde a ambulância vai e qual chamado ela escolheu
func proximoMovimento(ambulancia Ambulancia, hospitais []Hospital,
	atribuidos []int, turno int) (string, int) {

	if ambulancia.Livre {
		chamado, encontrou := escolherAlvo(ambulancia.X, ambulancia.Y, turno, atribuidos, hospitais)
		if !encontrou {
			return "P", -1
		}
		return direcao(ambulancia.X, ambulancia.Y, chamado.X, chamado.Y), chamado.ID
	}

	//ambulância ocupada vai para o hospital mais próximo
	hospital, encontrou := escolherHospital(ambulancia, hospitais)
	if !encontrou {
		return "P", -1
	}
	return direcao(ambulancia.X, ambulancia.Y, hospital.X, hospital.Y), -1
}
