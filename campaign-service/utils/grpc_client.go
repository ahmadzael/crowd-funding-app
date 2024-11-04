package utils

import (
    "context"
    "log"
    "myapp/proto"
    "time"

    "google.golang.org/grpc"
)

func ValidateTokenWithAuthService(token string) (bool, string) {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("Failed to connect to gRPC server: %v", err)
    }
    defer conn.Close()

    client := proto.NewAuthServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    res, err := client.ValidateToken(ctx, &proto.TokenRequest{Token: token})
    if err != nil {
        log.Fatalf("Failed to call ValidateToken: %v", err)
    }

    return res.Valid, res.Message
}
