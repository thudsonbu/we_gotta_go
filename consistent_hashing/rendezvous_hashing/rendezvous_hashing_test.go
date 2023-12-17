package rendezvous_hashing

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

type Node struct {
	name  string
	items []string
}

func TestRendezvousHashing(t *testing.T) {
	nodes := []Node{
		{
			name:  "node1",
			items: []string{},
		},
		{
			name:  "node2",
			items: []string{},
		},
		{
			name:  "node3",
			items: []string{},
		},
	}

	nodeNames := []string{}
	for _, node := range nodes {
		nodeNames = append(nodeNames, node.name)
	}

	for i := 0; i < 1000000; i++ {
		key := fmt.Sprint(rand.Int())
		nodeName := rendezvousHash(key, nodeNames)
		for i, node := range nodes {
			if node.name == nodeName {
				nodes[i].items = append(nodes[i].items, key)
			}
		}
	}

	lengthDiff := 0
	for i := 1; i < len(nodes)-1; i++ {
		lengthDiff += int(math.Abs(
			float64(len(nodes[i].items)) - float64(len(nodes[i+1].items)),
		))
	}

	fmt.Println("length diff:", lengthDiff)
	if lengthDiff > 10000 {
		t.Errorf("length diff is too big: %d", lengthDiff)
	}
}
