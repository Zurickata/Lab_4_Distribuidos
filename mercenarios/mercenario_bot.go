package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	pb "github.com/Zurickata/Lab_4_Distribuidos/proto"
	"google.golang.org/grpc"
)

func reportStatus(client pb.MercenaryServiceClient, mercenaryID int32) {
	status := fmt.Sprintf("Status report from mercenary %d", mercenaryID)
	res, err := client.ReportStatus(context.Background(), &pb.StatusReportRequest{
		MercenaryId: mercenaryID,
		Status:      status,
	})
	if err != nil {
		log.Printf("Error reporting status: %v", err)
		return
	}
	log.Printf("Reported status: %s", res.Message)
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMercenaryServiceClient(conn)

	mercenaryID := int32(rand.Intn(100))
	for {
		reportStatus(client, mercenaryID)
		time.Sleep(5 * time.Second)
	}
}
