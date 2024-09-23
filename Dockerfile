# Etapa de compilação
FROM golang:1.16-alpine AS builder
WORKDIR /app

# Copiar apenas os arquivos necessários para download de dependências
# Isso aproveita melhor o cache do Docker em builds subsequentes
COPY jwt-app/go.mod jwt-app/go.sum ./
RUN go mod download

# Copiar os arquivos de código fonte e compilar o aplicativo
COPY jwt-app/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o jwt-app .

# Etapa de imagem final
FROM scratch

# Copia o binário do builder para a imagem scratch
COPY --from=builder /app/jwt-app /jwt-app

# Define o comando para executar o aplicativo
CMD ["/jwt-app"]
