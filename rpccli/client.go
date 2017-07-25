package rpccli

import (
	"context"
	"time"

	"github.com/heqzha/dcache/pb"
	"google.golang.org/grpc"
)

type CacheServClient struct {
	conn *grpc.ClientConn
	cli  pb.CacheServClient
}

func (c *CacheServClient) NewRPCClient(addr string, timeout time.Duration) error {
	conn, err := grpc.Dial(addr, grpc.WithBlock(), grpc.WithTimeout(timeout), grpc.WithInsecure())
	if err != nil {
		return err
	}
	c.conn = conn
	c.cli = pb.NewCacheServClient(conn)
	return nil
}

func (c *CacheServClient) Get(group, key string) (*pb.GetRes, error) {
	return c.cli.Get(context.Background(), &pb.GetReq{
		Group: group,
		Key:   key,
	})
}

func (c *CacheServClient) Set(group, key string, value []byte) (*pb.SetRes, error) {
	return c.cli.Set(context.Background(), &pb.SetReq{
		Group: group,
		Key:   key,
		Value: value,
	})
}

func (c *CacheServClient) Del(group, key string) (*pb.DelRes, error) {
	return c.cli.Del(context.Background(), &pb.DelReq{
		Group: group,
		Key:   key,
	})
}

func (c *CacheServClient) Register(group, addr string) (*pb.RegisterRes, error) {
	return c.cli.Register(context.Background(), &pb.RegisterReq{
		Group: group,
		Addr:  addr,
	})
}

func (c *CacheServClient) Unregister(group, addr string) (*pb.UnregisterRes, error) {
	return c.cli.Unregister(context.Background(), &pb.UnregisterReq{
		Group: group,
		Addr:  addr,
	})
}

func (c *CacheServClient) SyncSrvGroup(srvgroup []byte) (*pb.SyncSrvGroupRes, error) {
	return c.cli.SyncSrvGroup(context.Background(), &pb.SyncSrvGroupReq{
		SrvGroup: srvgroup,
	})
}

func (c *CacheServClient) Ping(group, addr string) (*pb.PingRes, error) {
	return c.cli.Ping(context.Background(), &pb.PingReq{
		Group: group,
		Addr:  addr,
	})
}

func (c *CacheServClient) Close() error {
	return c.conn.Close()
}

type CSClientPool map[string]*CacheServClient

func (p CSClientPool) Add(addr string) (*CacheServClient, error) {
	_, ok := p[addr]
	if !ok {
		p[addr] = new(CacheServClient)
		err := p[addr].NewRPCClient(addr, time.Minute)
		if err != nil {
			return nil, err
		}
	}
	return p[addr], nil
}

func (p CSClientPool) Get(addr string) (*CacheServClient, error) {
	_, ok := p[addr]
	if !ok {
		return p.Add(addr)
	}
	return p[addr], nil
}

func (p CSClientPool) Del(addr string) error {
	c, ok := p[addr]
	defer func() {
		delete(p, addr)
	}()
	if ok && c != nil {
		return c.Close()
	}
	return nil
}

func (p CSClientPool) Len() int {
	return len(p)
}
