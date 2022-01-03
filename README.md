# API REST utilizando Docker, Go, e MongoDB

## As rotas implementadas foram:

- GET /movies/ 
Listagem de todos os filmes
- GET /movies/{id} 
Recuperar algum filme utilizando o ID como parâmetro
- POST /movies/ 
Realiza a inserção de um novo filme
- DELETE /movies/{id} 
Realiza a remoção de determinado filme com determinado ID

## Dependências

- [Docker Compose](https://docs.docker.com/compose/) 1.25

Dentro do diretorio do projeto execute:
'''

docker-compose build
'''

'''

docker-compose up
'''

