package memory

import (
	"errors"
	"fmt"
)

// Memory interface
type ddr2 struct {
	size uint64
}

// Load data to memory
func (d *ddr2) Load(position uint64, data []byte) error {
	if position >= d.size {
		return errors.New("DDR2 load error")
	}
	fmt.Println("DDR2 load memory:", position, data)
	return nil
}

// NewDDR2 create new ddr2 memory
func NewDDR2() Memory {
	return &ddr2{size: 2048}
}
