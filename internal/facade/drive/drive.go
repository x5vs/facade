package drive

import "context"

// Drive interface
type Drive interface {
	// Read data from address
	Read(ctx context.Context, lba, size uint64) ([]byte, error)
}
