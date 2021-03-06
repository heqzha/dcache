package dcache

import (
	"time"

	"github.com/bluele/gcache"
	"github.com/heqzha/dcache/utils"
)

type DCache struct {
	cache      gcache.Cache
	isRoot     bool
	rootAddr   string
	localGroup string
	localAddr  string
}

func (dc *DCache) Set(group, key string, value []byte) error {
	addr, err := sgh.Pick(group, key)
	if err != nil {
		return err
	}
	if addr == dc.localAddr {
		return dc.cache.Set(key, value)
	}
	cli, err := cliPool.GetOrAdd(addr)
	if err != nil {
		return err
	}
	if _, err := cli.set(group, key, value); err != nil {
		return err
	}
	return nil
}

func (dc *DCache) SetWithExpire(group, key string, value []byte, exp int64) error {
	addr, err := sgh.Pick(group, key)
	if err != nil {
		return err
	}
	if addr == dc.localAddr {
		return dc.cache.SetWithExpire(key, value, time.Duration(exp))
	}
	cli, err := cliPool.GetOrAdd(addr)
	if err != nil {
		return err
	}
	if _, err := cli.setWithExpire(group, key, value, exp); err != nil {
		return err
	}
	return nil
}

func (dc *DCache) Del(group, key string) ([]byte, error) {
	addr, err := sgh.Pick(group, key)
	if err != nil {
		return nil, err
	}
	if addr == dc.localAddr {
		value, err := dc.cache.Get(key)
		if err != nil {
			return nil, err
		}
		dc.cache.Remove(key)
		return value.([]byte), nil
	}
	cli, err := cliPool.GetOrAdd(addr)
	if err != nil {
		return nil, err
	}
	res, err := cli.del(group, key)
	if err != nil {
		return nil, err
	}
	return res.GetValue(), nil
}

func (dc *DCache) Get(group, key string) ([]byte, error) {
	addr, err := sgh.Pick(group, key)
	if err != nil {
		return nil, err
	}
	if addr == dc.localAddr {
		value, err := dc.cache.Get(key)
		if err != nil {
			return nil, err
		} else if value != nil {
			return value.([]byte), nil
		}
		return nil, nil
	}
	cli, err := cliPool.GetOrAdd(addr)
	if err != nil {
		return nil, err
	}
	res, err := cli.get(group, key)
	if err != nil {
		return nil, err
	}
	return res.GetValue(), nil
}

func (dc *DCache) GetIfExist(group, key string) ([]byte, error) {
	addr, err := sgh.Pick(group, key)
	if err != nil {
		return nil, err
	}
	if addr == dc.localAddr {
		value, err := dc.cache.GetIFPresent(key)
		if err != nil {
			if err == gcache.KeyNotFoundError {
				return nil, nil
			}
			return nil, err
		} else if value != nil {
			return value.([]byte), nil
		}
		return nil, nil
	}
	cli, err := cliPool.GetOrAdd(addr)
	if err != nil {
		return nil, err
	}
	res, err := cli.get(group, key)
	if err != nil {
		return nil, err
	}
	return res.GetValue(), nil
}

func (dc *DCache) Cache() gcache.Cache {
	return dc.cache
}

func (dc *DCache) IsRoot() bool {
	return dc.isRoot
}

func (dc *DCache) RootAddr() string {
	return dc.rootAddr
}

func (dc *DCache) LocalGroup() string {
	return dc.localGroup
}

func (dc *DCache) LocalAddr() string {
	return dc.localAddr
}

type DCacheBuilder struct {
	cacheBuilder *gcache.CacheBuilder
	isRoot       bool
	rootAddr     string
	localGroup   string
	localAddr    string
}

func New(size int) *DCacheBuilder {
	if size <= 0 {
		panic("dcache.New: size must larger than 0")
	}
	return &DCacheBuilder{
		cacheBuilder: gcache.New(size),
		isRoot:       true,
		rootAddr:     "127.0.0.1:10000",
		localGroup:   "default",
		localAddr:    "127.0.0.1:10000",
	}
}

func (b *DCacheBuilder) Simple() *DCacheBuilder {
	b.cacheBuilder.Simple()
	return b
}

func (b *DCacheBuilder) LFU() *DCacheBuilder {
	b.cacheBuilder.LFU()
	return b
}

func (b *DCacheBuilder) LRU() *DCacheBuilder {
	b.cacheBuilder.LRU()
	return b
}

func (b *DCacheBuilder) ARC() *DCacheBuilder {
	b.cacheBuilder.ARC()
	return b
}

func (b *DCacheBuilder) LoaderFunc(loaderFunc func(interface{}) (interface{}, error)) *DCacheBuilder {
	b.cacheBuilder.LoaderFunc(func(key interface{}) (interface{}, error) {
		value, err := loaderFunc(key)
		if err != nil {
			return nil, err
		} else if value != nil {
			data, err := utils.DCacheEncode(value)
			if err != nil {
				return nil, err
			}
			return data, nil
		}
		return nil, nil
	})
	return b
}

func (b *DCacheBuilder) LoaderExpireFunc(loaderExpireFunc func(interface{}) (interface{}, *time.Duration, error)) *DCacheBuilder {
	b.cacheBuilder.LoaderExpireFunc(loaderExpireFunc)
	return b
}

func (b *DCacheBuilder) EvictedFunc(evictedFunc func(interface{}, interface{})) *DCacheBuilder {
	b.cacheBuilder.EvictedFunc(evictedFunc)
	return b
}

func (b *DCacheBuilder) AddedFunc(addedFunc func(interface{}, interface{})) *DCacheBuilder {
	b.cacheBuilder.AddedFunc(addedFunc)
	return b
}

func (b *DCacheBuilder) IsRoot(isRoot bool) *DCacheBuilder {
	b.isRoot = isRoot
	return b
}

func (b *DCacheBuilder) RootAddr(addr string) *DCacheBuilder {
	b.rootAddr = addr
	return b
}

func (b *DCacheBuilder) LocalAddr(addr string) *DCacheBuilder {
	b.localAddr = addr
	return b
}

func (b *DCacheBuilder) LocalGroup(group string) *DCacheBuilder {
	b.localGroup = group
	return b
}

func (b *DCacheBuilder) Build() *DCache {
	dc := &DCache{
		cache:      b.cacheBuilder.Build(),
		isRoot:     b.isRoot,
		rootAddr:   b.rootAddr,
		localGroup: b.localGroup,
		localAddr:  b.localAddr,
	}
	return dc
}
