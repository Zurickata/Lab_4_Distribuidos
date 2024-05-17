package mercenarios

import (
    pb "github.com/Zurickata/Lab_4_Distribuidos/proto"
    "google.golang.org/grpc"
    "context"
    "log"
    "math/rand"
    "time"
    "fmt"
)

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
            r, err := c.Preparacion(context.Background(), mercenario)
            if err != nil {
                log.Fatalf("could not prepare: %v", err)
            }
            log.Printf("Preparación Bot-%d: %s\n", id, r.Mensaje)

            // Ejemplo de decisiones por piso
            for j := 1; j <= 10; j++ {
                piso := &pb.Piso{Numero: int32(j), Decision: "Moverse", Mercenario: mercenario}
                r2, err := c.DecisionesPiso(context.Background(), piso)
                if err != nil {
                    log.Fatalf("could not make decision: %v", err)
                }
                log.Printf("Decisión piso Bot-%d: %s\n", id, r2.Mensaje)
                time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
            }

            // Consultar el monto en el doshbank
            m, err := c.ConsultarDoshBank(context.Background(), &pb.Vacio{})
            if err != nil {
                log.Fatalf("could not consult doshbank: %v", err)
            }
            log.Printf("Monto en doshbank Bot-%d: %d\n", id, m.Monto)
        }(i)
    }

    select {}
}
