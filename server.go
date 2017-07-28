package dcache

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/heqzha/dcache/core"
	"github.com/heqzha/dcache/pb"
	"github.com/heqzha/dcache/process"
	"github.com/heqzha/goutils/logger"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	grpcServer *grpc.Server
)

func Run(port int, dc *DCache) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for _ = range signalChan {
			fmt.Println("Unregistering...")
			msgQ.Push("srvgroup", map[string]interface{}{
				"type": "unregister",
			})
			//handle closing
			<-*cleanUpFlag
			cliPool.Close()
			process.StopAll()
			grpcServer.GracefulStop()
			fmt.Println("Graceful Closed.")
		}
	}()

	sgm.RegisterLocalAddr(dc.localGroup, dc.localAddr)
	sgh.Load(sgm.GetGroup())

	process.MaintainSvrGroups()

	if !dc.isRoot {
		root, err := cliPool.GetOrAdd(dc.rootAddr)
		if err != nil {
			panic(fmt.Sprintf("failed to connect root node: %s", err.Error()))
		}
		res, err := root.Register(dc.localGroup, dc.localAddr)
		if err != nil {
			panic(fmt.Sprintf("failed to register to root node: %s", err.Error()))
		} else if !res.GetStatus() {
			panic(fmt.Sprintf("register denied"))
		}
	}
	//Register and run server
	runRPCServer(port, func(grpc *grpc.Server, services ...interface{}) {
		for _, s := range services {
			switch s.(type) {
			case *DCacheService:
				pb.RegisterCacheServServer(grpc, s.(*DCacheService))
			}
		}
	}, newDCacheService(dc))
}

func runRPCServer(port int, register func(*grpc.Server, ...interface{}), services ...interface{}) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
		os.Exit(1)
	}
	opts := []grpc.ServerOption{}
	grpcServer = grpc.NewServer(opts...)
	register(grpcServer, services...)

	grpcServer.Serve(lis)
}

type DCacheService struct {
	cache *DCache
}

func newDCacheService(cache *DCache) *DCacheService {
	return &DCacheService{
		cache: cache,
	}
}

func (s *DCacheService) Get(ctx context.Context, in *pb.GetReq) (*pb.GetRes, error) {
	value, err := s.cache.Get(in.GetGroup(), in.GetKey())
	if err != nil {
		return nil, err
	}
	return &pb.GetRes{
		Status: true,
		Value:  value,
	}, nil
}

func (s *DCacheService) Set(ctx context.Context, in *pb.SetReq) (*pb.SetRes, error) {
	if err := s.cache.Set(in.GetGroup(), in.GetKey(), in.GetValue()); err != nil {
		return nil, err
	}
	return &pb.SetRes{
		Status: true,
	}, nil
}

func (s *DCacheService) Del(ctx context.Context, in *pb.DelReq) (*pb.DelRes, error) {
	value, err := s.cache.Del(in.GetGroup(), in.GetKey())
	if err != nil {
		return nil, err
	}
	return &pb.DelRes{
		Status: true,
		Value:  value,
	}, nil
}

func (s *DCacheService) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterRes, error) {
	logger.Debug("DCacheService.Register", fmt.Sprintf("%s: %s", in.GetGroup(), in.GetAddr()))
	if err := sgm.Register(in.GetGroup(), in.GetAddr()); err != nil {
		return nil, err
	}
	Sync()
	return &pb.RegisterRes{
		Status: true,
	}, nil
}

func (s *DCacheService) Unregister(ctx context.Context, in *pb.UnregisterReq) (*pb.UnregisterRes, error) {
	logger.Debug("DCacheService.Unregister", fmt.Sprintf("%s: %s", in.GetGroup(), in.GetAddr()))
	if err := sgm.Unregister(in.GetGroup(), in.GetAddr()); err != nil {
		return nil, err
	}
	cliPool.Del(in.GetAddr())
	Sync()
	return &pb.UnregisterRes{
		Status: true,
	}, nil
}

func (s *DCacheService) SyncSrvGroup(ctx context.Context, in *pb.SyncSrvGroupReq) (*pb.SyncSrvGroupRes, error) {
	logger.Debug("DCacheService.SyncSrvGroup", "")
	tmpSGM := core.SGM{}
	tmpSGM.Load(in.GetSrvGroup())
	logger.Debug("SGM.Merge", fmt.Sprintf("Before Merge: %s", sgm.CompareReadable(tmpSGM)))
	cond, err := sgm.Merge(tmpSGM)
	if err != nil {
		return nil, err
	}
	logger.Debug("SGM.Merge", fmt.Sprintf("After Merge: %s", sgm.CompareReadable(tmpSGM)))
	srvGroup, err := sgm.Dump()
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
	logger.Debug("DCacheService.Ping", fmt.Sprintf("%s-%s", in.GetGroup(), in.GetAddr()))
	srvGroup, err := sgm.Dump()
	if err != nil {
		return nil, err
	}
	return &pb.PingRes{
		Status:   true,
		SrvGroup: srvGroup,
	}, nil
}
