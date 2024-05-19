FROM golang:1.22 AS build-stage

WORKDIR /app

# Copiamos los archivos necesarios para construir la aplicación
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY proto/ /app/proto
COPY director/director.go .
COPY datanode1/datanode.go .
COPY datanode2/datanode.go .
COPY datanode3/datanode.go .
COPY mercenarios/mercenario_bot.go .
COPY mercenarios/mercenario_user.go .
COPY namenode/namenode.go .

# Compilamos el binario para Linux
RUN GOOS=linux GOARCH=amd64 go build -o director director.go

# Aseguramos que el binario tenga permisos de ejecución
RUN chmod +x director

# Exponemos el puerto
EXPOSE 50051

# Definimos el comando de inicio
CMD ["./director"]
