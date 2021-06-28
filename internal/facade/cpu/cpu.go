package cpu

// Proccessor interface
type Proccessor interface {
	// Freeze
	Freeze()
	// Jump to address
	Jump(uint64)
	// Execute command
	Execute()
}
