// Package ptr provides convenience methods for retrieving the pointer reference to any value
// or dereferencing any pointer to its value.
package ptr

// Return the pointer of any value V
func P[V any](v V) *V {
	return &v
}

// Return the dereferenced value of a pointer to any P
func V[P any](p *P) P {
	return *p
}
