package memory

import "context"

// Memory interface
type Memory interface {
	// Load data to memory
	Load(ctx context.Context, position uint64, data []byte) error
}
