package react

const testVersion = 4

// Base implements the Cell interface
type Base struct {
	value int
}

// Value returns the value of b
func (b *Base) Value() int {
	return b.value
}

// Input implements the InputCell interface
type Input struct {
	value    int
	callback func(int) int
}

// Value returns the value of i
func (i *Input) Value() int {
	return i.value
}

// SetValue sets the value of the i
func (i *Input) SetValue(v int) {
	if v != i.value {
		i.value = v
		if i.callback != nil {
			i.callback(v)
		}
	}
}

// Compute implements the ComputeCell interface
type Compute struct {
	c1        Input
	c2        Input
	compute1  func(int) int
	compute2  func(int, int) int
	callbacks []func(int)
}

// Value returns the value of ComputeCell
func (c *Compute) Value() int {
	if c.compute1 != nil {
		c.compute1(c.c1.Value())
	} else if c.compute2 != nil {
		c.compute2(c.c1.Value(), c.c2.Value())
	}

	return 0
}

// AddCallback adds a callback to a compute cell
func (c *Compute) AddCallback(cb func(int)) CallbackHandle {
	if c.callbacks == nil {
		c.callbacks = make([]func(int), 0)
	}

	c.callbacks = append(c.callbacks, cb)

	return len(c.callbacks) - 1
}

// RemoveCallback removes a callback from a compute cell
func (c *Compute) RemoveCallback(cbh CallbackHandle) {
	if i, ok := cbh.(int); ok {
		c.callbacks = append(c.callbacks[:i], c.callbacks[i+1:]...)
	}
}

// React implements the Reactor interface
type React struct {
}

// CreateInput creates an input cell
func (r *React) CreateInput(v int) InputCell {
	return &Input{v, nil}
}

// CreateCompute1 creates a compute cell which computes its value ...
func (r *React) CreateCompute1(c Cell, f func(int) int) ComputeCell {
	return Compute{c1: c, compute1: f}
}

// New creates a new reactor
func New() *React {
	return &React{}
}
