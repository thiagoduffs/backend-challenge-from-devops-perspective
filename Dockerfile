# Etapa de compilação
FROM golang:1.16-alpine AS builder
WORKDIR /app

# Copia apenas o diretório jwt-app, que contém o go.mod, go.sum e o código fonte
COPY jwt-app/ .

# Compila o binário como estático
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o jwt-app .

# Etapa de imagem final
FROM scratch

# Copia o binário do builder para a imagem scratch
COPY --from=builder /app/jwt-app /jwt-app

# Define o comando para executar o aplicativo
CMD ["/jwt-app"]
