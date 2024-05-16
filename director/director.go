package main

import (
    pb "github.com/Zurickata/Lab_4_Distribuidos/proto"
    "google.golang.org/grpc"
    "context"
    "fmt"
    "net"
    "strconv"
    "time"
	"sync"
)

type server struct {
    pb.UnimplementedDirectorServiceServer
}

func (s *server) StartMission(ctx context.Context, req *pb.StartMissionRequest) (*pb.StartMissionResponse, error) {
    // Implementar lógica de inicio de misión
    fmt.Printf("Starting mission %s for mercenaries: %v\n", req.MissionId, req.MercenaryIds)
    
	return &pb.StartMissionResponse{}, nil
}

func (s *server) UpdateOrder(ctx context.Context, req *pb.OrderUpdateRequest) (*pb.OrderUpdateResponse, error) {
    fmt.Printf("Updating order for mercenary %d: %s\n", req.MercenaryId, req.NewOrder)
    // Implementar lógica para actualizar órdenes
    return &pb.OrderUpdateResponse{Success: true, Message: "Order updated"}, nil
}

func main() {
	conn, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("No se pudo crear la conexion TCP: " + err.Error())
		return
	}
    serv := grpc.NewServer()
    pb.RegisterDirectorServiceServer(serv, &server{})

    fmt.Println("Director server is running on port 50051...")
    if err = serv.Serve(conn); err != nil{
        fmt.Println("No se pudo levantar el servidor: " + err.Error())
        return
    } 
}
