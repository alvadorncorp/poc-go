# Poc go

Esse repositório é para teste de conhecimento em golang, que será feito de forma evolutiva por tarefas!

## TECH-1

Requisitos:
Criar um endpoint /auth/signin que é um endpoint POST

Aceita como parametros um json com o seguinte formato:

```json
{
"user": <user>,
"password": <password>
}
```

O usuario e senha válido, deve ser uma constante do programa, ou seja, não precisa se preocupar com banco de dados nesse primeiro momento, ex:

```go
var authorizedUser = "pedro"

var authorizedPassword = "senha123"
```

Para que o usuario e senha sejam válidos, basta checar se o usuario do json é igual ao usuario da variavel authorizedUser e a senha do json é igual a senha do authorizedPassword.

Se a autenticação for com sucesso, deve gerar um token JWT e retornar um payload da seguinte forma:

```json
{
"token": <token jwt>
}
```

Se a autenticação falhar, retornar o status 401 com o payload da seguinte forma:

```json
{
  "error": "authentication_failure"
}
```

Uso de bibliotecas:

Quero que você utilize essa biblioteca em Go para criar a API: https://github.com/gin-gonic/gin

Documentação em português: https://gin-gonic.com/pt/docs/

Essa é uma biblioteca que permite a definição de endpoints em Go, assim como você deve ter feito utilizando a http.Handler ou Gorilla nas video aulas.

Para geração do JWT, use a seguinte biblioteca:
https://github.com/golang-jwt/jwt

Para saber o que é JWT, da uma olhadinha nesse vídeo: https://www.youtube.com/watch?v=sHyoMWnnLGU

Observação, JWT não depende do Go, ou seja, ele existe independente do Go e quase toda linguagem de programação tem uma biblioteca que gera um token JWT.

## TECH-2

Requisitos:
Criar um arquivo com um banco de dados interno, no proprio projeto.

Dentro desse arquivo criar um map passando a variavel de usuario e senha, com a chave email.

Criar uma função chamada BuscaUsuario, onde ela precisa encontrar oos usuarios informados no map.

No controller, fazer a verificação e a autenticação do usuario e se passou, autenticar a senha, retornando o tocken.

## TECH-3

Criar uma nova rota do tipo POST, para criar um novo usuario.

Criar essa função no db, e passa-la para o controller.

Verificar o email, nao pode criar um novo usuario com mesmo email.

## TECH-4 && TECH-5

salvar senha criptografada

Criar um docker-compose que suba um postgres na versão 17

O seu container postgres deve ter dados persistentes, ou seja, se você fizer um docker-compose down, ele deverá manter os dados anteriores.

No arquivo main.go instanciar uma conexão com o postgres usando a biblioteca: https://github.com/jackc/pgx

Não precisa alterar nada do código que vc já fez relativo a autenticação, essa primeira tarefa é só pra estabelecer uma comunicação entre um programa feito em go e um container postgres.

## TECH-6

Manter o padrao dos nomes da função, com uma so linguagem.

Criar um banco com postgres
