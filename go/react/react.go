package react

const testVersion = 4

// Input implements the InputCell interface
type Input struct {
	value    int
	callback func(int) int
}

// Value returns the value of i
func (i *Input) Value() int {
	return i.value
}

// SetValue sets the value of i
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
	i1        *Input
	i2        *Input
	compute1  func(int) int
	compute2  func(int, int) int
	callbacks []func(int)
}

// Value returns the value of ComputeCell
func (c *Compute) Value() int {
	if c.compute1 != nil {
		return c.compute1(c.i1.Value())
	} else if c.compute2 != nil {
		return c.compute2(c.i1.Value(), c.i2.Value())
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
	in, ok := c.(*Input)
	if !ok {
		in = &Input{c.Value(), nil}
	}

	return &Compute{
		i1:        in,
		i2:        nil,
		compute1:  f,
		compute2:  nil,
		callbacks: make([]func(int), 0),
	}
}

// CreateCompute2 creates a compute cell which computes its value ...
func (r *React) CreateCompute2(cell1, cell2 Cell, f func(int, int) int) ComputeCell {
	in1, ok := cell1.(*Input)
	if !ok {
		in1 = &Input{cell1.Value(), nil}
	}

	in2, ok := cell2.(*Input)
	if !ok {
		in2 = &Input{cell2.Value(), nil}
	}

	return &Compute{
		i1:        in1,
		i2:        in2,
		compute1:  nil,
		compute2:  f,
		callbacks: make([]func(int), 0),
	}
}

// New creates a new reactor
func New() *React {
	return &React{}
}
