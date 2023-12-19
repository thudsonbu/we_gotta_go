package bloomfilter

import (
	"crypto/sha256"
	"encoding/binary"
)

// hash returns a hash of the string s
func hash(s string) uint64 {
	h := sha256.New()
	h.Write([]byte(s))
	return binary.BigEndian.Uint64(h.Sum(nil))
}

// bloomFilter is a bloom filter
type bloomFilter struct {
	m uint64
	k uint64
	b []bool
}

// NewBloomFilter returns a new bloom filter with m bits and k hash functions
func NewBloomFilter(m uint64, k uint64) *bloomFilter {
	return &bloomFilter{
		m: m,
		k: k,
		b: make([]bool, m),
	}
}

// Add adds the string s to the bloom filter
func (bf *bloomFilter) Add(s string) {
	for i := uint64(0); i < bf.k; i++ {
		bf.b[hash(s+string(rune(i)))%bf.m] = true
	}
}

// Contains returns true if the string s is in the bloom filter
func (bf *bloomFilter) Contains(s string) bool {
	for i := uint64(0); i < bf.k; i++ {
		if !bf.b[hash(s+string(rune(i)))%bf.m] {
			return false
		}
	}
	return true
}
