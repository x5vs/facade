package memory

// Memory interface
type Memory interface {
	// Load data to memory
	Load(position uint64, data []byte) error
}
