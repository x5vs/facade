package drive

import (
	"errors"
	"fmt"
)

// implement Drive interface
type hdd struct {
	size uint64
}

// Read data from drive
func (h *hdd) Read(lba uint64, size int) ([]byte, error) {
	if lba > h.size {
		return nil, errors.New("HDD test error")
	}
	fmt.Println("HDD read: ", lba, size)
	return []byte{3, 2, 1}, nil
}

// NewHDD create new HDD drive
func NewHDD() Drive {
	return &hdd{size: 1024}
}
