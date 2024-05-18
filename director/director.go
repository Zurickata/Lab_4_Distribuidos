package director

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Zurickata/Lab_4_Distribuidos/proto"
	"google.golang.org/grpc"
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

func sendExampleDecision() {
	conn, err := grpc.Dial("namenode:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to namenode: %v", err)
	}
	defer conn.Close()

	client := pb.NewNameNodeServiceClient(conn)

	decisionRequest := &pb.DecisionRequest{
		Mercenario: "Felipe Güicharrousse",
		Piso:       1,
		Decision:   "Hola como estas",
	}

	response, err := client.RegisterDecision(context.Background(), decisionRequest)
	if err != nil {
		log.Fatalf("Could not register decision: %v", err)
	}

	fmt.Printf("Decision registered: %v\n", response.Success)
}

func main() {

	go sendExampleDecision() // Enviar mensaje de prueba al iniciar

	conn, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("No se pudo crear la conexion TCP: " + err.Error())
		return
	}
	serv := grpc.NewServer()
	pb.RegisterDirectorServiceServer(serv, &server{})

	fmt.Println("Director server is running on port 50051...")
	if err = serv.Serve(conn); err != nil {
		fmt.Println("No se pudo levantar el servidor: " + err.Error())
		return
	}
}
