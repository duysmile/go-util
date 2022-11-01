package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// HashFunc
// we need a function to convert []byte to uint32 to put it in ring hash
type HashFunc func([]byte) uint32

// Hash
// define a ring hash
type Hash struct {
	hashFunc HashFunc
	replicas int
	keys     []int
	hashMap  map[int]string
}

func NewHash(replicas int, hashFunk HashFunc) *Hash {
	h := &Hash{
		hashFunc: hashFunk,
		replicas: replicas,
		keys:     make([]int, 0),
		hashMap:  make(map[int]string),
	}
	if h.hashFunc == nil {
		h.hashFunc = crc32.ChecksumIEEE
	}

	return h
}

func (h *Hash) IsEmpty() bool {
	return len(h.keys) == 0
}

func (h *Hash) AddMulti(keys ...string) {
	for _, key := range keys {
		for i := 0; i < h.replicas; i++ {
			hashValue := int(h.hashFunc([]byte(key + strconv.Itoa(i))))
			h.keys = append(h.keys, hashValue)
			h.hashMap[hashValue] = key
		}
	}

	sort.Ints(h.keys)
}

func (h *Hash) Add(key string) {
	for i := 0; i < h.replicas; i++ {
		hashValue := int(h.hashFunc([]byte(key + strconv.Itoa(i))))
		h.keys = append(h.keys, hashValue)
		h.hashMap[hashValue] = key
	}

	sort.Ints(h.keys)
}

func (h *Hash) Remove(key string) {
	shouldRemoveValues := make(map[int]bool)
	for i := 0; i < h.replicas; i++ {
		hashValue := int(h.hashFunc([]byte(key + strconv.Itoa(i))))
		shouldRemoveValues[hashValue] = true
		delete(h.hashMap, hashValue)
	}

	newKeys := make([]int, 0, len(h.keys))
	for _, key := range h.keys {
		if !shouldRemoveValues[key] {
			newKeys = append(newKeys, key)
		}
	}

	h.keys = newKeys
	sort.Ints(h.keys)
}

func (h *Hash) Get(key string) string {
	hashValue := int(h.hashFunc([]byte(key)))

	// use binary find the first key in list greater than equal `hashValue`
	idx := sort.Search(len(h.keys), func(i int) bool {
		return h.keys[i] >= int(hashValue)
	})

	// if no keys match above condition, map it with first key
	if idx == len(h.keys) {
		idx = 0
	}

	return h.hashMap[h.keys[idx]]
}
