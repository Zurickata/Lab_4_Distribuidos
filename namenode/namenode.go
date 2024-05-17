package main

import (
	pb "github.com/Zurickata/Lab_4_Distribuidos/proto"
	"google.golang.org/grpc"

	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
)

type Server struct {
	pb.UnimplementedNameNodeServiceServer
	dataNodeAddresses []string
}

func (s *Server) RegisterDecision(ctx context.Context, req *pb.DecisionRequest) (*pb.DecisionResponse, error) {
	dataNodeAddress := s.dataNodeAddresses[rand.Intn(len(s.dataNodeAddresses))]
	conn, err := grpc.Dial(dataNodeAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("Could not connect to datanode: %v", err)
	}
	defer conn.Close()

	client := pb.NewDataNodeServiceClient(conn)
	_, err = client.StoreDecision(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Could not store decision: %v", err)
	}

	// Registrar decisión

	return &pb.DecisionResponse{Success: true}, nil
}

func (s *Server) GetDecisions(ctx context.Context, req *pb.GetDecisionsRequest) (*pb.GetDecisionsResponse, error) {
	// Lógica para obtener decisiones
	return &pb.GetDecisionsResponse{Decisions: []string{}}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
		return
	}

	s := &Server{
		dataNodeAddresses: []string{"datanode1:50052", "datanode2:50053", "datanode3:50054"},
	}

	serv := grpc.NewServer()
	pb.RegisterNameNodeServiceServer(serv, s)

	fmt.Println("NameNode server is running on port 50051...")
	if err = serv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
		return
	}
}
