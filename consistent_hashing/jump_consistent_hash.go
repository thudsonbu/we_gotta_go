package main

import "fmt"

func jumpConsistentHash(key uint64, numBuckets int) int {
	var b int64 = -1
	var j int64

	for j < int64(numBuckets) {
		b = j
		key = key*2862933555777941757 + 1
		j = int64(float64((b+1)*(1<<31)) / float64((key>>33)+1))
	}

	return int(b)
}

func main() {
	// We have 5 nodes in our distributed cache system
	numNodes := 5

	// We have a list of cache keys
	cacheKeys := []uint64{123, 456, 789, 101112, 131415}

	// We distribute the keys across the nodes
	for _, key := range cacheKeys {
		node := jumpConsistentHash(key, numNodes)
		fmt.Printf("Key %d is assigned to Node %d\n", key, node)
	}
}
