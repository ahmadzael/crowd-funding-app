package main

import (
    "context"
    "log"
    "net"
    "myapp/proto"
    "myapp/utils"

    "google.golang.org/grpc"
)

type server struct {
    proto.UnimplementedAuthServiceServer
}

func (s *server) ValidateToken(ctx context.Context, req *proto.TokenRequest) (*proto.TokenResponse, error) {
    tokenString := req.Token
    token, err := utils.ValidateJWT(tokenString)

    if err != nil || !token.Valid {
        return &proto.TokenResponse{Valid: false, Message: "Invalid token"}, nil
    }
    return &proto.TokenResponse{Valid: true, Message: "Token is valid"}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    proto.RegisterAuthServiceServer(grpcServer, &server{})

    log.Println("gRPC Auth Service running on port :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
