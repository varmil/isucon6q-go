package main

import (
	"fmt"
	"html"
	"strconv"
	"strings"
	"sync"

	"github.com/wangjia184/sortedset"
)

// SortedSet contains cmap, the cmap has pointer of reservation as value
type SortedSet struct {
	mx   *sync.RWMutex
	r    *sortedset.SortedSet
	reps *strings.Replacer
}

// NewSortedSet returns the instance
func NewSortedSet() *SortedSet {
	return &SortedSet{mx: &sync.RWMutex{}, r: sortedset.New()}
}

// Store the instance
func (s *SortedSet) Store(keyword string, updateReplacer bool) {
	// s.mx.Lock()
	s.r.AddOrUpdate(keyword, sortedset.SCORE(len(keyword)), true)
	// s.mx.Unlock()

	if updateReplacer {
		s.UpdateReplacer()
	}
}

// LoadAllSortedWords the instances
func (s *SortedSet) LoadAllSortedWords() *[]*string {
	var result []*string

	// s.mx.RLock()
	// defer s.mx.RUnlock()

	for _, foo := range s.r.GetByRankRange(-1, 1, false) {
		bar := foo.Key()
		result = append(result, &bar)
	}

	return &result
}

// Delete the instance
func (s *SortedSet) Delete(keyword string) {
	s.r.Remove(keyword)
	s.UpdateReplacer()
}

// Count all keywords
func (s *SortedSet) Count() int {
	return s.r.GetCount()
}

// UpdateReplacer do strings.NewReplacer() from current keywords
// strings.NewReplacer() 自体が重いのでキャッシュすると爆発的にスコアが上がる
func (s *SortedSet) UpdateReplacer() {
	var kwLinkPairs []string
	sorted := s.LoadAllSortedWords()

	for _, keyword := range *sorted {
		kw := *keyword

		// link 生成。相対パスで十分だった…
		link := fmt.Sprintf("<a href=\"/keyword/%s\">%s</a>",
			pathURIEscape(kw),
			html.EscapeString(kw))

		// NewReplacer用のslice
		kwLinkPairs = append(kwLinkPairs, kw, link)
	}

	s.mx.Lock()
	s.reps = strings.NewReplacer(kwLinkPairs...)
	s.mx.Unlock()
}

// Replace do strings.Replace() from current keywords
func (s *SortedSet) Replace(content string) string {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return s.reps.Replace(content)
}

func toInt64(s string) int64 {
	r, _ := strconv.ParseInt(s, 10, 64)
	return r
}
