package main

import (
	"io/ioutil"
	"strings"

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
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(req.Decision + "\n")
	if err != nil {
		return nil, fmt.Errorf("could not write to file: %v", err)
	}

	return &pb.DecisionResponse{Success: true}, nil
}

func (s *Server) FetchDecisions(ctx context.Context, req *pb.FetchDecisionsRequest) (*pb.FetchDecisionsResponse, error) {
	var decisions []string

	files, err := ioutil.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("could not read directory: %v", err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".txt") {
			content, err := ioutil.ReadFile(file.Name())
			if err != nil {
				return nil, fmt.Errorf("could not read file %s: %v", file.Name(), err)
			}
			decisions = append(decisions, strings.Split(string(content), "\n")...)
		}
	}

	return &pb.FetchDecisionsResponse{Decisions: decisions}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("Failed to listen on port 50054: %v", err)
	}

	s := &Server{}
	serv := grpc.NewServer()
	pb.RegisterDataNodeServiceServer(serv, s)

	fmt.Println("DataNode server is running on port 50054...")
	if err := serv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
