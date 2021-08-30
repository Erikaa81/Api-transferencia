# Imagem oficial do golang com suporte a go modules
FROM golang

# Acessando o diretório de trabalho
WORKDIR /app/src/Banco-api

# aponta a variavel gopath do go para o diretorio app
ENV GOPATH=/app

# copia os arquivos do projeto para o workdir do container
COPY . /app/src/Banco-api/

# execulta o main.go e baixa as dependencias do projeto
RUN go build main.go


# Habilitando a API
CMD ./main




