# Relatório do projeto
Nome: Ketlyn Gomes de Oliveira, Luigi Favaro
GRR: 20252054, 20254073

## Achar os chamados próximos de uma ambulância
**Estruturas: duas AVL por linha do mapa (uma para chamados comuns, outra só para os graves), com os dados na tabela hash**
Cada linha y do mapa tem uma árvore AVL. A chave de cada chamado junta a coluna e o id num só inteiro 
 
**Custo por turno:**
1) Inserir e remover: O(log L), com L o número de chamados da linha.

2) Busca de uma ambulância: O((R + c) · log L), com R = linhas percorridas (no máximo a distância até o candidato encontrado) e c = chaves puladas por não servirem. Cada chave pulada custa uma nova consulta na árvore, O(log L), mais um acesso O(1) à hash. Linhas vazias custam O(1).

3) Por turno: uma busca nos comuns e uma nos graves para cada ambulância livre.

**Alternativa descartada: grade por blocos (hash espacial) E heap máximo por valor para os graves**

 
## Descartar os chamados expirados
**Estrutura: heap mínimo por prazo**

 
**Custo por turno:**
1) Inserir: O(log n) por chamado novo; olhar o topo: O(1).
2) Expirar: O(k log n), com k = chamados vencidos naquele turno. Os que não venceram não são tocados.
**Alternativa descartada: percorrer todos os chamados a cada turno**

 
## Achar e remover um chamado recolhido sabendo só o identificador
**Estrutura: tabela hash com encadeamento e pool de nós**

 
**Custo por turno:**
1) Inserir, buscar e remover: O(1) em média. O tamanho médio da lista de cada caixa é n / 1000003, cerca de 0,5 no pico do estresse.
2) Para remover um recolhido também é preciso tirar a chave da AVL correspondente: O(log L).

**Alternativa descartada: árvore rubro-negra ou AVL por id**

 
## Evitando duas ambulâncias no mesmo chamado
**Estrutura: vetor de ids já atribuídos no turno**

 
**Custo por turno:**
1) Consultar: O(A), com A = ambulâncias que já escolheram alvo neste turno (no máximo 100 no estresse, num vetor pequeno que cabe no cache). Reservar: O(1).
2) No total, a consulta é feita uma vez por chamado candidato examinado, então o custo por turno fica em O(A) vezes o número de candidatos examinados.

**Alternativa descartada: marcar o chamado no registro da hash (campo de dono)**

 
## Achar o hospital mais próximo
**Estrutura: vetor de hospitais**

**Custo por turno:**
1) O(H) por ambulância ocupada, com H <= 20.
2) O(H) para cada um dos dois finalistas de cada ambulância livre: no máximo 2 · A · H, cerca de 4 mil operações no estresse.

**Alternativa descartada: guardar a distância ao hospital em cada chamado quando ele chega**

 
## Escolhendo entre chamados comuns e graves
**Estrutura: critério de taxa de retorno**

 
**Alternativas descartadas:**
1) Sempre o mais próximo (primeira versão do bot): ignora o valor e deixa os graves passarem para trás de comuns de poucos pontos.
2) Sempre um grave quando existe: a ambulância atravessaria o mapa todo por um grave, enquanto comuns valiosos e próximos seriam ignorados.