package stats

import "sync/atomic"

// Stats tracks active SSH sessions and total visitor count.
type Stats struct {
	activeSessions atomic.Int64
	totalVisitors  atomic.Int64
}

// Global is the singleton stats tracker.
var Global = &Stats{}

func (s *Stats) Connect() {
	s.activeSessions.Add(1)
	s.totalVisitors.Add(1)
}

func (s *Stats) Disconnect() {
	s.activeSessions.Add(-1)
}

func (s *Stats) Active() int64 {
	return s.activeSessions.Load()
}

func (s *Stats) Total() int64 {
	return s.totalVisitors.Load()
}
