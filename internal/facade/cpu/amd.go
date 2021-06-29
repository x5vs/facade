package cpu

import (
	"context"
	"fmt"
)

// implement CPU interface
type amd struct {
	pos uint64
}

// Freeze cpu
func (a *amd) Freeze(ctx context.Context) error {
	fmt.Println("AMD freeze: ", a.pos)
	return nil
}

// Jump to position
func (a *amd) Jump(ctx context.Context, pos uint64) error {
	a.pos = pos
	fmt.Println("AMD jump: ", a.pos)
	return nil
}

// Execute command on position
func (a *amd) Execute(ctx context.Context) error {
	fmt.Println("AMD execute: ", a.pos)
	return nil
}

// NewAMD create new AMD processor
func NewAMD() CPU {
	return &amd{}
}
