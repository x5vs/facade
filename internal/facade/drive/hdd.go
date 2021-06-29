package drive

import (
	"context"
	"errors"
	"fmt"
)

// implement Drive interface
type hdd struct {
	size uint64
	data []byte
}

// Read data from drive
func (h *hdd) Read(ctx context.Context, lba, size uint64) ([]byte, error) {
	if size == 0 {
		return nil, errors.New("Invaild data size")
	}
	if lba+size > h.size {
		return nil, errors.New("Invalid address params")
	}
	data := make([]byte, size)
	copy(data, h.data[lba:lba+size])
	fmt.Println("HDD read: ", lba, size, data)
	return data, nil
}

// NewHDD create new HDD drive
func NewHDD(size uint64) Drive {
	data := make([]byte, size)
	return &hdd{size: size, data: data}
}
