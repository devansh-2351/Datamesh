package execution

import (
	"context"
	"net"
	"log"
	datameshpb "github.com/datamesh-labs/datamesh/proto"
	"google.golang.org/grpc"
)

type ExecutionServer struct {
	datameshpb.UnimplementedExecutionServiceServer
}

func (s *ExecutionServer) ExecuteTx(ctx context.Context, req *datameshpb.ExecuteTxRequest) (*datameshpb.ExecuteTxResponse, error) {
	return &datameshpb.ExecuteTxResponse{Result: []byte("execution-result-stub")}, nil
}

func RunExecutionGRPCServer(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	datameshpb.RegisterExecutionServiceServer(grpcServer, &ExecutionServer{})
	log.Printf("Execution gRPC server listening on %s", addr)
	return grpcServer.Serve(lis)
} 