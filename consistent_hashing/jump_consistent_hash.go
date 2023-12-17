package main

import "fmt"

// jumpConsistentHash takes a key and the number of buckets as input.
// It returns the number of the bucket to which the key should be mapped.
func jumpConsistentHash(key uint64, numBuckets int) int {
	// b is the bucket we're currently considering, initialized to -1
	var b int64 = -1
	// j is the next bucket to consider, initialized to 0
	var j int64

	// We keep considering buckets until j is not less than the number of buckets
	for j < int64(numBuckets) {
		// We set b to the current bucket we're considering
		b = j
		// We update the key using a large prime number and incrementing by 1
		key = key*2862933555777941757 + 1
		// We calculate the next bucket to consider using a formula that ensures a
		// uniform distribution
		j = int64(float64((b+1)*(1<<31)) / float64((key>>33)+1))
	}

	// We return the last bucket we considered, which is the bucket the key should
	// be mapped to
	return int(b)
}

func main() {
	// We have 5 nodes in our distributed cache system
	numNodes := 5

	// We have a list of cache keys
	cacheKeys := []uint64{123, 456, 789, 101112, 131415}

	// We distribute the keys across the nodes
	for _, key := range cacheKeys {
		// We use the jumpConsistentHash function to decide which node the key
		// should be assigned to
		node := jumpConsistentHash(key, numNodes)
		// We print the key and the node it's assigned to
		fmt.Printf("Key %d is assigned to Node %d\n", key, node)
	}
}
