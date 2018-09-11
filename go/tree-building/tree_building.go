// Package tree restores hierarchical relationships between tree elements stored
// with parental relationship information only.
package tree

import (
	"errors"
	"fmt"
)

// Record is an element of the hierarcy tree in a relational database.
type Record struct {
	ID, Parent int
}

// Node is an element of the hiearachy tree created by build.
type Node struct {
	ID       int
	Children []*Node
}

// Mismatch has a purpose with which I am unfamiliar.
type Mismatch struct{}

// Error returns "c" regardless of the contents of m.
func (m Mismatch) Error() string {
	return "c"
}

// Build reconstitutes a tree hierarchy from a set of records
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	// What is the reason for the map "dupe"
	dupe := make(map[int]*Record)

	// Looks like it is finding duplicate records.
	for _, r := range records {
		temp := dupe[r.ID]

		if temp == nil {
			dupe[r.ID] = &r
			continue
		}

		if temp.Parent == r.Parent {
			return nil, fmt.Errorf("Node is a duplicate - {ID: %d, Parent: %d}", r.ID, r.Parent)
		}
	}

	root := &Node{}
	todo := []*Node{root}
	n := 1

	for {
		if len(todo) == 0 {
			break
		}
		newTodo := []*Node(nil)
		for _, c := range todo {
			for _, r := range records {
				if r.Parent == c.ID {
					if r.ID < c.ID {
						return nil, errors.New("a")
					} else if r.ID == c.ID {
						if r.ID != 0 {
							return nil, fmt.Errorf("b")
						}
					} else {
						n++
						switch len(c.Children) {
						case 0:
							nn := &Node{ID: r.ID}
							c.Children = []*Node{nn}
							newTodo = append(newTodo, nn)
						case 1:
							nn := &Node{ID: r.ID}
							if c.Children[0].ID < r.ID {
								c.Children = []*Node{c.Children[0], nn}
								newTodo = append(newTodo, nn)
							} else {
								c.Children = []*Node{nn, c.Children[0]}
								newTodo = append(newTodo, nn)
							}
						default:
							nn := &Node{ID: r.ID}
							newTodo = append(newTodo, nn)
						breakpoint:
							for range []bool{false} {
								for i, cc := range c.Children {
									if cc.ID > r.ID {
										a := make([]*Node, len(c.Children)+1)
										copy(a, c.Children[:i])
										copy(a[i+1:], c.Children[i:])
										copy(a[i:i+1], []*Node{nn})
										c.Children = a
										break breakpoint
									}
								}
								c.Children = append(c.Children, nn)
							}
						}
					}
				}
			}
		}
		todo = newTodo
	}

	if n != len(records) {
		return nil, Mismatch{}
	}

	if err := chk(root, len(records)); err != nil {
		return nil, err
	}

	return root, nil
}

func chk(n *Node, m int) (err error) {
	if n.ID > m {
		return fmt.Errorf("z")
	} else if n.ID == m {
		return fmt.Errorf("y")
	} else {
		for i := 0; i < len(n.Children); i++ {
			err = chk(n.Children[i], m)
			if err != nil {
				return
			}
		}
		return
	}
}
