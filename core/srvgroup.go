package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type SrvTable map[string]int64

func (st *SrvTable) Add(t SrvTable) {
	for addr, v := range t {
		(*st)[addr] = v
	}
}

func (st *SrvTable) Clone() SrvTable {
	nst := make(SrvTable, len(*st))
	for addr, v := range *st {
		nst[addr] = v
	}
	return nst
}

func (st *SrvTable) String() string {
	addrs := []string{}
	for addr, ts := range *st {
		addrs = append(addrs, fmt.Sprintf("%s-%d", addr, ts))
	}
	return strings.Join(addrs, ";")
}

type SrvGroup map[string]*SrvTable

func (s *SrvGroup) NewGroup(group string) (*SrvTable, error) {
	if group == "" {
		return nil, fmt.Errorf("SrvGroup.NewGroup: %s", "missing group name")
	}
	(*s)[group] = &SrvTable{}
	return (*s)[group], nil
}

func (s *SrvGroup) GetTable(group string) (*SrvTable, error) {
	if group == "" {
		return nil, fmt.Errorf("SrvGroup.GetTable: %s", "missing group name")
	}
	table, ok := (*s)[group]
	if ok && table != nil {
		return table, nil
	}
	return s.NewGroup(group)
}

func (s *SrvGroup) SetTable(group string, table *SrvTable) error {
	if group == "" {
		return fmt.Errorf("SrvGroup.SetTable: %s", "missing group name")
	}
	(*s)[group] = table
	return nil
}

type SGM struct {
	localGroup string
	localAddr  string
	group      SrvGroup
	clock      VClock
	mutex      *sync.RWMutex
}

func (s *SGM) Init(localGroup, localAddr string) {
	s.localGroup = localGroup
	s.localAddr = localAddr
	s.group = SrvGroup{}
	s.clock = VClock{}
	s.mutex = &sync.RWMutex{}
	s.Register(localGroup, localAddr)
}

func (s *SGM) Register(group string, addrs ...string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	table, err := s.group.GetTable(group)
	if err != nil {
		return err
	}
	for _, addr := range addrs {
		(*table)[addr] = time.Now().Unix()
	}

	s.clock.Tick(s.localAddr)
	return nil
}

func (s *SGM) Unregister(group, addr string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	table, err := s.group.GetTable(group)
	if err != nil {
		return err
	}
	_, eixst := (*table)[addr]
	if eixst {
		delete((*table), addr)
	}
	s.clock.Tick(s.localAddr)
	return nil
}

func (s *SGM) UpdateTimestamp(group, addr string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	table, err := s.group.GetTable(group)
	if err != nil {
		return err
	}
	_, eixst := (*table)[addr]
	if eixst {
		(*table)[addr] = time.Now().Unix()
	}
	s.clock.Tick(s.localAddr)
	return nil
}

func (s *SGM) GetGroupNames() []string {
	group := []string{}
	for g := range s.group {
		group = append(group, g)
	}
	return group
}

func (s *SGM) Merge(as SGM) (Condition, error) {
	s.mutex.Lock()
	defer func() {
		s.clock.Merge(as.clock)
		s.mutex.Unlock()
	}()
	if s.clock.Compare(as.clock, Equal) {
		//s and as are equal
		return Equal, nil
	} else if s.clock.Compare(as.clock, Concurrent) {
		//s and as are concurrent
		for tg, tt := range as.group {
			st, err := s.group.GetTable(tg)
			if err != nil {
				return Condition(-1), err
			}
			st.Add((*tt))
		}
		return Concurrent, nil
	} else if s.clock.Compare(as.clock, Descendant) {
		//s is older than as
		for tg, tt := range as.group {
			ttc := tt.Clone()
			s.group.SetTable(tg, &ttc)
		}
		return Descendant, nil
	} else if s.clock.Compare(as.clock, Ancestor) {
		//s is newer than t
		return Ancestor, nil
	}
	return Condition(-1), nil
}

func (s *SGM) CompareReadable(as SGM) string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.clock.Compare(as.clock, Equal) {
		return "Equal"
	} else if s.clock.Compare(as.clock, Concurrent) {
		return "Concurrent"
	} else if s.clock.Compare(as.clock, Descendant) {
		return "Descendant"
	} else if s.clock.Compare(as.clock, Ancestor) {
		return "Ancestor"
	}
	return "Unknown"
}

func (s *SGM) GetTable(group string) (*SrvTable, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.group.GetTable(group)
}

func (s *SGM) GetGroup() SrvGroup {
	return s.group
}

func (s *SGM) GetClock() VClock {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.clock
}

func (s *SGM) Dump() ([]byte, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	b := new(bytes.Buffer)
	enc := gob.NewEncoder(b)
	if err := enc.Encode(s.group); err != nil {
		return nil, err
	}
	if err := enc.Encode(s.clock); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (s *SGM) Load(buf []byte) error {
	r := bytes.NewBuffer(buf)
	dec := gob.NewDecoder(r)
	if err := dec.Decode(&s.group); err != nil {
		return err
	}
	if err := dec.Decode(&s.clock); err != nil {
		return err
	}
	return nil
}

type HashFn func(data []byte) uint32

type RepTable map[int]string

type SGHash struct {
	group    map[string]*RepTable
	replicas int
	hashes   map[string][]int
	hashfn   HashFn
	mutex    *sync.RWMutex
}

func (s *SGHash) Init(replicas int, fn HashFn) {
	s.group = map[string]*RepTable{}
	s.replicas = replicas
	s.hashes = map[string][]int{}
	s.hashfn = fn
	s.mutex = &sync.RWMutex{}
	if s.hashfn == nil {
		s.hashfn = crc32.ChecksumIEEE
	}
}

func (s *SGHash) Load(sg SrvGroup) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for tg, tt := range sg {
		for addr := range *tt {
			for i := 0; i < s.replicas; i++ {
				hash := int(s.hashfn([]byte(strconv.Itoa(i) + addr)))
				if s.hashes[tg] == nil {
					s.hashes[tg] = make([]int, 0)
				}
				s.hashes[tg] = append(s.hashes[tg], hash)
				if s.group[tg] == nil {
					s.group[tg] = &RepTable{}
				}
				(*s.group[tg])[hash] = addr
			}
		}
		sort.Ints(s.hashes[tg])
	}
}

func (s *SGHash) Pick(group, key string) string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if len(s.hashes) == 0 || len(s.hashes[group]) == 0 {
		return ""
	}
	hash := int(s.hashfn([]byte(key)))

	idx := sort.Search(len(s.hashes[group]), func(i int) bool { return s.hashes[group][i] >= hash })

	if idx == len(s.hashes[group]) {
		idx = 0
	}
	return (*s.group[group])[s.hashes[group][idx]]
}
