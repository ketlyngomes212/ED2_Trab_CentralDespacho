package main

// Escolha do alvo de uma ambulância livre.
//
// Critério: TAXA DE RETORNO = valor / tempo da viagem
//   taxa = valor / (distância até o chamado + 0,75 * distância do chamado ao hospital + 1)
// O valor só vira ponto na entrega, então importa quantos pontos a ambulância
// rende por turno gasto. Esse único número equilibra graves e comuns sem regra
// fixa de prioridade. O peso 0,75 na volta ao hospital veio de testes: a volta
// é certa, enquanto a ida arrisca perder o chamado para o adversário.
//
// Onde procurar:
//  1. graves: percorre a listaGraves inteira;
//  2. comuns: percorre a AVL por linha em losangos de raio crescente (16, 32, ...)
//     e para quando nada mais longe pode ganhar do melhor achado

const (
	valorMaxComum    = 20   // maior valor possível de um chamado comum
	histerese        = 1.25 // bônus do alvo atual, para não trocar de alvo todo turno
	pesoHospital     = 0.75 // peso da distância chamado -> hospital na taxa
	orcamentoVisitas = 600  // máx. de comuns avaliados por ambulância/turno, depois de achar algum
	orcamentoMaximo  = 3000 // máx. de comuns avaliados por ambulância/turno, mesmo sem achar nenhum
	raioInicial      = 16
	raioMaximo = 1 << 30 
)

// estado de uma busca de alvo para UMA ambulância
type Busca struct {
	ax, ay  int
	turno   int
	amb     int // id da ambulância
	atual   int // id do alvo atual da ambulância (0 se não tem)
	melhor  *Chamado
	taxa    float64
	visitas int
}

// considera o chamado c como alvo e guarda-o se for o melhor visto até agora
func (b *Busca) avaliar(c *Chamado) {
	if c.Dona != -1 && c.Dona != b.amb {
		return // outra ambulância nossa já vai para ele
	}
	d := calculaDistancia(c.X, c.Y, b.ax, b.ay)
	if b.turno+d > c.Prazo {
		return // não dá tempo de chegar
	}
	if b.turno+d+c.DistHospital > totalTurnos {
		return // a entrega não acontece antes do fim da partida
	}
	t := float64(c.Valor) / (float64(d) + pesoHospital*float64(c.DistHospital) + 1)
	if c.ID == b.atual {
		t *= histerese
	}
	if t > b.taxa {
		b.taxa = t
		b.melhor = c
	}
}

// avalia todos os graves abertos
// avalia os graves abertos, pulando antes de qualquer conta cara os que não
// podem ganhar do melhor atual
func (b *Busca) buscarGraves(l *ListaGraves) {
	for i := range l.itens {
		g := &l.itens[i]
		dx := g.x - b.ax
		if dx < 0 {
			dx = -dx
		}
		dy := g.y - b.ay
		if dy < 0 {
			dy = -dy
		}
		d := dx + dy
		if d > raioMaximo || b.turno+d > g.prazo {
			continue // longe demais, ou não dá tempo de chegar
		}
		if b.melhor != nil && float64(g.valor) <= b.taxa*float64(d+1) {
			continue // nem com o hospital colado ganharia do melhor atual
		}
		b.avaliar(&tab.nos[g.no].chamado)
	}
}

// avalia os comuns da linha y com coluna entre xmin e xmax
func (e *indiceEspacial) visitarFaixa(b *Busca, y, xmin, xmax int) {
	if xmin < 0 {
		xmin = 0
	}
	if xmax >= e.largura {
		xmax = e.largura - 1
	}
	if xmin > xmax {
		return
	}
	ady := y - b.ay
	if ady < 0 {
		ady = -ady
	}
	percorrerIntervalo(e.linhas[y], chave(xmin, 0), chave(xmax, deslocamento-1), func(k int) {
		// Poda antes de consultar a tabela hash, só com o que a chave já diz (x e y):
		dx := xDaChave(k) - b.ax
		if dx < 0 {
			dx = -dx
		}
		d := dx + ady
		if b.turno+d > totalTurnos {
			return // nem chegando lá sobra tempo de entregar
		}
		if b.melhor != nil && float64(valorMaxComum) <= b.taxa*float64(d+1) {
			return // nem um comum de valor máximo ganharia do melhor atual
		}
		n := tab.AcharNo(idDaChave(k))
		if n != -1 {
			b.visitas++
			b.avaliar(&tab.nos[n].chamado)
		}
	})
}

// avalia os comuns em losangos (distância de Manhattan <= raio) cada vez maiores.
// Em cada passo só percorre a coroa nova (entre o losango anterior e o atual).
func (e *indiceEspacial) buscarComuns(b *Busca) {
	altura := len(e.linhas)
	raioAnt := -1
	for raio := raioInicial; ; raio *= 2 {
		for dy := -raio; dy <= raio; dy++ {
			y := b.ay + dy
			if y < 0 || y >= altura {
				continue
			}
			ady := dy
			if ady < 0 {
				ady = -ady
			}
			meio := raio - ady
			interior := raioAnt - ady // metade da largura do losango anterior nesta linha
			if interior < 0 {
				e.visitarFaixa(b, y, b.ax-meio, b.ax+meio)
			} else {
				e.visitarFaixa(b, y, b.ax-meio, b.ax-interior-1)
				e.visitarFaixa(b, y, b.ax+interior+1, b.ax+meio)
			}
		}

		// nada além do raio consegue taxa maior que valorMaxComum/(raio+1)
		if b.taxa >= float64(valorMaxComum)/float64(raio+1) {
			return
		}
		if b.melhor != nil && b.visitas >= orcamentoVisitas {
			return
		}
		if b.visitas >= orcamentoMaximo {
			return
		}
		if raio >= altura+e.largura {
			return // o losango já cobre o mapa inteiro
		}
		if b.turno+raio >= totalTurnos {
			return // mais longe que isso não dá tempo nem de chegar antes do fim
		}
		if raio >= raioMaximo {
			return
		}
		raioAnt = raio
	}
}
