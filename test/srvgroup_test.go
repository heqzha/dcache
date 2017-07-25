package test

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"fmt"

	"github.com/heqzha/dcache/core"
)

func TestSrvTable(t *testing.T) {
	tb1 := core.SrvTable{
		"127.0.0.1:1001": 0,
		"127.0.0.1:1002": 0,
		"127.0.0.1:1003": 0,
	}
	tb2 := core.SrvTable{
		"127.0.0.1:1004": 0,
		"127.0.0.1:1005": 0,
	}
	tb1.Add(tb2)
	fmt.Println(tb1.String())

	tb2 = tb1.Clone()
	fmt.Println(tb2.String())
}

func TestSrvGroup(t *testing.T) {
	group := core.SrvGroup{}
	group.NewGroup("test1")
	tb1 := core.SrvTable{
		"127.0.0.1:1001": 0,
		"127.0.0.1:1002": 0,
		"127.0.0.1:1003": 0,
	}
	tb2 := core.SrvTable{
		"127.0.0.1:1004": 0,
		"127.0.0.1:1005": 0,
	}
	group.SetTable("test1", &tb1)

	group.SetTable("test2", &tb2)

	fmt.Println(group.GetTable("test1"))

	fmt.Println(group.GetTable("test2"))
}

func TestSGM(t *testing.T) {
	var (
		sgm  core.SGM
		sgm2 core.SGM
	)
	sgm.Init("group1", "127.0.0.1:1001")
	sgm2.Init("group2", "127.0.0.2:1001")

	for index := 1001; index < 1004; index++ {
		sgm.Register("group1", "127.0.0.1:"+strconv.Itoa(index))
	}
	for index := 1002; index < 1006; index++ {
		sgm.Register("group2", "127.0.0.1:"+strconv.Itoa(index))
	}
	fmt.Println(sgm.GetGroupNames())
	tb1g1, _ := sgm.GetTable("group1")
	tb1g2, _ := sgm.GetTable("group2")
	clk1 := sgm.GetClock()
	fmt.Println(tb1g1.String())
	fmt.Println(tb1g2.String())
	fmt.Println(clk1.ReturnVCString())

	time.Sleep(2 * time.Second)

	sgm.UpdateTimestamp("group1", "127.0.0.1:1001")

	for index := 1001; index < 1003; index++ {
		sgm2.Register("group1", "127.0.0.2:"+strconv.Itoa(index))
	}
	tb2gp1, _ := sgm2.GetTable("group1")
	clk2 := sgm2.GetClock()
	fmt.Println(tb2gp1.String())
	fmt.Println(clk2.ReturnVCString())

	sgm2.Merge(sgm)
	fmt.Println("Condition:", sgm2.CompareReadable(sgm))
	tb2mgp1, _ := sgm2.GetTable("group1")
	tb2mgp2, _ := sgm2.GetTable("group2")
	clk2 = sgm2.GetClock()
	fmt.Println(tb2mgp1.String())
	fmt.Println(tb2mgp2.String())
	fmt.Println(clk2.ReturnVCString())

	sgm.Merge(sgm2)
	fmt.Println("Condition:", sgm.CompareReadable(sgm2))
	tb1mgp1, _ := sgm.GetTable("group1")
	tb1mgp2, _ := sgm.GetTable("group2")
	clk1 = sgm.GetClock()
	fmt.Println(tb1mgp1.String())
	fmt.Println(tb1mgp2.String())
	fmt.Println(clk1.ReturnVCString())
}

func TestSTMDumpLoad(t *testing.T) {
	var (
		sgm  core.SGM
		sgm2 core.SGM
	)
	sgm.Init("group1", "127.0.0.1:1001")
	sgm2.Init("group2", "127.0.0.2:1001")

	for index := 1001; index < 1004; index++ {
		sgm.Register("group1", "127.0.0.1:"+strconv.Itoa(index))
	}

	for index := 1003; index < 1007; index++ {
		sgm.Register("group2", "127.0.0.1:"+strconv.Itoa(index))
	}

	d, err := sgm.Dump()
	if err != nil {
		t.Error(err)
		return
	}

	if err := sgm2.Load(d); err != nil {
		t.Error(err)
		return
	}
	tableg1, _ := sgm2.GetTable("group1")
	tableg2, _ := sgm2.GetTable("group2")
	fmt.Println(tableg1.String())
	fmt.Println(tableg2.String())
	fmt.Println(sgm.CompareReadable(sgm2))
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func TestSTHash(t *testing.T) {
	var (
		sth core.SGHash
		stm core.SGM
	)
	stm.Init("group1", "127.0.0.1:1001")

	for index := 1001; index < 1010; index++ {
		stm.Register("group1", "127.0.0.1:"+strconv.Itoa(index))
	}

	sth.Init(3, nil)
	sth.Load(stm.GetGroup())

	for index := 0; index < 20; index++ {
		key := RandStringBytes(10)
		srv := sth.Pick("group1", key)
		fmt.Println(key, srv)
	}
}
