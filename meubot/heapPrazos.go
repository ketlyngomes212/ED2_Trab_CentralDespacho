package main

// Fila de prioridade (heap mínimo) ordenada pelo prazo dos chamados.

type item struct {
	prazo int
	id    int
}

type heapPrazos struct {
	itens []item
}

func pai(i int) int { return (i - 1) / 2 }
func esq(i int) int { return 2*i + 1 }
func dir(i int) int { return 2*i + 2 }

func (h *heapPrazos) subir(i int) {
	for i > 0 && h.itens[i].prazo < h.itens[pai(i)].prazo {
		h.itens[i], h.itens[pai(i)] = h.itens[pai(i)], h.itens[i]
		i = pai(i)
	}
}

func (h *heapPrazos) descer(i int) {
	n := len(h.itens)
	for esq(i) < n {
		menor := esq(i)
		if dir(i) < n && h.itens[dir(i)].prazo < h.itens[menor].prazo {
			menor = dir(i)
		}
		if h.itens[menor].prazo >= h.itens[i].prazo {
			break
		}
		h.itens[i], h.itens[menor] = h.itens[menor], h.itens[i]
		i = menor
	}
}

// Inserir acrescenta um item no fim e o faz subir até o lugar certo.
func (h *heapPrazos) Inserir(prazo, id int) {
	h.itens = append(h.itens, item{prazo: prazo, id: id})
	h.subir(len(h.itens) - 1)
}

// Topo devolve o item de menor prazo sem remover.
func (h *heapPrazos) Topo() (item, bool) {
	if len(h.itens) == 0 {
		return item{}, false
	}
	return h.itens[0], true
}

// RemoverTopo remove e devolve o item de menor prazo.
func (h *heapPrazos) RemoverTopo() (item, bool) {
	if len(h.itens) == 0 {
		return item{}, false
	}
	topo := h.itens[0]
	ultimo := len(h.itens) - 1
	h.itens[0] = h.itens[ultimo]
	h.itens = h.itens[:ultimo]
	h.descer(0)
	return topo, true
}

// expirar remove da tabela e dos índices todos os chamados cujo prazo já passou.
func expirar(turno int, h *heapPrazos, tab *tabelaChamados) {
	for {
		topo, ok := h.Topo()
		if !ok || topo.prazo >= turno {
			break
		}
		h.RemoverTopo()
		c, achou := tab.Remover(topo.id)
		if achou {
			removerDoIndice(c)
		}
	}
}
