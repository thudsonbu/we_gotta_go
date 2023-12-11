package main

import (
	"hash/fnv"
	"sort"
	"strconv"
)

type Node struct {
	Id   int
	Hash uint32
}

type ConsistentHashing struct {
	Nodes []Node
}

func NewConsistentHashing() *ConsistentHashing {
	return &ConsistentHashing{}
}

func (c *ConsistentHashing) AddNode(id int) {
	node := Node{
		Id:   id,
		Hash: hashInt(id),
	}
	c.Nodes = append(c.Nodes, node)
	sort.Slice(c.Nodes, func(i, j int) bool {
		return c.Nodes[i].Hash < c.Nodes[j].Hash
	})
}

func (c *ConsistentHashing) GetNode(key string) Node {
	hash := hashString(key)
	i := sort.Search(len(c.Nodes), func(i int) bool {
		return c.Nodes[i].Hash >= hash
	})

	if i == len(c.Nodes) {
		i = 0
	}

	return c.Nodes[i]
}

func hashInt(i int) uint32 {
	return hashString(strconv.Itoa(i))
}

func hashString(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
