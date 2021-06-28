package memory

import (
	"errors"
	"fmt"
)

// Memory interface
type ddr1 struct {
	size uint64
}

// Load data to memory
func (d *ddr1) Load(position uint64, data []byte) error {
	if position >= d.size {
		return errors.New("DDR1 load error")
	}
	fmt.Println("DDR1 load memory:", position, data)
	return nil
}

// NewDDR1 create new ddr1 memory
func NewDDR1() Memory {
	return &ddr1{size: 1024}
}
