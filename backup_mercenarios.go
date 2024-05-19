package main

import (
    pb "github.com/Zurickata/Lab_4_Distribuidos/proto"
    "google.golang.org/grpc"
    "context"
    "log"
    "math/rand"
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "bufio"
    "strconv"
    "sync"
    "strings"
)

var mu sync.Mutex

func main() {
    go RunUserMercenario()
    go RunBotMercenarios()

    // Espera a que se presione Ctrl+C para finalizar el programa
    c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
    <-c

    log.Println("Cerrando aplicación...")
}

func RunBotMercenarios() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewDirectorClient(conn)

    for i := 2; i <= 8; i++ {
        go func(id int) {
            mercenario := &pb.Mercenario{Nombre: fmt.Sprintf("Bot-%d", id), Id: int32(id)}
            for pisoNum := 1; pisoNum <= 3; pisoNum++ {
                mu.Lock()
                // Mercenario indica que está preparado
                r, err := c.Preparacion(context.Background(), mercenario)
                if err != nil {
                    log.Fatalf("could not prepare: %v", err)
                }
                log.Printf("Preparación Bot-%d: %s\n", id, r.Mensaje)

                // Espera la señal para iniciar el siguiente piso
                r3, err := c.InicioPiso(context.Background(), &pb.Piso{Numero: int32(pisoNum)})
                if err != nil {
                    log.Fatalf("could not start floor: %v", err)
                }
                log.Printf("Bot-%d: %s\n", id, r3.Mensaje)

                // Mercenario indica decision del piso según piso
                decision := int32(1)
                switch pisoNum {
                case 1:
                    decision = int32(randomInRange(1, 3))
                case 2:
                    decision = int32(randomInRange(1, 2))
                case 3:
                    decision = int32(randomInRange(1, 15))
                }
                piso := &pb.Piso{
                    Numero:    int32(pisoNum),
                    Decision:  decision,
                    Mercenario: mercenario,
                }

                r2, err := c.DecisionesPiso(context.Background(), piso)
                if err != nil {
                    log.Fatalf("could not make decision: %v", err)
                }
                log.Printf("Decisión (%d) piso %d Bot-%d: %s\n", piso.Decision, pisoNum, id, r2.Mensaje)

                // Muerte del Mercenario
                if r2.Mensaje == "Murió" {
                    log.Printf("Bot-%d murió en el piso %d\n", id, pisoNum)
                    mu.Unlock()
                    break
                }

                // Indicar que está listo para el siguiente piso
                r4, err := c.ListoParaSiguientePiso(context.Background(), mercenario)
                if err != nil {
                    log.Fatalf("could not signal readiness: %v", err)
                }
                log.Printf("Bot-%d listo para el siguiente piso: %s\n", id, r4.Mensaje)
                mu.Unlock()
            }
        }(i)
    }

    select {}
}

func RunUserMercenario() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewDirectorClient(conn)

    mercenario := &pb.Mercenario{Nombre: "Usuario", Id: 1}
    for pisoNum := 1; pisoNum <= 3; pisoNum++ {
        mu.Lock()
        // Mercenario indica que está preparado
        r, err := c.Preparacion(context.Background(), mercenario)
        if err != nil {
            log.Fatalf("could not prepare: %v", err)
        }
        log.Printf("Preparación Usuario: %s\n", r.Mensaje)

        // Espera la señal para iniciar el siguiente piso
        r3, err := c.InicioPiso(context.Background(), &pb.Piso{Numero: int32(pisoNum)})
        if err != nil {
            log.Fatalf("could not start floor: %v", err)
        }
        log.Printf("Usuario: %s\n", r3.Mensaje)

        // Mercenario indica decision del piso según piso
        decision := getUserDecision(pisoNum)
        piso := &pb.Piso{
            Numero:    int32(pisoNum),
            Decision:  decision,
            Mercenario: mercenario,
        }

        r2, err := c.DecisionesPiso(context.Background(), piso)
        if err != nil {
            log.Fatalf("could not make decision: %v", err)
        }
        log.Printf("Decisión (%d) piso %d Usuario: %s\n", piso.Decision, pisoNum, r2.Mensaje)

        // Muerte del Mercenario
        if r2.Mensaje == "Murió" {
            log.Printf("Usuario murió en el piso %d\n", pisoNum)
            mu.Unlock()
            break
        }

        // Indicar que está listo para el siguiente piso
        r4, err := c.ListoParaSiguientePiso(context.Background(), mercenario)
        if err != nil {
            log.Fatalf("could not signal readiness: %v", err)
        }
        log.Printf("Usuario listo para el siguiente piso: %s\n", r4.Mensaje)
        mu.Unlock()
    }
}

func getUserDecision(pisoNum int) int32 {
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("Ingrese la decisión para el piso %d: ", pisoNum)
    decisionStr, _ := reader.ReadString('\n')
    decisionStr = strings.TrimSpace(decisionStr) // Elimina espacios en blanco incluyendo \r y \n
    decision, err := strconv.Atoi(decisionStr)
    if err != nil {
        log.Fatalf("Invalid decision: %v", err)
    }
    return int32(decision)
}

// Función para generar un número aleatorio dentro de un rango
func randomInRange(min, max int) int {
	return min + rand.Intn(max-min+1)
}
