package main

import (
    "log"
    "os"
    "os/signal"
    "syscall"
	. "Lab_4_Distribuidos/mercenarios"
)

func main() {
    go RunUserMercenario()
    go RunBotMercenarios()

    // Espera a que se presione Ctrl+C para finalizar el programa
    c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
    <-c

    log.Println("Cerrando aplicación...")
}
