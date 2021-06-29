package memory

import (
	"context"
	"errors"
	"fmt"
)

// Memory interface
type ddr1 struct {
	size uint64
	data []byte
}

// Load data to memory
func (d *ddr1) Load(ctx context.Context, position uint64, data []byte) error {
	total := uint64(len(data))
	if position+total > d.size {
		return errors.New("DDR1 load error")
	}
	copy(d.data[position:], data)
	fmt.Println("DDR1 load memory:", position, data)
	fmt.Println("DDR1 memory:", d.data)
	return nil
}

// NewDDR1 create new ddr1 memory
func NewDDR1(size uint64) Memory {
	return &ddr1{size: size, data: make([]byte, size)}
}
