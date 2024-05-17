package main

import (
	pb "github.com/Zurickata/Lab_4_Distribuidos/proto"
	"google.golang.org/grpc"

	"context"
	"fmt"
	"log"
	"net"
	"os"
)

type Server struct {
	pb.UnimplementedDataNodeServiceServer
}

func (s *Server) StoreDecision(ctx context.Context, req *pb.DecisionRequest) (*pb.DecisionResponse, error) {
	filename := fmt.Sprintf("%s_%d.txt", req.Mercenario, req.Piso)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Could not open file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(req.Decision + "\n")
	if err != nil {
		return nil, fmt.Errorf("Could not write to file: %v", err)
	}

	return &pb.DecisionResponse{Success: true}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("Failed to listen on port 50054: %v", err)
		return
	}

	s := &Server{}
	serv := grpc.NewServer()
	pb.RegisterDataNodeServiceServer(serv, s)

	fmt.Println("DataNode server is running on port 50054...")
	if err := serv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
		return
	}
}
