package goxoshiro256starstar

import (
	"github.com/epes/gosplitmix64"
	"math/rand"
	"sync"
)

// LockedXoshiro256StarStar is a go translation of https://prng.di.unimi.it/xoshiro256starstar.c
// Implements rand.Source and rand.Source64
// Concurrency-safe
type LockedXoshiro256StarStar struct {
	s0  uint64
	s1  uint64
	s2  uint64
	s3  uint64
	mtx sync.Mutex
}

func (x *LockedXoshiro256StarStar) Seed(seed int64) {
	x.mtx.Lock()
	defer x.mtx.Unlock()

	sm := gosplitmix64.New(seed)

	x.s0 = sm.Next()
	x.s1 = sm.Next()
	x.s2 = sm.Next()
	x.s3 = sm.Next()
}

func (x *LockedXoshiro256StarStar) Next() uint64 {
	x.mtx.Lock()
	defer x.mtx.Unlock()

	result := rotl(x.s1*5, 7) * 9
	t := x.s1 << 17

	x.s2 ^= x.s0
	x.s3 ^= x.s1
	x.s1 ^= x.s2
	x.s0 ^= x.s3

	x.s2 ^= t
	x.s3 = rotl(x.s3, 45)

	return result
}

func (x *LockedXoshiro256StarStar) Int63() int64 {
	return int64(x.Uint64() >> 1)
}

func (x *LockedXoshiro256StarStar) Uint64() uint64 {
	return x.Next()
}

func NewLocked(seed int64) *LockedXoshiro256StarStar {
	x := &LockedXoshiro256StarStar{}
	x.Seed(seed)
	return x
}

func NewLockedSource() rand.Source {
	return &LockedXoshiro256StarStar{}
}

func NewLockedSource64() rand.Source64 {
	return &LockedXoshiro256StarStar{}
}
