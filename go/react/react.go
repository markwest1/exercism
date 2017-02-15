package react

const testVersion = 4

// Node implements the Cell, InputCell and ComputeCell interfaces
type Node struct {
	v         int
	c1        Cell
	c2        Cell
	f1        func(int) int
	f2        func(int, int) int
	callbacks map[int]func(int)
}

// React implements the Reactor interface
type React struct{}

// New creates a new reactor
func New() *React {
	return &React{}
}

// CreateInput is part of the Reactor interface
func (r *React) CreateInput(v int) InputCell {
	return &Node{v, nil, nil, nil, nil, nil}
}

// CreateCompute1 is part of the Reactor interface
func (r *React) CreateCompute1(c Cell, f func(int) int) ComputeCell {
	node := Node{f(c.Value()), c, nil, f, nil, nil}

	// When value of the input changes, update the node's value
	if n, ok := c.(*Node); ok {
		n.AddCallback(func(int) {
			node.SetValue(node.Value())
		})
	}

	return &node
}

// CreateCompute2 is part of the Reactor interface
func (r *React) CreateCompute2(c1, c2 Cell, f func(int, int) int) ComputeCell {
	node := Node{f(c1.Value(), c2.Value()), c1, c2, nil, f, nil}

	// When the value of the first input changes, update the node's value
	if n1, ok := c1.(*Node); ok {
		n1.AddCallback(func(int) {
			node.SetValue(node.Value())
		})
	}

	// When the value of the second input changes, update the node's value
	if n2, ok := c2.(*Node); ok {
		n2.AddCallback(func(int) {
			node.SetValue(node.Value())
		})
	}

	return &node
}

// Value is part of the Cell interface
func (n *Node) Value() int {
	if n.f1 != nil && n.c1 != nil {
		return n.f1(n.c1.Value())
	} else if n.f2 != nil && n.c1 != nil && n.c2 != nil {
		return n.f2(n.c1.Value(), n.c2.Value())
	}

	return n.v
}

// SetValue is part of the InputCell interface
func (n *Node) SetValue(v int) {
	if v != n.v {
		n.v = v

		if n.callbacks != nil {
			for _, f := range n.callbacks {
				f(n.v)
			}
		}
	}
}

// AddCallback is part of the ComputeCell interface
func (n *Node) AddCallback(f func(int)) CallbackHandle {
	if n.callbacks == nil {
		n.callbacks = make(map[int]func(int))
	}

	handle := len(n.callbacks)
	n.callbacks[handle] = f

	return handle
}

// RemoveCallback is part of the ComputeCell interface
func (n *Node) RemoveCallback(h CallbackHandle) {
	if handle, ok := h.(int); ok {
		if _, ok := n.callbacks[handle]; ok {
			delete(n.callbacks, handle)
		}
	}
}
