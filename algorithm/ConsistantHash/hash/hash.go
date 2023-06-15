package hash

import (
	"hash/crc32"
	"sort"
)

type hashRing struct {
	nodes      map[uint32]string
	replicates int
	keys       []uint32
}

func New(replicate int) *hashRing {
	return &hashRing{
		nodes:      make(map[uint32]string),
		replicates: replicate,
		keys:       []uint32{},
	}
}

func (h *hashRing) Add(nodename string) {
	for i := 0; i < h.replicates; i++ {
		hk := crc32.ChecksumIEEE([]byte(nodename + "-" + string(i)))
		h.nodes[hk] = nodename
		h.keys = append(h.keys, hk)
	}
	sort.Slice(h.keys, func(i, j int) bool {
		return h.keys[i] < h.keys[j]
	})
}

func (h *hashRing) Get(key string) string {
	number := crc32.ChecksumIEEE([]byte(key))
	idx := sort.Search(len(h.keys), func(i int) bool {
		return number <= h.keys[i]
	})
	if idx == len(h.keys) {
		idx = 0
	}
	return h.nodes[h.keys[idx]]
}
