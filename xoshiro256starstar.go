package goxoshiro256starstar

import (
	"github.com/epes/gosplitmix64"
	"math/rand"
)

// Xoshiro256StarStar is a go translation of https://prng.di.unimi.it/xoshiro256starstar.c
// Implements rand.Source and rand.Source64
type Xoshiro256StarStar struct {
	s0 uint64
	s1 uint64
	s2 uint64
	s3 uint64
}

func (x *Xoshiro256StarStar) Seed(seed int64) {
	sm := gosplitmix64.New(seed)

	x.s0 = sm.Next()
	x.s1 = sm.Next()
	x.s2 = sm.Next()
	x.s3 = sm.Next()
}

func (x *Xoshiro256StarStar) Next() uint64 {
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

func (x *Xoshiro256StarStar) Int63() int64 {
	return int64(x.Uint64() >> 1)
}

func (x *Xoshiro256StarStar) Uint64() uint64 {
	return x.Next()
}

func New(seed int64) *Xoshiro256StarStar {
	x := &Xoshiro256StarStar{}
	x.Seed(seed)
	return x
}

func NewSource() rand.Source {
	return &Xoshiro256StarStar{}
}

func NewSource64() rand.Source64 {
	return &Xoshiro256StarStar{}
}

func rotl(x uint64, k int) uint64 {
	return (x << k) | (x >> (64 - k))
}
