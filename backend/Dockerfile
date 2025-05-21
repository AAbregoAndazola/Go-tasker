# backend/Dockerfile

FROM golang:1.21-alpine

WORKDIR /app

# Copiamos los archivos de dependencias
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copiamos el resto del código
COPY . .

# Compilamos la app (el main está en /cmd/main.go)
RUN go build -o main ./cmd/main.go

EXPOSE 3000

CMD ["./main"]
