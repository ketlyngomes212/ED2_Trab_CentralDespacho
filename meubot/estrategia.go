package main

import "time"

// alvos[i] é o id do chamado para o qual a ambulância i está indo (0 = nenhum).
// O compromisso evita que a ambulância troque de alvo a cada turno.
var alvos []int

// Rede de segurança contra o teto de 100 ms por turno: se as buscas já gastaram
// este tempo, as ambulâncias que faltam só mantêm o alvo que já tinham.
const tempoMaximoBusca = 55 * time.Millisecond
var tempoGraves, tempoComuns time.Duration

// instante em que o turno atual começou a ser processado
var inicioTurno time.Time

// libera a reserva do alvo da ambulância
func liberarAlvo(amb int) {
	id := alvos[amb]
	if id == 0 {
		return
	}
	c := tab.Achar(id)
	if c != nil && c.Dona == amb {
		c.Dona = -1
	}
	alvos[amb] = 0
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

// distância do ponto (x, y) ao hospital mais próximo
func distanciaHospital(x, y int) int {
	menor := 1 << 30
	for _, h := range hospitais {
		d := calculaDistancia(x, y, h.X, h.Y)
		if d < menor {
			menor = d
		}
	}
	return menor
}

// proximoMovimento define para onde a ambulância vai
func proximoMovimento(ambulancia Ambulancia, turno int) string {
	for ambulancia.ID >= len(alvos) { // proteção: id fora do esperado (0 a k-1)
		alvos = append(alvos, 0)
	}

	// Ambulância ocupada vai para o hospital mais próximo
	if !ambulancia.Livre {
		liberarAlvo(ambulancia.ID)
		hospital, encontrou := escolherHospital(ambulancia, hospitais)
		if !encontrou {
			return "P"
		}
		return direcao(ambulancia.X, ambulancia.Y, hospital.X, hospital.Y)
	}

	// Ambulância livre vai para o chamado de maior taxa valor/tempo
	b := Busca{ax: ambulancia.X, ay: ambulancia.Y, turno: turno, amb: ambulancia.ID}

	// O alvo atual (se ainda existir) entra primeiro, com bônus de histerese
	if alvos[ambulancia.ID] != 0 {
		atual := tab.Achar(alvos[ambulancia.ID])
		if atual != nil {
			b.atual = atual.ID
			b.avaliar(atual)
		}
	}

	if time.Since(inicioTurno) < tempoMaximoBusca {
		t0 := time.Now()
		b.buscarGraves(graves)
		t1 := time.Now()
		indice.buscarComuns(&b)
		tempoGraves += t1.Sub(t0)
		tempoComuns += time.Since(t1)
	}

	if b.melhor == nil {
		liberarAlvo(ambulancia.ID)
		return "P"
	}

	if b.melhor.ID != alvos[ambulancia.ID] {
		liberarAlvo(ambulancia.ID)
		b.melhor.Dona = ambulancia.ID
		alvos[ambulancia.ID] = b.melhor.ID
	}

	return direcao(ambulancia.X, ambulancia.Y, b.melhor.X, b.melhor.Y)
}
