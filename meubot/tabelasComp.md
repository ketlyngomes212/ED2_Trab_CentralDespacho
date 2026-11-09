## tabela comparativa

### Cenário Pequeno, etapa 1, semente 1: 500 turnos, 1000 chamados, 0.2 s, limite 100 ms (relogio)

| Modelo | Pontos | recolhidos | entregas | infrações | estouros | Média (ms) | p99 (ms) | Máx (ms) |
|---|---|---|---|---|---|---|---|---|
| ingenuo | 718 | 71 | 68 | 0 | 0 | 0.03 | 0.53 | 0.58 |
| avl, hash, heap, buscando + prox | 754 | 73 | 69 | 0 | 0 | 0.03 | 0.57 | 1.16 |
| taxa de retorno e lista de graves | 1263 | 49 | 49 | 0 | 0 | 0.04 | 0.54 | 0.54  
| Indexado | 719 | 36 | 34 | 0 | 0 | 0.04 | 0.86 | 1.86 |


### Cenário médio , etapa 1, semente 1: 1000 turnos, 39997 chamados, 0.7 s, limite 100 ms (relogio)

| Modelo | Pontos | recolhidos | entregas | infrações | estouros | P Média (ms) | P p99 (ms) | P Máx (ms) |
|---|---|---|---|---|---|---|---|---|
| ingenuo | 4804 | 437 | 431 | 0 | 0 | 0.55 | 1.12 | 6.76 |
| avl, hash, heap, buscando + prox| 5156 | 451 | 443 | 0 | 0 | 0.09 | 0.61 | 0.78 |
| taxa de retorno e lista de graves | 10168 | 289 | 289 | 0 | 0 | 0.14 | 0.56 | 1.13
| Indexado | 8633 | 437 | 219 | 0 | 0 | 0.07 | 0.63 | 0.94 |



### Cenário Estresse etapa 1, semente 1: 1000 turnos, 996846 chamados, 94.0 s, limite 100 ms (relogio)

| Modelo | Pontos | recolhidos | entregas | infrações | estouros | P Média (ms) | P p99 (ms) | P Máx (ms) |
|---|---|---|---|---|---|---|---|---|
| ingenuo | 20091 | 1676 | 1633 | 0 | 478 | 55.56 | 10.33 | 36.42 |
| avl, hash, heap, buscando + prox| 39390 | 3265 | 3218 | 0 | 0 | 3.31 | 10.26 | 28.04 |
| taxa de retorno e lista de graves | 61038 | 2134 | 2134 | 0 | 0 | 6.27 | 14.42 | 78.28
| Indexado | 59709 | 2031 | 1992 | 0 | 0 | 3.07 | 21.50 | 72.21 |