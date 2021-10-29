# Banco-api
Banco-api é uma API para algumas rotinas bancarias, como criação de contas, listagem de contas, obter o saldo para uma conta específica e transferencia entre contas internas de um banco digital.

# Requisitos/dependências
- Dockerfile
- Docker-compose

# Iniciando


# API
Regras gerais
Usar formato JSON para leitura e escrita. (ex: GET /accounts/ retorna json, POST /accounts/ {name: 'james bond'})

# Rotas esperadas
/accounts
A entidade Account possui os seguintes atributos:

id
name
cpf
secret
balance
created_at

# Espera-se as seguintes ações:

GET /accounts - obtém a lista de contas
GET /accounts/{account_id}/balance - obtém o saldo da conta
POST /accounts - cria uma conta
Regras para esta rota

balance pode iniciar com 0 ou algum valor para simplificar
secret deve ser armazenado como hash
/login

# A entidade Login possui os seguintes atributos:

cpf
secret

# Espera-se as seguintes ações:

POST /login - autentica a usuaria
Regras para esta rota

Deve retornar token para ser usado nas rotas autenticadas
/transfers

# A entidade Transfer possui os seguintes atributos:

id
account_origin_id
account_destination_id
amount
created_at

# Espera-se as seguintes ações:

GET /transfers - obtém a lista de transferencias da usuaria autenticada.
POST /transfers - faz transferencia de uma conta para outra.

# Regras para esta rota

Quem fizer a transferência precisa estar autenticada.
O account_origin_id deve ser obtido no Token enviado.
Caso Account de origem não tenha saldo, retornar um código de erro apropriado
Atualizar o balance das contas.

