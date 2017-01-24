package react

const testVersion = 4

// Manager manages linked cells.
type Manager struct {
	inputs []InputNode
}

// Node is the basic element of a reactive system.
type Node struct {
	value int
}

// InputNode has a changeable value, chaning the value triggers updates to
// other cells.
type InputNode struct {
	Node
}

// ComputeNode always computes its value based on other cells and can
// call callbacks upon changes.
type ComputeNode struct {
	Node
}

// New creates a new Reactor.
func New() Reactor {
	return Manager{}
}

// CreateInput creates an input cell linked into the reactor
// with the given initial value.
func (m Manager) CreateInput(n int) InputCell {
	if m.inputs == nil {
		m.inputs = make([]InputNode, 0)
	}

	input := InputNode{Node{n}}
	m.inputs = append(m.inputs, input)
	return input
}

// CreateCompute1 creates a compute cell which computes its value
// based on one other cell. The compute function will only be called
// if the value of the passed cell changes.
func (m Manager) CreateCompute1(c Cell, compute func(int) int) ComputeCell {
	return nil
}

// CreateCompute2 is like CreateCompute1, but depending on two cells.
// The compute function will only be called if the value of any of the
// passed cells changes.
func (m Manager) CreateCompute2(c1 Cell, c2 Cell, compute func(int, int) int) ComputeCell {
	return nil
}

// Value returns the current value of the cell
func (n Node) Value() int {
	return n.value
}

// SetValue sets the value of the cell
func (i InputNode) SetValue(v int) {
	i.value = v
}

// AddCallback adds a callback which will be called when the value changes.
// It returns a callback handle which can be used to remove the callback.
func (c *ComputeNode) AddCallback(func(int)) CallbackHandle {
	return nil
}

// RemoveCallback removes a previously added callback, if it exists.
func (c *ComputeNode) RemoveCallback(CallbackHandle) {
}
