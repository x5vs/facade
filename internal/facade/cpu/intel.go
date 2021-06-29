package cpu

import (
	"context"
	"fmt"
)

// implement CPU interface
type intel struct {
	pos uint64
}

// Freeze cpu
func (i *intel) Freeze(ctx context.Context) error {
	fmt.Println("Intel freeze: ", i.pos)
	return nil
}

// Jump to position
func (i *intel) Jump(ctx context.Context, pos uint64) error {
	i.pos = pos
	fmt.Println("Intel jump: ", i.pos)
	return nil
}

// Execute command on position
func (i *intel) Execute(ctx context.Context) error {
	fmt.Println("Intel execute: ", i.pos)
	return nil
}

// NewIntel create new Intel proccessor
func NewIntel() CPU {
	return &intel{}
}
