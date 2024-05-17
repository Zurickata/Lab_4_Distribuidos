package main

import (
    pb "github.com/Zurickata/Lab_4_Distribuidos/proto"
    "google.golang.org/grpc"
    "context"
    "log"
    "net"
    "sync"
    "fmt"
)

type server struct {
    pb.UnimplementedDirectorServer
    mu sync.Mutex
    mercenarios map[int32]*pb.Mercenario
    monto int64
}

func (s *server) Preparacion(ctx context.Context, m *pb.Mercenario) (*pb.Respuesta, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.mercenarios[m.Id] = m
    return &pb.Respuesta{Mensaje: "Preparación confirmada", Exito: true}, nil
}

func (s *server) DecisionesPiso(ctx context.Context, p *pb.Piso) (*pb.Respuesta, error) {
    // Lógica para manejar decisiones en cada piso
    return &pb.Respuesta{Mensaje: "Decisión recibida", Exito: true}, nil
}

func (s *server) ConsultarDoshBank(ctx context.Context, v *pb.Vacio) (*pb.Monto, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    return &pb.Monto{Monto: s.monto}, nil
}

func (s *server) IniciarPiso(ctx context.Context, v *pb.Vacio) (*pb.Iniciar, error) {
    return &pb.Iniciar{Mensaje: "Inicio de piso"}, nil
}

func (s *server) InformarMuerte(ctx context.Context, m *pb.Mercenario) (*pb.Respuesta, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    delete(s.mercenarios, m.Id)
    s.monto += 100000000
    return &pb.Respuesta{Mensaje: "Muerte reportada", Exito: true}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    
    fmt.Println("Servidor en ejecución en el puerto 50051...")
    s := grpc.NewServer()
    pb.RegisterDirectorServer(s, &server{mercenarios: make(map[int32]*pb.Mercenario)})
    log.Printf("Server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
