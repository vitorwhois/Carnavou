# Meu Projeto Carnaval

O projeto consistem uma api construida em Golang, utilizando o banco de dados Supabase.
Este projeto tem como objetivo facilitar a visualização dos blocos de carnaval em São Paulo.
Os blocos serão organizados para o usuario buscar por nome, data e categorias (tamanho, Estilo e caracteristica).
Sera disponibilizado a possibilidade de salvar uma lista dos blocos escolhidos e compartilhar o link com amigos.

## Conteúdo
![Pagina inicial](https://i.imgur.com/TLWmoh2.jpg)
- Página inicial:
    - Barra de pesquisa de blocos por nome.
    - Botão Filtrar para encontrar blocos por região e data.
    - Seção agenda:
        - Carrega os blocos do dia atual, limitado ao último dia de carnaval.
    - Mapa:
        - Mapa com todos os blocos de rua de São Paulo e localização.


![Pagina de resultado a busca](https://i.imgur.com/2Df9esh.jpeg)  
- Resultado da sua busca:
    - O resultado da busca aparece na seção principal da Home, com mensagem de erro caso não encontre.
- Página minha lista:
    - O usuário pode salvar os blocos ao clicar no botão "salvar +" do card.
    - Os blocos são salvos no localstorage.
    - A página minha lista carrega todos os blocos salvos.
    - Botão "compartilhar" para receber o link da lista criada pelo usuário e poder compartilhar com os amigos.
    - Botão "Remover -" para excluir um bloco da lista.
    - Mensagem de erro direcionando para a página de busca, quando a minha lista é carregada vazia, sem blocos salvos.
![Pagina sobre](https://i.imgur.com/0duCjdU.jpg)  
- Página Sobre:
      - Informações sobre a equipe.
  
## Autores
Vitor Ruis da Silva - Desenvolvedor e Mariana do Brasil - Designer.

## Agradecimentos
Agradeça qualquer pessoa ou recurso que tenha contribuído para o seu projeto.
(...)

## Status do Projeto
MVP lançado na semana o carnaval de 2024 com mais de 2mil visualizações.
Novas features e melhorias serão implementadas no segundo semestre.

