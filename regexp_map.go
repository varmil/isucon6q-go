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

type KeywordGlobTuple struct {
	keyword string
	glob    *glob.Glob
}

// type keywordsSyncSlice struct {
// 	mx   *sync.RWMutex
// 	data []string
// }

// var keywordsCache keywordsSyncSlice

// NewSyncMap returns the instance
func NewSyncMap() *SyncMap {
	// keywordsCache = keywordsSyncSlice{mx: &sync.RWMutex{}}
	return &SyncMap{r: cmap.New()}
}

// Store the instance
func (s *SyncMap) Store(keyword string, value *glob.Glob) {
	s.r.Set(keyword, value)

	// keywordsCache.mx.Lock()
	// keywordsCache.data = append(keywordsCache.data, keyword)
	// keywordsCache.mx.Unlock()
}

// Load the instance, return nil if not exists
func (s *SyncMap) Load(keyword string) (*glob.Glob, bool) {
	t, ok := s.r.Get(keyword)
	if !ok {
		return nil, false
	}
	return t.(*glob.Glob), true
}

// Has checks if the key exists or not
func (s *SyncMap) Has(keyword string) bool {
	return s.r.Has(keyword)
}

// LoadAllSortedWords the instances
func (s *SyncMap) LoadAllSortedWords() []*KeywordGlobTuple {
	var result []*KeywordGlobTuple

	for word, regexp := range s.r.Items() {
		result = append(result, &KeywordGlobTuple{
			keyword: word,
			glob:    regexp.(*glob.Glob),
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return len(result[i].keyword) > len(result[j].keyword)
	})

	return result
}

// Delete the instance
func (s *SyncMap) Delete(keyword string) {
	s.r.Remove(keyword)

	// keywordsCache.mx.Lock()
	// keywordsCache.data = append(keywordsCache.data, keyword)
	// keywordsCache.mx.Unlock()
}

func toString(n int64) string {
	return strconv.FormatInt(n, 10)
}
