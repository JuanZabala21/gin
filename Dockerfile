FROM golang:latest

LABEL maintainer="Juan Zabala"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o gin .

# No elimines el ejecutable
# RUN find . -name "*.go" -type f -delete

# Exponer el puerto 5000, ya que parece ser el que usas
EXPOSE 5000

# Ejecuta la aplicación
CMD ["./gin"]
