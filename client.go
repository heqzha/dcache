package dcache

import (
	"context"
	"sync"
	"time"

	"github.com/heqzha/dcache/pb"
	"github.com/heqzha/goutils/logger"
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

func (c *CacheServClient) get(group, key string) (*pb.GetRes, error) {
	return c.cli.Get(context.Background(), &pb.GetReq{
		Group: group,
		Key:   key,
	})
}

func (c *CacheServClient) set(group, key string, value []byte) (*pb.SetRes, error) {
	return c.cli.Set(context.Background(), &pb.SetReq{
		Group: group,
		Key:   key,
		Value: value,
	})
}

func (c *CacheServClient) del(group, key string) (*pb.DelRes, error) {
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

type CSClientPool struct {
	pool  map[string]*CacheServClient
	mutex *sync.RWMutex
}

func (p *CSClientPool) Init() {
	p.pool = make(map[string]*CacheServClient)
	p.mutex = &sync.RWMutex{}
}

func (p *CSClientPool) GetOrAdd(addr string) (*CacheServClient, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	_, ok := p.pool[addr]
	if !ok {
		p.pool[addr] = new(CacheServClient)
		err := p.pool[addr].NewRPCClient(addr, time.Minute)
		if err != nil {
			return nil, err
		}
	}
	return p.pool[addr], nil
}

func (p *CSClientPool) Del(addr string) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	c, ok := p.pool[addr]
	defer func() {
		delete(p.pool, addr)
	}()
	if ok && c != nil {
		return c.Close()
	}
	return nil
}

func (p *CSClientPool) Len() int {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return len(p.pool)
}

func (p *CSClientPool) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for addr, cli := range p.pool {
		if err := cli.Close(); err != nil {
			logger.Error("CacheServClient.Close", err.Error())
		}
		delete(p.pool, addr)
	}
}
