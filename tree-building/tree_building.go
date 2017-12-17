package tree

import (
	"errors"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {

	l := len(records)
	idNodeMap := make(map[int]*Node, l)
	idRecordMap := make(map[int]Record, l)

	for _, r := range records {
		if r.ID < 0 || r.ID >= l {
			return nil, errors.New("invalid ID")
		}
		idNodeMap[r.ID] = &Node{ID: r.ID}
		idRecordMap[r.ID] = r
	}

	var root *Node = nil
	for id := l - 1; id >= 0; id-- {

		r := idRecordMap[id]
		c := idNodeMap[id]

		if r.ID == r.Parent {
			if root != nil {
				return nil, errors.New("cycle")
			}
			root = c
			continue
		}

		if r.ID < r.Parent {
			return nil, errors.New("invalid ID")
		}

		p := idNodeMap[r.Parent]
		p.Children = append([]*Node{c}, p.Children...)
	}

	return root, nil
}
