# Desafio de System Design: Instagram

O objetivo deste desafio é projetar um sistema que atenda aos seguintes requisitos de engenharia, detalhando o que será abordado e o que está fora do escopo.

## Requisitos de Engenharia

## Requisitos Funcionais

- O usuário poderá criar novos posts com possibilidade de fazer upload de imagens e vídeos
- O usuário poderá seguir outros usuários e também ser seguido
- O usuário terá acesso ao feed com posts dos usuários que ele está seguindo.

Não será abordado:

- Stories
- Comentários
- Lives
- Busca
- Compartilhamento de posts para outros usuários
- Listagem de usuários que ele está seguindo

## Requisitos Não Funcionais

- 100 milhões de usuários por dia (DAU)
- Disponibilidade > Consistência
- Latência de < 400ms para exibição do feed
- Baixa latência para exibição de imagens e vídeos

O sistema deve operar de forma assíncrona, sempre que possível, para garantir resiliência e otimização de recursos.

## Não será abordado:

- Implementação de observabilidade
- Processo de entrega de software
- Detalhes de infraestrutura (como uso de Kubernetes)
- Procedimentos de recuperação em caso de falhas

## Tarefas para Completar o System Design (Ferramenta recomendada: Excalidraw)

- Identificar as principais entidades do sistema e suas relações.
- Estimar a capacidades básicas de volume de requisições e storage
- Definir o design da API de alto nível com os principais endpoints e métodos.
- Criar o diagrama visual com as interações dos componentes do sistema.
- Considerar perguntas complexas que possam abordar:
  - Banco de dados
  - Caching
  - Fan out strategies
  - Necessidade de mensageria
  - Processos de upload
  - Escalabilidade horizontal
