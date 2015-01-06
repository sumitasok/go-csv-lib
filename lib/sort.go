package lib

import "sort"

type By func(r1, r2 *Row) bool

func (by By) Sort(rows []Row) {
	rs := &RowSorter{
		rows: rows,
		by:   by,
	}
	sort.Sort(rs)
}

type RowSorter struct {
	rows []Row
	by   func(r1, r2 *Row) bool
}

// Len is part of sort.Interface.
func (s *RowSorter) Len() int {
	return len(s.rows)
}

// Swap is part of sort.Interface.
func (s *RowSorter) Swap(i, j int) {
	s.rows[i], s.rows[j] = s.rows[j], s.rows[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *RowSorter) Less(i, j int) bool {
	return s.by(&s.rows[i], &s.rows[j])
}
