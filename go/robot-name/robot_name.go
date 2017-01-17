package robotname

// Robot represents a robot.
type Robot struct {
	name string
}

// Name returns the name of robot r.
func (r *Robot) Name() string {
	return r.name
}

// Reset gives robot r a new name, different from its previous name.
func (r *Robot) Reset() {

}
