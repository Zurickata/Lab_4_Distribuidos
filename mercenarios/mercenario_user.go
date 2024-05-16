package main

import (
    pb "path/to/proto"
    "context"
    "fmt"
    "log"
    "bufio"
    "os"
    "google.golang.org/grpc"
)

func reportStatus(client pb.MercenaryServiceClient, mercenaryID int32) {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter status report: ")
    status, _ := reader.ReadString('\n')

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

    mercenaryID := int32(1) // ID fijo para usuario

    for {
        reportStatus(client, mercenaryID)
    }
}
