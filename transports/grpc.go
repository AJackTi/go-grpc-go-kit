package transports

import (
	"context"

	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"

	"go-grpc-go-kit/endpoints"
	"go-grpc-go-kit/pb"
)

type gRPCServer struct {
	pb.UnimplementedMathServiceServer
	add      grpc.Handler
	subtract grpc.Handler
	multiply grpc.Handler
	divide   grpc.Handler
}

// NewGRPCServer initializes a new gRPC server
func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.MathServiceServer {
	return &gRPCServer{
		add: grpc.NewServer(
			endpoints.Add,
			decodeMathRequest,
			encodeMathResponse,
		),
		subtract: grpc.NewServer(
			endpoints.Subtract,
			decodeMathRequest,
			encodeMathResponse,
		),
		multiply: grpc.NewServer(
			endpoints.Multiply,
			decodeMathRequest,
			encodeMathResponse,
		),
		divide: grpc.NewServer(
			endpoints.Divide,
			decodeMathRequest,
			encodeMathResponse,
		),
	}
}

func (s *gRPCServer) Add(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.MathResponse), nil
}

func (s *gRPCServer) Subtract(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.subtract.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.MathResponse), nil
}

func (s *gRPCServer) Multiply(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.multiply.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.MathResponse), nil
}

func (s *gRPCServer) Divide(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.divide.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.MathResponse), nil
}

func decodeMathRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.MathRequest)
	return endpoints.MathReq{
		NumA: req.NumA,
		NumB: req.NumB,
	}, nil
}

func encodeMathResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.MathResp)
	return &pb.MathResponse{
		Result: resp.Result,
	}, nil
}
