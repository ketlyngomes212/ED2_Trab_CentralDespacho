# Relatório do projeto
Nome: Ketlyn Gomes de Oliveira, Luigi Favaro
GRR: 20252054, 20254073

## Achar os chamados próximos de uma ambulância
**Estruturas: uma AVL por linha do mapa (chamados comuns) + lista (chamados graves)**
Os chamados comuns ficam em árvores AVL, sendo uma por linha. A chave de cada chamado junta a coluna que ele pertence e o id, o que permite vários chamados por célula e mantém a ordem por coluna. Uma mbulância busca um chamado analisando linhas vizinhas, através de losangos com raios 16, 32, 64, etc. Em cada linha, pega apenas a faixa de colunas do raio. A busca termina quando nenhum chamado mais longe consegue ganhar do melhor encontrado. Os graves não ficam na AVL, mas em um vetor, pois são poucos em relação aos comuns e vale a pena analisar todos eles.

**Custo por turno:**
1) Para chamados comuns: inserir e remover O(log L), pois a árvore é só da linha específica.
2) Busca de uma ambulância: O(R * log L + c), com R linhas percorridas e c chaves percorridas nas faixas. Os orçamentos (600 avaliações depois de achar um candidato, 3000 no máximo) são conferidos ao fim de cada losango, então limitam quantos são percorridos, mas não o tamanho de cada um.
3) Para chamados graves: inserir e remover O(1); avaliar custa O(G) por ambulância livre, ou O(A * G) por turno

**Alternativa descartada: Grade por blocos, Hash**
Na grade, embora inserir e remover sejam O(1) e a consulta leia os blocos em volta, analisamos que 30% dos chamados se concentram em poucos pontos: os blocos desses pontos ficam enormes e a consulta nesses lugares vira quase uma varredura, e é preciso escolher o tamanho do bloco para um mapa e uma densidade. Enquanto isso, a AVL por linha não depende da distribuição: o custo é logarítmico em qualquer concentração. Uma hash por célula tem a mesma ideia, por não ter ordem, para fazer a varredura ela precisa verificar célula por célula, o que custa caro.

## Descartar os chamados expirados
**Estrutura: heap mínimo por prazo**
Cada chamado entra na heap com (prazo, id). A cada turno, enquanto o topo tiver prazo < turno, ele sai e o chamado é removido das demais estruturas. O prazo é o último turno válido, então prazo == turno ainda vale. Chamados recolhidos continuam no heap até vencer. Ao saírem, a tabela hash não os acha mais e nada acontece, evitando uma busca na heap.

**Custo por turno:**
1) Inserir: O(log n) por chamado novo; olhar o topo: O(1).
2) Expirar: O(k log n), com k = chamados vencidos naquele turno. Os que não venceram não são tocados.

**Alternativa descartada: percorrer todos os chamados a cada turno**
Custaria O(n) por turno. No estresse n chega a 500 mil, e isso daria milhões de operações ao final, apenas para descobrir os que venceram. O heap só olha quem venceu. Uma alternativa melhor seria uma fila indexada pelo turno de vencimento, já que os prazos são inteiros limitados: ambas as operações ficariam com custo O(1). O heap foi escolhido porque não depende de saber o prazo máximo e seu custo já é o suficiente para não exceder o limite.

 ## Achar / Remover um chamado recolhido sabendo só o identificador
 **Estrutura: tabela hash com encadeamento**
O servidor informa só o id, e ele não é sequencial, portanto, não serve como índice para vetor. A hash usa 100003 caixas (número primo), função `id % 100003` e encadeamento por índices num vetor com nós pré-alocado, com uma lista que possui os nós livres (sem `map` e sem alocar durante a partida). O registro guarda todos os dados do chamado, e as outras estruturas guardam apenas a chave.

**Custo por turno:**
1) Inserir, buscar e remover: O(1) em média; o tamanho médio da fila de cada caixa
  é n / 100003, cerca de 5 no pico em um cenário de estresse.
2) Para remover um recolhido também precisamos remover o chamado da AVL ou da lista de
  graves: O(log L) ou O(1).

**Alternativa descartada: rubro-negra, AVL por id**
Ambas garantem O(log n) no pior caso e mantém os ids ordenados. O problema surge pois ficam mais lentas na média (cerca de 18 comparações para 500 mil chamados, contra cerca de 5 da hash.), além de que a ordem por id não é útil aqui. A árvore seria melhor se os ids tivessem, por exemplo, mesmo resto na divisão, o que faria a hash ter custo O(n) ou se precisássemos percorrer os chamados em ordem de id.

## Evitando duas ambulâncias no mesmo chamado
**Estrutura: campo Dona no registro do chamado**
Ao escolher um alvo, a ambulância grava seu id no chamado, no campo Dona. Na avaliação os chamados já reservados são ignorados. O alvo só será liberado caso haja troca ou perda pela ambulância.

**Custo por turno:**
1) Consultar e reservar: ambas com O(1).

**Alternativa descartada: lista com os chamados já atribuídos no turno**
Na primeira versão do bot, cada consulta custava O(A). Com 100 ambulâncias avaliando muitos candidatos, o custo era multiplicado por 100. A lista seria interessante apenas em cenários pequenos, com poucas ambulâncias e poucos chamados.

## Analisando o hospital mais próximo
**Estrutura: vetor de hospitais**
A distância de um chamado ao hospital mais próximo é calculada somente uma vez, quando o chamado chega. Esse valor fica guardado em DistHospital. Para ambulâncias ocupadas, ocorre uma varredura.

**Custo por turno:**
1) O(H) com H <= 20 para cada chamado que surge e cada ambulância que fica ocupada.

## Escolhendo entre chamados comuns e graves
**Estrutura: critério de taxa de retorno**
A taxa de retorno é um valor utilizado na escolha do alvo, através de constantes (definidas após testes) e das distâncias. A distância até o chamado tem um peso menor no cálculo, pois na ida a outra equipe pode roubar o chamado, enquanto na volta isso é impossível. O cálculo é o seguinte: taxa = valor do chamado / (distância até o chamado + 0,75 * distância chamado ao hospital + 1)

**Alternativas descartadas:**
1) Mais próximos (a primeira versão do bot): ignora completamente o valor, ignorando graves por comuns de poucos pontos.

2) Priorizar apenas graves: a ambulância percorre o mapa todo se necessário, apenas para coletar um grave, enquanto alguns comuns valiosos mais próximos são ignorados.