package main

import (
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
func (s *SyncMatchMap) Store(contentHash string, value bool) {
	s.r.Set(contentHash, value)
}

// Load the instance, return nil if not exists
func (s *SyncMatchMap) Load(contentHash string) (bool, bool) {
	t, ok := s.r.Get(contentHash)
	if !ok {
		return false, false
	}
	return t.(bool), true
}

// Delete the instance
func (s *SyncMatchMap) Delete(contentHash string) {
	s.r.Remove(contentHash)
}
