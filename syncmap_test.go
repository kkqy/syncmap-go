package syncmap_test

import (
	"fmt"
	"testing"

	"github.com/kkqy/syncmap-go"
)

func TestSyncMap(t *testing.T) {
	p := syncmap.SyncMap[int, int]{}
	key := 1
	p.Store(key, 1)
	fmt.Println(p.Load(key))
}
