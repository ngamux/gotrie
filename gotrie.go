package gotrie

import (
	"strings"
	"sync"
)

type Trie struct {
	root   *node
	size   uint
	mu     sync.RWMutex
	config Config
}

type node struct {
	children map[string]*node
	data     interface{}
}

func NewTrie(config Config) *Trie {
	return &Trie{
		root: &node{
			children: make(map[string]*node),
		},
		config: config,
	}
}

func (trie *Trie) buildKey(key string) []string {
	return strings.Split(key, trie.config.Separator)
}

func (trie *Trie) Put(key string, data interface{}) {
	trie.mu.Lock()
	keys := trie.buildKey(key)
	root := trie.root
	for _, key := range keys {
		child := root.children[key]
		if child == nil {
			child = &node{
				children: make(map[string]*node),
			}
			root.children[key] = child
		}

		root = child
	}

	root.data = data
	trie.size++

	trie.mu.Unlock()
}

func (trie *Trie) Get(key string) interface{} {
	trie.mu.RLock()
	keys := trie.buildKey(key)
	defer trie.mu.RUnlock()
	root := trie.root
	for _, key := range keys {
		child := root.children[key]
		if child == nil {
			return nil
		}

		root = child
	}

	return root.data
}

func (trie *Trie) Size() uint {
	trie.mu.RLock()
	defer trie.mu.RUnlock()
	return trie.size
}
