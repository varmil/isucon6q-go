package main

// NOTE: { eid: { keyword: isMatch } }
// のようなネストマップは逆にかなり低速

import (
	"strconv"

	cmap "github.com/orcaman/concurrent-map"
)

// SyncMatchMap contains cmap
// contentHash: contentHash + keyword
type SyncMatchMap struct {
	r cmap.ConcurrentMap
}

// NewSyncMatchMap returns the instance
func NewSyncMatchMap() *SyncMatchMap {
	return &SyncMatchMap{r: cmap.New()}
}

// Store the instance
func (s *SyncMatchMap) Store(entryID int, content string) {
	s.r.Set(strconv.Itoa(entryID), content)
}

// Load the instance, return nil if not exists
func (s *SyncMatchMap) Load(entryID int) (string, bool) {
	t, ok := s.r.Get(strconv.Itoa(entryID))
	if !ok {
		return "", false
	}
	return t.(string), true
}

// Delete the instance
func (s *SyncMatchMap) Delete(entryID int) {
	s.r.Remove(strconv.Itoa(entryID))
}
