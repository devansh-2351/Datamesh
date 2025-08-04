package data_availability

import (
	"context"
	"net"
	"log"
	datameshpb "github.com/datamesh-labs/datamesh/proto"
	"google.golang.org/grpc"
)

type DAServer struct {
	datameshpb.UnimplementedDAServiceServer
	engine *ErasureKZGEngine
}

func (s *DAServer) PostBlob(ctx context.Context, req *datameshpb.PostBlobRequest) (*datameshpb.PostBlobResponse, error) {
	id, _ := s.engine.PostBlob(ctx, req.Data)
	return &datameshpb.PostBlobResponse{Id: id}, nil
}

func (s *DAServer) SampleBlob(ctx context.Context, req *datameshpb.SampleBlobRequest) (*datameshpb.SampleBlobResponse, error) {
	data, _ := s.engine.SampleBlob(ctx, req.Id)
	return &datameshpb.SampleBlobResponse{Data: data}, nil
}

func RunDAGRPCServer(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	datameshpb.RegisterDAServiceServer(grpcServer, &DAServer{engine: &ErasureKZGEngine{}})
	log.Printf("DA gRPC server listening on %s", addr)
	return grpcServer.Serve(lis)
} 