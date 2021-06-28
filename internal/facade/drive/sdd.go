package drive

import (
	"errors"
	"fmt"
)

// implement Drive interface
type ssd struct {
	size uint64
}

// Read data from drive
func (s *ssd) Read(lba uint64, size int) ([]byte, error) {
	if lba > s.size {
		return nil, errors.New("SSH test error")
	}
	fmt.Println("SSD read: ", lba, size)
	return []byte{1, 2, 3}, nil
}

// NewSDD create new SSD drive
func NewSDD() Drive {
	return &ssd{size: 512}
}
