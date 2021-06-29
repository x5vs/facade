package drive

import (
	"context"
	"errors"
	"fmt"
)

// implement Drive interface
type ssd struct {
	size uint64
	data []byte
}

// Read data from drive
func (s *ssd) Read(ctx context.Context, lba, size uint64) ([]byte, error) {
	if size == 0 {
		return nil, errors.New("Invaild data size")
	}
	if lba+size > s.size {
		return nil, errors.New("Invalid address params")
	}
	data := make([]byte, size)
	copy(data, s.data[lba:lba+size])
	fmt.Println("SSD read: ", lba, size, data)
	return data, nil
}

// NewSDD create new SSD drive
func NewSDD(size uint64) Drive {
	data := make([]byte, size)
	return &ssd{size: size, data: data}
}
