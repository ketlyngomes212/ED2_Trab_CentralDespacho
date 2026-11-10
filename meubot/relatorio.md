# Relatório do projeto
Nome: Ketlyn Gomes de Oliveira, Luigi Favaro
GRR: 20252054, 20254073

## Achar os chamados próximos de uma ambulância
**Estruturas: duas AVL por linha do mapa (uma para chamados comuns, outra só para os graves), com os dados na tabela hash**
Cada linha y do mapa tem uma árvore AVL. A chave de cada chamado junta a coluna e o id num só inteiro., garantindo valores diferentes para chaves na mesma coluna
 
A busca expande em faixas horizontais a partir da posição da ambulância, pulando chamados já atribuídos ou com prazos "impossíveis".

**Custo por turno:**
1) Inserir e remover: O(log L), com L o número de chamados da linha.
2) Busca de uma ambulância: O((R + c) · log L), com R = linhas percorridas (no máximo a distância até o candidato encontrado) e c = chaves puladas por não servirem. Cada chave pulada custa uma nova consulta na árvore, O(log L), mais um acesso O(1) à hash. Linhas vazias custam O(1).
3) Por turno: uma busca nos comuns e uma nos graves para cada ambulância livre.

**Alternativa descartada: grade por blocos (hash espacial) E heap máximo por valor para os graves**
A grade por blocos dividiria o mapa em células e guardaria listas com chamados em cada uma das células. A busca visitaria as células dentro do raio de interesse. Embora essa implementação seja simples, notamos que o custo é proporcional aos chamados por célula (inviável). 

O heap máximo por valor para graves permitiria pegar sempre o mais valioso, mas sem informação espacial: a ambulância poderia atravessar o mapa inteiro por um grave que vale um pouco mais, perdendo vários comuns que valessem a pena no caminho.

## Descartar os chamados expirados
**Estrutura: heap mínimo por prazo**
Todos os chamados abertos entram em uma fila ordenada pelo prazo. A cada turno, a função expirar retira do topo enquanto o prazo do topo for menor que o turno atual, removendo cada chamado expirado.

**Custo por turno:**
1) Inserir: O(log n) por chamado novo; olhar o topo: O(1).
2) Expirar: O(k log n), com k = chamados vencidos naquele turno. Os que não venceram não são tocados.

**Alternativa descartada: percorrer todos os chamados a cada turno**
Percorrer todos os chamados abertos custaria O(n) por turno, o que ficaria inviável em cenários de estresse.

## Achar e remover um chamado recolhido sabendo só o identificador
**Estrutura: tabela hash com encadeamento e pool de nós**
A tabela hash é composta por nós, cada um guarda o chamado e o índice do próximo nó da mesma caixa. 
 
**Custo por turno:**
1) Inserir, buscar e remover: O(1) em média. O tamanho médio da lista de cada caixa é n / 1000003, cerca de 0,5 no pico do estresse.
2) Para remover um recolhido também é preciso tirar a chave da AVL correspondente: O(log L).

**Alternativa descartada: árvore rubro-negra ou AVL por id**
Uma árvore por id daria O(log n) para busca e remoção por id, pior que o custo O(1) da hash. 
 
## Evitando duas ambulâncias no mesmo chamado
**Estrutura: vetor de ids já atribuídos no turno**
A cada turno, guardamos um slice atribuidos com os chamados escolhidos por ambulâncias. Antes de confirmar um chamado como destino, chamadoOcupado percorre esse slice, para garantir que não haja desencontro.
 
**Custo por turno:**
1) Consultar: O(A), com A = ambulâncias que já escolheram alvo neste turno (no máximo 100 no estresse, num vetor pequeno que cabe no cache). Reservar: O(1).
2) No total, a consulta é feita uma vez por chamado candidato examinado, então o custo por turno fica em O(A) vezes o número de candidatos examinados.

**Alternativa descartada: marcar o chamado no registro da hash (campo de dono)**
Na versão inicial existia um campo Dona em cada chamado, que dispensaria o vetor de atribuídos e reduziria a consulta a O(1). No entanto, essa abordagem exigiria limpar os campos Dona a cada turno (O(A)), complicando na remoção.
 
## Achar o hospital mais próximo
**Estrutura: vetor de hospitais**
Como há no máximo 20 hospitais, guardamos em um vetor normal e escolherHospital percorre  para devolver o de menor distância até a ambulância em questão.

**Custo por turno:**
1) O(H) por ambulância ocupada, com H <= 20.
2) O(H) para cada um dos dois finalistas de cada ambulância livre: no máximo 2 · A · H, cerca de 4 mil operações no estresse.

**Alternativa descartada: guardar a distância ao hospital em cada chamado quando ele chega**
Embora calcular previamente a distância economize tempo, o ganho real é pequeno. Manter a lista simples de hospitais é mais simples e não viola nenhuma restrição de tempo.
 
## Escolhendo entre chamados comuns e graves
**Estrutura: critério de taxa de retorno**
A função escolherAlvo em graves.go obtém o melhor grave dentro de alcanceGrave e o melhor comum sem limite de distância, depois compara os dois pelo critério de taxa de retorno:

taxa = Valor / (distância_até_chamado + distância_chamado_hospital + 1)

O parâmetro pesoGrave multiplica o valor do grave antes da comparação, ou seja, com pesoGrave = 1, graves e comuns teriam mesma "importância".

 
**Alternativas descartadas:**
1) Sempre o mais próximo (primeira versão do bot): ignora o valor e deixa os graves passarem para trás de comuns de poucos pontos.
2) Sempre um grave quando existe: a ambulância atravessaria o mapa todo por um grave, enquanto comuns valiosos e próximos seriam ignorados.
3) Lista densa com varredura total (segunda versão): listaGraves avaliava todos os graves abertos por ambulância livre, sem se importar com distância. Embora fizesse mais pontos, cometia muitos estouros de tempo.