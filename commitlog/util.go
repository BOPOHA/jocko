package commitlog

import (
	"sort"
)

// findSegments returns the nearest segment whose base offset is greater than or equal to the given
// offset.
func findSegment(segments []*segment, offset int64) (*segment, int) {
	n := len(segments)
	idx := sort.Search(n, func(i int) bool {
		return segments[i].nextOffset > offset
	})
	if idx == n {
		return nil, idx
	}
	return segments[idx], idx
}

func roundDown(total, factor int64) int64 {
	return factor * (total / factor)
}
