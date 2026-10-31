# Relatório do projeto
Nome: Ketlyn Gomes de Oliveira, Luigi Favaro
GRR: 20252054, 20254073

## Achar os chamados próximos de uma ambulância
**Estrutura: uma AVL por linha**

**Custo por turno:**
**Alternativa descartada: Grade por blocos, Hash**
 porque ...



## Descartar os chamados expirados
**Estrutura: heap mínimo por prazo**

**Custo por turno:**
**Alternativa descartada: percorrer todos os chamados a cada turno**
 porque ...

 ## Achar / Remover um chamado recolhido sabendo só o identificador
 **Estrutura: tabela hash com encadeamento**

**Custo por turno:**
**Alternativa descartada: rubro-negra, ??**
 porque ...

 
## tabela comparativa

### Cenário Pequeno, etapa 1, semente 1: 500 turnos, 1000 chamados, 0.2 s, limite 100 ms (relogio)

| Modelo | Pontos | recolhidos | entregas | infrações | estouros | Média (ms) | p99 (ms) | Máx (ms) |
|---|---|---|---|---|---|---|---|---|
| ingenuo | | | | |
| melhoria do ingenuo | 718 | 71 | 68 | 0 | 0 | 0.06 | 0.96 | 3.03 |
| avl, hash, heap, buscando + prox| 754 | 73 | 69 | 0 | 0 | 0.03 | 0.57 | 1.16 |
| avl, hash, heap, buscando chamado vermelho|
| Indexado | 719 | 36 | 34 | 0 | 0 |  0.04 | 0.86 | 1.86 |


### Cenário médio , etapa 1, semente 1: 1000 turnos, 39997 chamados, 0.7 s, limite 100 ms (relogio)

| Modelo | Pontos | recolhidos | entregas | infrações | estouros | P Média (ms) | P p99 (ms) | P Máx (ms) |
|---|---|---|---|---|---|---|---|---|
| ingenuo | | | | |
| melhoria do ingenuo | 4804 | 437 | 431 | 0 | 0 | 0.41 | 1.25 | 1.67 |
| avl, hash, heap, buscando + prox| 5156 | 451 | 443 | 0 | 0 | 0.09 | 0.61 | 0.78 |
| avl, hash, heap, buscando chamado vermelho|
| Indexado | 8633 | 437 | 219 | 0 | 0 | 0.07 | 0.63 | 0.94 |



### Cenário Estresse etapa 1, semente 1: 1000 turnos, 996846 chamados, 94.0 s, limite 100 ms (relogio)

| Modelo | Pontos | recolhidos | entregas | infrações | estouros | P Média (ms) | P p99 (ms) | P Máx (ms) |
|---|---|---|---|---|---|---|---|---|
| ingenuo | | | | |
| melhoria do ingenuo | 11302 | 973 | 930 | 0 | 731 | 56.59 | 196.27 | 244.58 |
| avl, hash, heap, buscando + prox| 39390 | 3265 | 3218 | 0 | 0 | 3.31 | 10.26 | 28.04 |
| avl, hash, heap, buscando chamado vermelho|
| Indexado | 59709 | 2031 | 1992 | 0 | 0 | 3.07 | 21.50 | 72.21 |