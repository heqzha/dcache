package rpcserv

import (
	"fmt"
	"net"
	"os"

	"github.com/heqzha/dcache/pb"
	"github.com/heqzha/dcache/rpcserv/handler"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func runRPCServer(port int, register func(*grpc.Server, ...interface{}), services ...interface{}) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
		os.Exit(1)
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	register(grpcServer, services...)

	grpcServer.Serve(lis)
}

func Run(port int) {
	runRPCServer(port, func(grpc *grpc.Server, services ...interface{}) {
		for _, s := range services {
			switch s.(type) {
			case *DCacheService:
				pb.RegisterCacheServServer(grpc, s.(*DCacheService))
			}
		}
	}, new(DCacheService))
}

type DCacheService struct{}

func (s *DCacheService) Get(ctx context.Context, in *pb.GetReq) (*pb.GetRes, error) {
	//TODO
	return nil, nil
}

func (s *DCacheService) Set(ctx context.Context, in *pb.SetReq) (*pb.SetRes, error) {
	//TODO
	return nil, nil
}

func (s *DCacheService) Del(ctx context.Context, in *pb.DelReq) (*pb.DelRes, error) {
	//TODO
	return nil, nil
}

func (s *DCacheService) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterRes, error) {
	if err := handler.Register(in.GetGroup(), in.GetAddr()); err != nil {
		return nil, err
	}
	return &pb.RegisterRes{
		Status: true,
	}, nil
}

func (s *DCacheService) Unregister(ctx context.Context, in *pb.UnregisterReq) (*pb.UnregisterRes, error) {
	if err := handler.Unregister(in.GetGroup(), in.GetAddr()); err != nil {
		return nil, err
	}
	return &pb.UnregisterRes{
		Status: true,
	}, nil
}

func (s *DCacheService) SyncSrvGroup(ctx context.Context, in *pb.SyncSrvGroupReq) (*pb.SyncSrvGroupRes, error) {
	cond, srvGroup, err := handler.SyncSrvGroups(in.GetSrvGroup())
	if err != nil {
		return nil, err
	}
	return &pb.SyncSrvGroupRes{
		Status:    true,
		Condition: int32(cond),
		SrvGroup:  srvGroup,
	}, nil
}

func (s *DCacheService) Ping(ctx context.Context, in *pb.PingReq) (*pb.PingRes, error) {
	srvGroup, err := handler.Ping(in.GetGroup(), in.GetAddr())
	if err != nil {
		return nil, err
	}
	return &pb.PingRes{
		Status:   true,
		SrvGroup: srvGroup,
	}, nil
}
