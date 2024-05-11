
# API com Go

Esse é um projeto de uma API desenvolvida com Go utilizando as seguintes ferramentas:
- <a href="https://gin-gonic.com">***Go-Gin***</a>
- <a href="https://gorm.io/index.html">***Go-Gorm***</a>
- <a href="https://github.com/swaggo/swag">***Swaggo***</a>

Essa aplicação conta com uma API REST com os métodos GET, POST, PUT e DELETE, conexão com banco de dados, documentação e para o ambiente de desenvolvimento e prod foi usado o Docker.

![logo go](https://go.dev/images/gophers/ladder.svg)

## Como fazer deploy da aplicação

Primeiro de tudo é preciso fazer um clone do projeto com o comando abaixo:
```github
  git clone https://github.com/Alym62/rest-go.git 
```

Logo após basta entrar no diretório do projeto e realizar o seguinte comando para realizar o deploy. Esse comando vai subir a imagem da aplicação juntamente ao banco de dados que ele precisa se conectar:
```docker
  docker-compose up
```

#### Ambiente de desenvolvimento

Para rodar a aplicação em ambiente de desenvolvimento, basta entrar dentro do diretório ***📂 ./db/*** lá você encontra a configuração e a imagem do banco de dados utilizado na aplicação. depois disso basta voltar para o ***root*** e rodar o comando abaixo:
```go
  go run main.go
```

## Stack utilizada

[![My Skills](https://skillicons.dev/icons?i=go,postgres,docker&perline=3)](https://skillicons.dev)


