package main

import (
	"chess-user/cmd/rpc/internal/config"
	"chess-user/cmd/rpc/internal/service"
	"chess-user/cmd/rpc/proto/user"
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
)

func main() {
	var (
		usrCfg = config.GetUserCfg()
		srv    = grpc.NewServer(
			grpc.StreamInterceptor(streamInterceptor),
			grpc.UnaryInterceptor(unaryInterceptor),
		)

		lis, err = net.Listen("tcp", usrCfg.UserRpc.Hosts+":"+usrCfg.UserRpc.Port)
	)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)

		return
	}
	user.RegisterUserServiceServer(srv, &service.UserServer{})

	log.Println("user-rpc listen to " + usrCfg.UserRpc.Port)
	// 开始处理
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func streamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if err := authorize(stream.Context()); err != nil {
		return err
	}

	return handler(srv, stream)
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if err := authorize(ctx); err != nil {
		return "", err
	}

	return handler(ctx, req)
}

func authorize(ctx context.Context) error {

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if len(md["username"]) > 0 && md["username"][0] == config.GetUserCfg().Name &&
			len(md["password"]) > 0 && md["password"][0] == config.GetUserCfg().UserRpc.Key {
			return nil
		}

		return errors.New("AccessDeniedErr")

	}

	return errors.New("EmptyMetadataErr")
}
