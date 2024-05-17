package mercenarios

import (
    pb "github.com/Zurickata/Lab_4_Distribuidos/proto"
    "google.golang.org/grpc"
    "context"
    "log"
    "fmt"
)

func RunUserMercenario() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewDirectorClient(conn)

    mercenario := &pb.Mercenario{Nombre: "Usuario", Id: 1}
    r, err := c.Preparacion(context.Background(), mercenario)
    if err != nil {
        log.Fatalf("could not prepare: %v", err)
    }
    fmt.Printf("Preparación: %s\n", r.Mensaje)

    // Ejemplo de decisiones por piso
    piso := &pb.Piso{Numero: 1, Decision: "Moverse", Mercenario: mercenario}
    r2, err := c.DecisionesPiso(context.Background(), piso)
    if err != nil {
        log.Fatalf("could not make decision: %v", err)
    }
    fmt.Printf("Decisión piso: %s\n", r2.Mensaje)

    // Consultar el monto en el doshbank
    m, err := c.ConsultarDoshBank(context.Background(), &pb.Vacio{})
    if err != nil {
        log.Fatalf("could not consult doshbank: %v", err)
    }
    fmt.Printf("Monto en doshbank: %d\n", m.Monto)
}
