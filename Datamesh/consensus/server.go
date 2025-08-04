package consensus

import (
	"context"
	"net"
	"log"
	datameshpb "github.com/datamesh-labs/datamesh/proto"
	"google.golang.org/grpc"
)

type ConsensusServer struct {
	datameshpb.UnimplementedConsensusServiceServer
	engine *TendermintEngine
}

func (s *ConsensusServer) GetStatus(ctx context.Context, req *datameshpb.StatusRequest) (*datameshpb.StatusResponse, error) {
	s.engine.IncrementBlock() // For demo, increment block height on each call
	return &datameshpb.StatusResponse{Status: s.engine.Status()}, nil
}

func (s *ConsensusServer) AddValidator(ctx context.Context, req *datameshpb.AddValidatorRequest) (*datameshpb.AddValidatorResponse, error) {
	success := s.engine.AddValidator(req.Validator)
	return &datameshpb.AddValidatorResponse{Success: success}, nil
}

func (s *ConsensusServer) RemoveValidator(ctx context.Context, req *datameshpb.RemoveValidatorRequest) (*datameshpb.RemoveValidatorResponse, error) {
	success := s.engine.RemoveValidator(req.Id)
	return &datameshpb.RemoveValidatorResponse{Success: success}, nil
}

func (s *ConsensusServer) ListValidators(ctx context.Context, req *datameshpb.ListValidatorsRequest) (*datameshpb.ListValidatorsResponse, error) {
	vals := s.engine.ListValidators()
	return &datameshpb.ListValidatorsResponse{Validators: vals}, nil
}

func RunConsensusGRPCServer(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	datameshpb.RegisterConsensusServiceServer(grpcServer, &ConsensusServer{engine: &TendermintEngine{}})
	log.Printf("Consensus gRPC server listening on %s", addr)
	return grpcServer.Serve(lis)
} 