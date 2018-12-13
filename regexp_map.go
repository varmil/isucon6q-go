package main

import (
	"sort"
	"strconv"

	"github.com/gobwas/glob"
	cmap "github.com/orcaman/concurrent-map"
)

// SyncMap contains cmap, the cmap has pointer of reservation as value
type SyncMap struct {
	r cmap.ConcurrentMap
}

// NewSyncMap returns the instance
func NewSyncMap() *SyncMap {
	return &SyncMap{r: cmap.New()}
}

// Store the instance
func (s *SyncMap) Store(keyword string, value glob.Glob) {
	s.r.Set(keyword, value)
}

// Load the instance, return nil if not exists
func (s *SyncMap) Load(keyword string) (glob.Glob, bool) {
	t, ok := s.r.Get(keyword)
	if !ok {
		return nil, false
	}
	return t.(glob.Glob), true
}

// Has checks if the key exists or not
func (s *SyncMap) Has(keyword string) bool {
	return s.r.Has(keyword)
}

// LoadAll the instances, return nil if not exists
func (s *SyncMap) LoadAllSortedWords() []string {
	var result []string

	for word := range s.r.Items() {
		result = append(result, word)
	}
	sort.Slice(result, func(i, j int) bool {
		return len(result[i]) > len(result[j])
	})

	return result
}

// Delete the instance
func (s *SyncMap) Delete(keyword string) {
	s.r.Remove(keyword)
}

func toString(n int64) string {
	return strconv.FormatInt(n, 10)
}
