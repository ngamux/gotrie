package gotrie

import (
	"reflect"
	"strings"
	"testing"
)

var dummies = []struct {
	Key  string
	Data interface{}
}{
	{"name", "john"},
	{"names", []string{"john", "chena"}},
}

var keyPathDummies = []struct {
	Key  string
	Data interface{}
}{
	{"/api/v2/users/:id", "john"},
	{"/api/v2/users", []string{"john", "chena"}},
}

func TestNewTrie(t *testing.T) {
	trie := NewTrie()

	if trie.Size() != 0 {
		t.Errorf("trie size should 0, got %d", trie.Size())
	}
}

func TestNewTrieWithConfig(t *testing.T) {
	trie := NewTrie(Config{})

	if trie.Size() != 0 {
		t.Errorf("trie size should 0, got %d", trie.Size())
	}
}

func TestPut(t *testing.T) {
	trie := NewTrie(Config{})

	for _, dummy := range dummies {
		trie.Put(dummy.Key, dummy.Data)
	}

	if trie.Size() != uint(len(dummies)) {
		t.Errorf("trie size should %d, got %d", len(dummies), trie.Size())
	}
}

func TestGet(t *testing.T) {
	trie := NewTrie(Config{})

	for _, dummy := range dummies {
		trie.Put(dummy.Key, dummy.Data)
	}

	for _, dummy := range dummies {
		value := trie.Get(dummy.Key)
		if !reflect.DeepEqual(value, dummy.Data) {
			t.Errorf("trie should contains %s, got %s", dummy.Data, value)
		}
	}
}

func TestBuildKey(t *testing.T) {
	trie := NewTrie(Config{})

	for _, dummy := range dummies {
		keys := trie.buildKey(dummy.Key)

		if len(keys) != len(dummy.Key) {
			t.Errorf("trie key length should %d, got %d", len(keys), len(dummy.Key))
		}
	}

	trie = NewTrie(Config{
		Separator: "/",
	})

	for _, dummy := range keyPathDummies {
		keys := trie.buildKey(dummy.Key)
		t.Log(keys)

		if len(keys) != len(strings.Split(dummy.Key, "/")) {
			t.Errorf("trie key length should %d, got %d", len(keys), len(dummy.Key))
		}
	}
}
