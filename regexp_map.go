package main

import (
	"strconv"
	"sync"

	"github.com/wangjia184/sortedset"
)

// SortedSet contains cmap, the cmap has pointer of reservation as value
type SortedSet struct {
	mx *sync.RWMutex
	r  *sortedset.SortedSet
}

// NewSortedSet returns the instance
func NewSortedSet() *SortedSet {
	return &SortedSet{mx: &sync.RWMutex{}, r: sortedset.New()}
}

// Store the instance
func (s *SortedSet) Store(keyword string) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.r.AddOrUpdate(keyword, sortedset.SCORE(len(keyword)), true)
}

// LoadAllSortedWords the instances
func (s *SortedSet) LoadAllSortedWords() *[]*string {
	var result []*string

	s.mx.RLock()
	defer s.mx.RUnlock()

	for _, foo := range s.r.GetByRankRange(-1, 1, false) {
		bar := foo.Key()
		result = append(result, &bar)
	}

	return &result
}

// Delete the instance
func (s *SortedSet) Delete(keyword string) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.r.Remove(keyword)
}

// Count all keywords
func (s *SortedSet) Count() int {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return s.r.GetCount()
}

func toInt64(s string) int64 {
	r, _ := strconv.ParseInt(s, 10, 64)
	return r
}
