package cpu

import "context"

// CPU interface
type CPU interface {
	// Freeze
	Freeze(context.Context) error
	// Jump to address
	Jump(context.Context, uint64) error
	// Execute command
	Execute(context.Context) error
}
