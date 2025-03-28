package main

import (
	"context"

	"github.com/abhijeet-singh800/grpc-test/grpcPkg"
)

type userServer struct {
	grpcPkg.UnimplementedUserServiceServer
}

func (s *userServer) GetUser(ctx context.Context)

func main() {

}
