package drive

// Drive interface
type Drive interface {
	// Read data from address
	Read(lba, size uint64) ([]byte, error)
}
