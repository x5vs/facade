package memory

import (
	"errors"
	"fmt"
)

// Memory interface
type ddr2 struct {
	size uint64
	data []byte
}

// Load data to memory
func (d *ddr2) Load(position uint64, data []byte) error {
	total := uint64(len(data))
	if position+total > d.size {
		return errors.New("DDR2 load error")
	}
	copy(d.data[position:], data)
	fmt.Println("DDR2 load memory:", position, data)
	fmt.Println("DDR2 memory:", d.data)
	return nil
}

// NewDDR2 create new ddr2 memory
func NewDDR2(size uint64) Memory {
	return &ddr2{size: size, data: make([]byte, size)}
}
