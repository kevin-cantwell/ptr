// Package ptr contains various utilities for working with pointers in-line.
// There are times when taking the address of a pointer or dereferencing a pointer
// requires more than one line of code (often a nil check or variable assignment).
// The ptr suite maps a common set of such patterns into functions that can be
// invoked on a single line.
package ptr
