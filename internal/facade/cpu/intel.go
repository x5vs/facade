package cpu

import "fmt"

// implement Proccessor interface
type intel struct {
	pos uint64
}

// Freeze cpu
func (i *intel) Freeze() {
	fmt.Println("Intel freeze: ", i.pos)
}

// Jump to position
func (i *intel) Jump(pos uint64) {
	i.pos = pos
	fmt.Println("Intel jump: ", i.pos)
}

// Execute command on position
func (i *intel) Execute() {
	fmt.Println("Intel execute: ", i.pos)
}

// NewIntel create new Intel proccessor
func NewIntel() Proccessor {
	return &intel{pos: 0}
}
