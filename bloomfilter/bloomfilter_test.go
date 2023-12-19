package bloomfilter

import "testing"

func TestBloomFilter(t *testing.T) {
	bf := NewBloomFilter(100, 3)
	bf.Add("hello")
	bf.Add("world")
	if !bf.Contains("hello") {
		t.Fatal("bloom filter should contain hello")
	}
	if !bf.Contains("world") {
		t.Fatal("bloom filter should contain world")
	}
	if bf.Contains("foo") {
		t.Fatal("bloom filter should not contain foo")
	}
}
