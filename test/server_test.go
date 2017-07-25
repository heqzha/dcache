package test

import (
	"context"
	"testing"

	"github.com/heqzha/dcache/rpcserv"
)

func TestDCacheService(t *testing.T) {
	ser := rpcserv.DCacheService{}
	ser.Get(context.Background(), nil)
	ser.Set(context.Background(), nil)
	ser.Del(context.Background(), nil)
	ser.Register(context.Background(), nil)
	ser.Unregister(context.Background(), nil)
	ser.SyncSrvGroup(context.Background(), nil)
}
