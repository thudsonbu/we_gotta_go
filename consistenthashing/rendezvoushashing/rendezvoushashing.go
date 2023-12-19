package rendezvous_hashing

import (
	"crypto/sha256"
	"encoding/binary"
)

func hash(s string) uint64 {
	h := sha256.New()
	h.Write([]byte(s))
	return binary.BigEndian.Uint64(h.Sum(nil))
}

func rendezvousHash(key string, nodes []string) string {
	maxHash := uint64(0)
	maxNode := ""

	for _, node := range nodes {
		h := hash(key + node)
		if h > maxHash {
			maxHash = h
			maxNode = node
		}
	}

	return maxNode
}
