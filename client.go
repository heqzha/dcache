package dcache

import (
	"context"
	"sync"
	"time"

	"github.com/heqzha/dcache/pb"
	"github.com/heqzha/dcache/utils"
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

func (c *CacheServClient) Close() error {
	return c.conn.Close()
}

func (c *CacheServClient) Get(group, key string, value interface{}) error {
	res, err := c.get(group, key)
	if err != nil {
		return err
	}
	data := res.GetValue()
	if data != nil {
		return utils.DCacheDecode(value, data)
	}
	return nil
}

func (c *CacheServClient) Set(group, key string, value interface{}) error {
	data, err := utils.DCacheEncode(value)
	if err != nil {
		return err
	}
	_, err = c.set(group, key, data)
	return err
}

func (c *CacheServClient) Del(group, key string, value interface{}) error {
	res, err := c.del(group, key)
	if err != nil {
		return err
	}
	data := res.GetValue()
	return utils.DCacheDecode(value, data)
}

func (c *CacheServClient) GetIfExist(group, key string, value interface{}) error {
	res, err := c.getIfExist(group, key)
	if err != nil {
		return err
	}
	data := res.GetValue()
	if data != nil {
		return utils.DCacheDecode(value, data)
	}
	return KeyNotExistError
}

func (c *CacheServClient) SetWithExpire(group, key string, value interface{}, exp time.Duration) error {
	data, err := utils.DCacheEncode(value)
	if err != nil {
		return err
	}
	_, err = c.setWithExpire(group, key, data, int64(exp))
	return err
}

func (c *CacheServClient) get(group, key string) (*pb.GetRes, error) {
	return c.cli.Get(context.Background(), &pb.GetReq{
		Group: group,
		Key:   key,
	})
}

func (c *CacheServClient) getIfExist(group, key string) (*pb.GetIfExistRes, error) {
	return c.cli.GetIfExist(context.Background(), &pb.GetIfExistReq{
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

func (c *CacheServClient) setWithExpire(group, key string, value []byte, exp int64) (*pb.SetWithExpireRes, error) {
	return c.cli.SetWithExpire(context.Background(), &pb.SetWithExpireReq{
		Group:  group,
		Key:    key,
		Value:  value,
		Expire: exp,
	})
}

func (c *CacheServClient) del(group, key string) (*pb.DelRes, error) {
	return c.cli.Del(context.Background(), &pb.DelReq{
		Group: group,
		Key:   key,
	})
}

func (c *CacheServClient) register(group, addr string) (*pb.RegisterRes, error) {
	return c.cli.Register(context.Background(), &pb.RegisterReq{
		Group: group,
		Addr:  addr,
	})
}

func (c *CacheServClient) unregister(group, addr string) (*pb.UnregisterRes, error) {
	return c.cli.Unregister(context.Background(), &pb.UnregisterReq{
		Group: group,
		Addr:  addr,
	})
}

func (c *CacheServClient) syncSrvGroup(srvgroup []byte) (*pb.SyncSrvGroupRes, error) {
	return c.cli.SyncSrvGroup(context.Background(), &pb.SyncSrvGroupReq{
		SrvGroup: srvgroup,
	})
}

func (c *CacheServClient) ping(group, addr string) (*pb.PingRes, error) {
	return c.cli.Ping(context.Background(), &pb.PingReq{
		Group: group,
		Addr:  addr,
	})
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
