package cpu

import "fmt"

// implement Proccessor interface
type amd struct {
	pos uint64
}

// Freeze cpu
func (a *amd) Freeze() {
	fmt.Println("AMD freeze: ", a.pos)
}

// Jump to position
func (a *amd) Jump(pos uint64) {
	a.pos = pos
	fmt.Println("AMD jump: ", a.pos)
}

// Execute command on position
func (a *amd) Execute() {
	fmt.Println("AMD execute: ", a.pos)
}

// NewAMD creat new AMD processor
func NewAMD() Proccessor {
	return &amd{pos: 0}
}
