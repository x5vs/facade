package drive

// Drive interface
type Drive interface {
	// Read data from address
	Read(lba uint64, size int) ([]byte, error)
}
