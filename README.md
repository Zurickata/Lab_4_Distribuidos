# Lab_4_Distribuidos

Integrantes:
    Felipe Guicharrousse
    Rodrigo Ramirez
    Ignacio Riquelme

Para la ejecucion del programa en la raiz del proyecto ingresar por consola
    docker-compose up --build

Luego ejecutar los mercenarios de la siguiente forma
    cd mercenarios
    go run main.go

En otra consola ejecute:
    cd namenode
    go run namenode.go

Y en otras consolas ejecute los datanode ejemplo:
    cd datanodeX
    go run datanode.go
Siendo X los numeros [1,2,3].

Y listo! Juegue!